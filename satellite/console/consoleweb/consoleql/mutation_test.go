// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleql_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/common/uuid"
	"storj.io/storj/private/post"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/accounting"
	"storj.io/storj/satellite/accounting/live"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/console/consoleauth"
	"storj.io/storj/satellite/console/consoleweb/consoleql"
	"storj.io/storj/satellite/mailservice"
	"storj.io/storj/satellite/payments/paymentsconfig"
	"storj.io/storj/satellite/payments/stripecoinpayments"
	"storj.io/storj/satellite/rewards"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
	"storj.io/storj/storage/redis/redisserver"
)

// discardSender discard sending of an actual email.
type discardSender struct{}

// SendEmail immediately returns with nil error.
func (*discardSender) SendEmail(ctx context.Context, msg *post.Message) error {
	return nil
}

// FromAddress returns empty post.Address.
func (*discardSender) FromAddress() post.Address {
	return post.Address{}
}

func TestGraphqlMutation(t *testing.T) {
	satellitedbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db satellite.DB) {
		log := zaptest.NewLogger(t)

		partnersService := rewards.NewPartnersService(
			log.Named("partners"),
			rewards.DefaultPartnersDB,
			[]string{
				"https://us-central-1.tardigrade.io/",
				"https://asia-east-1.tardigrade.io/",
				"https://europe-west-1.tardigrade.io/",
			},
		)

		paymentsConfig := stripecoinpayments.Config{}
		payments := stripecoinpayments.NewService(log, paymentsConfig, db.StripeCoinPayments(), db.Console().Projects(), db.ProjectAccounting(), 0, 0, 0)

		service, err := console.NewService(
			log.Named("console"),
			&consoleauth.Hmac{Secret: []byte("my-suppa-secret-key")},
			db.Console(),
			db.ProjectAccounting(),
			db.Rewards(),
			partnersService,
			paymentsService.Accounts(),
			console.Config{PasswordCost: console.TestPasswordCost, DefaultProjectLimit: 5},
			5000,
		)
		require.NoError(t, err)

		mailService, err := mailservice.New(log, &discardSender{}, "testdata")
		require.NoError(t, err)
		defer ctx.Check(mailService.Close)

		rootObject := make(map[string]interface{})
		rootObject["origin"] = "http://doesntmatter.com/"
		rootObject[consoleql.ActivationPath] = "?activationToken="
		rootObject[consoleql.SignInPath] = "login"
		rootObject[consoleql.LetUsKnowURL] = "letUsKnowURL"
		rootObject[consoleql.ContactInfoURL] = "contactInfoURL"
		rootObject[consoleql.TermsAndConditionsURL] = "termsAndConditionsURL"

		schema, err := consoleql.CreateSchema(log, service, mailService)
		require.NoError(t, err)

		createUser := console.CreateUser{
			FullName:  "John Roll",
			ShortName: "Roll",
			Email:     "test@mail.test",
			PartnerID: "120bf202-8252-437e-ac12-0e364bee852e",
			Password:  "123a123",
		}
		refUserID := ""

		regToken, err := service.CreateRegToken(ctx, 1)
		require.NoError(t, err)

		rootUser, err := service.CreateUser(ctx, createUser, regToken.Secret, refUserID)
		require.NoError(t, err)
		require.Equal(t, createUser.PartnerID, rootUser.PartnerID.String())

		err = paymentsService.Accounts().Setup(ctx, rootUser.ID, rootUser.Email)
		require.NoError(t, err)

		activationToken, err := service.GenerateActivationToken(ctx, rootUser.ID, rootUser.Email)
		require.NoError(t, err)

		err = service.ActivateAccount(ctx, activationToken)
		require.NoError(t, err)

		token, err := service.Token(ctx, createUser.Email, createUser.Password)
		require.NoError(t, err)

		sauth, err := service.Authorize(consoleauth.WithAPIKey(ctx, []byte(token)))
		require.NoError(t, err)

		authCtx := console.WithAuth(ctx, sauth)

		testQuery := func(t *testing.T, query string) interface{} {
			result := graphql.Do(graphql.Params{
				Schema:        schema,
				Context:       authCtx,
				RequestString: query,
				RootObject:    rootObject,
			})

			for _, err := range result.Errors {
				if err.OriginalError() != nil {
					return nil, err
				}
			}
			require.False(t, result.HasErrors())

			return result.Data, nil
		}

		token, err = service.Token(ctx, rootUser.Email, createUser.Password)
		require.NoError(t, err)

		sauth, err = service.Authorize(consoleauth.WithAPIKey(ctx, []byte(token)))
		require.NoError(t, err)

		authCtx = console.WithAuth(ctx, sauth)

		var projectIDField string
		t.Run("Create project mutation", func(t *testing.T) {
			projectInfo := console.ProjectInfo{
				Name:        "Project name",
				Description: "desc",
			}

			query := fmt.Sprintf(
				"mutation {createProject(input:{name:\"%s\",description:\"%s\"}){name,description,id,createdAt}}",
				projectInfo.Name,
				projectInfo.Description,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			project := data[consoleql.CreateProjectMutation].(map[string]interface{})

			assert.Equal(t, projectInfo.Name, project[consoleql.FieldName])
			assert.Equal(t, projectInfo.Description, project[consoleql.FieldDescription])

			projectIDField = project[consoleql.FieldID].(string)
		})

		projectID, err := uuid.FromString(projectIDField)
		require.NoError(t, err)

		project, err := service.GetProject(authCtx, projectID)
		require.NoError(t, err)
		require.Equal(t, rootUser.PartnerID, project.PartnerID)

		regTokenUser1, err := service.CreateRegToken(ctx, 1)
		require.NoError(t, err)

		user1, err := service.CreateUser(authCtx, console.CreateUser{
			FullName: "User1",
			Email:    "u1@mail.test",
			Password: "123a123",
		}, regTokenUser1.Secret, refUserID)
		require.NoError(t, err)

		t.Run("Activation", func(t *testing.T) {
			activationToken1, err := service.GenerateActivationToken(
				ctx,
				user1.ID,
				"u1@mail.test",
			)
			require.NoError(t, err)

			err = service.ActivateAccount(ctx, activationToken1)
			require.NoError(t, err)

			user1.Email = "u1@mail.test"
		})

		regTokenUser2, err := service.CreateRegToken(ctx, 1)
		require.NoError(t, err)

		user2, err := service.CreateUser(authCtx, console.CreateUser{
			FullName: "User1",
			Email:    "u2@mail.test",
			Password: "123a123",
		}, regTokenUser2.Secret, refUserID)
		require.NoError(t, err)

		t.Run("Activation", func(t *testing.T) {
			activationToken2, err := service.GenerateActivationToken(
				ctx,
				user2.ID,
				"u2@mail.test",
			)
			require.NoError(t, err)

			err = service.ActivateAccount(ctx, activationToken2)
			require.NoError(t, err)

			user2.Email = "u2@mail.test"
		})

		t.Run("Add project members mutation", func(t *testing.T) {
			query := fmt.Sprintf(
				"mutation {addProjectMembers(projectID:\"%s\",email:[\"%s\",\"%s\"]){id,name,members(cursor: { limit: 50, search: \"\", page: 1, order: 1, orderDirection: 2 }){projectMembers{joinedAt}}}}",
				project.ID.String(),
				user1.Email,
				user2.Email,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			proj := data[consoleql.AddProjectMembersMutation].(map[string]interface{})

			members := proj[consoleql.FieldMembers].(map[string]interface{})
			projectMembers := members[consoleql.FieldProjectMembers].([]interface{})

			assert.Equal(t, project.ID.String(), proj[consoleql.FieldID])
			assert.Equal(t, project.Name, proj[consoleql.FieldName])
			assert.Equal(t, 3, len(projectMembers))
		})

		t.Run("Delete project members mutation", func(t *testing.T) {
			query := fmt.Sprintf(
				"mutation {deleteProjectMembers(projectID:\"%s\",email:[\"%s\",\"%s\"]){id,name,members(cursor: { limit: 50, search: \"\", page: 1, order: 1, orderDirection: 2 }){projectMembers{user{id}}}}}",
				project.ID.String(),
				user1.Email,
				user2.Email,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			proj := data[consoleql.DeleteProjectMembersMutation].(map[string]interface{})

			members := proj[consoleql.FieldMembers].(map[string]interface{})
			projectMembers := members[consoleql.FieldProjectMembers].([]interface{})
			rootMember := projectMembers[0].(map[string]interface{})[consoleql.UserType].(map[string]interface{})

			assert.Equal(t, project.ID.String(), proj[consoleql.FieldID])
			assert.Equal(t, project.Name, proj[consoleql.FieldName])
			assert.Equal(t, 1, len(members))

			assert.Equal(t, rootUser.ID.String(), rootMember[consoleql.FieldID])
		})

		var keyID string
		t.Run("Create api key mutation", func(t *testing.T) {
			keyName := "key1"
			query := fmt.Sprintf(
				"mutation {createAPIKey(projectID:\"%s\",name:\"%s\"){key,keyInfo{id,name,projectID,partnerId}}}",
				project.ID.String(),
				keyName,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			createAPIKey := data[consoleql.CreateAPIKeyMutation].(map[string]interface{})

			key := createAPIKey[consoleql.FieldKey].(string)
			keyInfo := createAPIKey[consoleql.APIKeyInfoType].(map[string]interface{})

			assert.NotEqual(t, "", key)

			assert.Equal(t, keyName, keyInfo[consoleql.FieldName])
			assert.Equal(t, project.ID.String(), keyInfo[consoleql.FieldProjectID])
			assert.Equal(t, rootUser.PartnerID.String(), keyInfo[consoleql.FieldPartnerID])

			keyID = keyInfo[consoleql.FieldID].(string)
		})

		t.Run("Delete api key mutation", func(t *testing.T) {
			id, err := uuid.FromString(keyID)
			require.NoError(t, err)

			info, err := service.GetAPIKeyInfo(authCtx, id)
			require.NoError(t, err)

			query := fmt.Sprintf(
				"mutation {deleteAPIKeys(id:[\"%s\"]){name,projectID}}",
				keyID,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			keyInfoList := data[consoleql.DeleteAPIKeysMutation].([]interface{})

			for _, k := range keyInfoList {
				keyInfo := k.(map[string]interface{})

				assert.Equal(t, info.Name, keyInfo[consoleql.FieldName])
				assert.Equal(t, project.ID.String(), keyInfo[consoleql.FieldProjectID])
			}
		})

		const testName = "testName"
		const testDescription = "test description"

		t.Run("Update project mutation", func(t *testing.T) {
			query := fmt.Sprintf(
				"mutation {updateProject(id:\"%s\",name:\"%s\",description:\"%s\"){id,name,description}}",
				project.ID.String(),
				testName,
				testDescription,
			)

			result, err := testQuery(t, query)
			require.NoError(t, err)

			data := result.(map[string]interface{})
			proj := data[consoleql.UpdateProjectMutation].(map[string]interface{})

			assert.Equal(t, project.ID.String(), proj[consoleql.FieldID])
			assert.Equal(t, testName, proj[consoleql.FieldName])
			assert.Equal(t, testDescription, proj[consoleql.FieldDescription])
		})
	})
}
