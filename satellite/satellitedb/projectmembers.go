// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"strings"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/satellite/console"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// ProjectMembers exposes methods to manage ProjectMembers table in database.
type projectMembers struct {
	methods dbx.Methods
	db      *dbx.DB
}

// GetByMemberID is a method for querying project member from the database by memberID.
func (pm *projectMembers) GetByMemberID(ctx context.Context, memberID uuid.UUID) (_ []console.ProjectMember, err error) {
	defer mon.Task()(&ctx)(&err)
	projectMembersDbx, err := pm.methods.All_ProjectMember_By_MemberId(ctx, dbx.ProjectMember_MemberId(memberID[:]))
	if err != nil {
		return nil, err
	}

	return projectMembersFromDbxSlice(ctx, projectMembersDbx)
}

// GetByProjectID is a method for querying project members from the database by projectID, offset and limit.
func (pm *projectMembers) GetPagedByProjectID(ctx context.Context, projectID uuid.UUID, cursor console.ProjectMembersCursor) (_ *console.ProjectMembersPage, err error) {
	defer mon.Task()(&ctx)(&err)

	search := "%" + strings.Replace(cursor.Search, " ", "%", -1) + "%"

	if cursor.Limit > 50 {
		cursor.Limit = 50
	}

	if cursor.Page == 0 {
		return nil, errs.New("page cannot be 0")
	}

	page := &console.ProjectMembersPage{
		Search:         cursor.Search,
		Limit:          cursor.Limit,
		Offset:         uint64((cursor.Page - 1) * cursor.Limit),
		Order:          cursor.Order,
		OrderDirection: cursor.OrderDirection,
	}

	countQuery := pm.db.Rebind(`
		SELECT COUNT(*)
		FROM project_members pm 
		INNER JOIN users u ON pm.member_id = u.id
		WHERE pm.project_id = ?
		AND ( u.email LIKE ? OR 
			  u.full_name LIKE ? OR
			  u.short_name LIKE ? 
		)`)

	countRow := pm.db.QueryRowContext(ctx,
		countQuery,
		projectID[:],
		search,
		search,
		search)

	err = countRow.Scan(&page.TotalCount)
	if err != nil {
		return nil, err
	}
	if page.TotalCount == 0 {
		return page, nil
	}
	if page.Offset > page.TotalCount-1 {
		return nil, errs.New("page is out of range")
	}
	// TODO: LIKE is case-sensitive postgres, however this should be case-insensitive and possibly allow typos
	reboundQuery := pm.db.Rebind(`
		SELECT pm.*
			FROM project_members pm
				INNER JOIN users u ON pm.member_id = u.id
				WHERE pm.project_id = ?
				AND ( u.email LIKE ? OR
					u.full_name LIKE ? OR
					u.short_name LIKE ? )
					ORDER BY ` + sanitizedOrderColumnName(cursor.Order) + `
					` + sanitizeOrderDirectionName(page.OrderDirection) + `	
					LIMIT ? OFFSET ?`)

	rows, err := pm.db.QueryContext(ctx,
		reboundQuery,
		projectID[:],
		search,
		search,
		search,
		page.Limit,
		page.Offset)

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	if err != nil {
		return nil, err
	}

	var projectMembers []console.ProjectMember
	for rows.Next() {
		pm := console.ProjectMember{}
		var memberIDBytes, projectIDBytes []uint8
		var memberID, projectID uuid.UUID

		err = rows.Scan(&memberIDBytes, &projectIDBytes, &pm.CreatedAt)
		if err != nil {
			return nil, err
		}

		memberID, err := bytesToUUID(memberIDBytes)
		if err != nil {
			return nil, err
		}

		projectID, err = bytesToUUID(projectIDBytes)
		if err != nil {
			return nil, err
		}

		pm.ProjectID = projectID
		pm.MemberID = memberID

		projectMembers = append(projectMembers, pm)
	}

	page.ProjectMembers = projectMembers
	page.Order = cursor.Order

	page.PageCount = uint(page.TotalCount / uint64(cursor.Limit))
	if page.TotalCount%uint64(cursor.Limit) != 0 {
		page.PageCount++
	}

	page.CurrentPage = cursor.Page

	return page, err
}

// Insert is a method for inserting project member into the database.
func (pm *projectMembers) Insert(ctx context.Context, memberID, projectID uuid.UUID) (_ *console.ProjectMember, err error) {
	defer mon.Task()(&ctx)(&err)
	createdProjectMember, err := pm.methods.Create_ProjectMember(ctx,
		dbx.ProjectMember_MemberId(memberID[:]),
		dbx.ProjectMember_ProjectId(projectID[:]))
	if err != nil {
		return nil, err
	}

	return projectMemberFromDBX(ctx, createdProjectMember)
}

// Delete is a method for deleting project member by memberID and projectID from the database.
func (pm *projectMembers) Delete(ctx context.Context, memberID, projectID uuid.UUID) (err error) {
	defer mon.Task()(&ctx)(&err)
	_, err = pm.methods.Delete_ProjectMember_By_MemberId_And_ProjectId(
		ctx,
		dbx.ProjectMember_MemberId(memberID[:]),
		dbx.ProjectMember_ProjectId(projectID[:]),
	)

	return err
}

// projectMemberFromDBX is used for creating ProjectMember entity from autogenerated dbx.ProjectMember struct
func projectMemberFromDBX(ctx context.Context, projectMember *dbx.ProjectMember) (_ *console.ProjectMember, err error) {
	defer mon.Task()(&ctx)(&err)
	if projectMember == nil {
		return nil, errs.New("projectMember parameter is nil")
	}

	memberID, err := bytesToUUID(projectMember.MemberId)
	if err != nil {
		return nil, err
	}

	projectID, err := bytesToUUID(projectMember.ProjectId)
	if err != nil {
		return nil, err
	}

	return &console.ProjectMember{
		MemberID:  memberID,
		ProjectID: projectID,
		CreatedAt: projectMember.CreatedAt,
	}, nil
}

// sanitizedOrderColumnName return valid order by column
func sanitizedOrderColumnName(pmo console.ProjectMemberOrder) string {
	switch pmo {
	case 2:
		return "u.email"
	case 3:
		return "u.created_at"
	default:
		return "u.full_name"
	}
}

func sanitizeOrderDirectionName(pmo console.ProjectMemberOrderDirection) string {
	if pmo == 2 {
		return "DESC"
	}

	return "ASC"
}

// projectMembersFromDbxSlice is used for creating []ProjectMember entities from autogenerated []*dbx.ProjectMember struct
func projectMembersFromDbxSlice(ctx context.Context, projectMembersDbx []*dbx.ProjectMember) (_ []console.ProjectMember, err error) {
	defer mon.Task()(&ctx)(&err)
	var projectMembers []console.ProjectMember
	var errors []error

	// Generating []dbo from []dbx and collecting all errors
	for _, projectMemberDbx := range projectMembersDbx {
		projectMember, err := projectMemberFromDBX(ctx, projectMemberDbx)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		projectMembers = append(projectMembers, *projectMember)
	}

	return projectMembers, errs.Combine(errors...)
}
