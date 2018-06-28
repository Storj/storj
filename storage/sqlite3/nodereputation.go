// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package nodereputation

import (
	"database/sql"
	"fmt"
	"strings"

	rep "storj.io/storj/pkg/nodereputation"
	proto "storj.io/storj/protos/nodereputation"

	// import of sqlite3 for side effects
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeebo/errs"
)

// CreateTableError is an error class for errors related to the reputation package
var CreateTableError = errs.Class("reputation table creation error")

// CreateNodeError is an error class for errors related to the reputation package
var CreateNodeError = errs.Class("reputation node creation error")

// SelectError is an error class for errors related to the reputation package
var SelectError = errs.Class("reputation selection error")

// UpdateError is an error class for errors related to the reputation package
var UpdateError = errs.Class("reputation update error")

// StartDBError is an error class for errors related to the reputation package
var StartDBError = errs.Class("reputation start sqlite3 error")

// startDB starts a sqlite3 database from the file path parameter
func startDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, StartDBError.Wrap(err)
	}

	return db, nil
}

// createReputationTable creates a table in sqlite3 based on the create table string parameter
func createReputationTable(db *sql.DB) error {
	var res []string

	for _, feature := range proto.Feature_name {
		for _, param := range proto.Parameter_name {
			res = append(res, fmt.Sprintf("%s_%s", feature, param))
		}
		for _, state := range proto.BetaStateCols_name {
			res = append(res, fmt.Sprintf("%s_%s", feature, state))
		}
	}

	timefmt := "%Y-%m-%d %H:%M:%f"

	createTableStmt := fmt.Sprintf(`CREATE table node_reputation (
		node_name TEXT NOT NULL,
		last_seen timestamp DEFAULT(STRFTIME('%s', 'NOW')) NOT NULL,
		%s,
		PRIMARY KEY(node_name, last_seen));`,
		timefmt,
		strings.Join(res, ",\n"),
	)

	_, err := db.Exec(createTableStmt)
	if err != nil {
		return CreateTableError.Wrap(err)
	}
	return nil
}

func createNewNodeRecord(db *sql.DB, nodeName string) error {

	type paramValue struct {
		param proto.Parameter
		val   float64
	}

	type stateValue struct {
		state proto.BetaStateCols
		val   proto.UpdateRepValue
	}

	params := []paramValue{
		paramValue{
			param: proto.Parameter_BAD_RECALL,
			val:   0.995,
		},
		paramValue{
			param: proto.Parameter_GOOD_RECALL,
			val:   0.99,
		},
		paramValue{
			param: proto.Parameter_WEIGHT_DENOMINATOR,
			val:   10000.0,
		},
	}
	states := []stateValue{
		stateValue{
			state: proto.BetaStateCols_CUMULATIVE_SUM_REPUTATION,
			val:   proto.UpdateRepValue_ZERO,
		},
		stateValue{
			state: proto.BetaStateCols_CURRENT_REPUTATION,
			val:   proto.UpdateRepValue_ZERO,
		},
		stateValue{
			state: proto.BetaStateCols_FEATURE_COUNTER,
			val:   proto.UpdateRepValue_ZERO,
		},
	}

	insertNewNodeName(db, nodeName)
	for _, feature := range proto.Feature_name {
		for _, pair := range params {
			err := updateNodeParameters(db, nodeName, feature, pair.param, pair.val)
			if err != nil {
				CreateNodeError.Wrap(err)
			}
		}

		for _, state := range states {
			err := updateNodeState(db, nodeName, feature, state.state, state.val)
			if err != nil {
				CreateNodeError.Wrap(err)
			}
		}
	}

	return nil
}

func insertNewNodeName(db *sql.DB, nodeName string) error {
	tx, err := db.Begin()
	if err != nil {
		return CreateNodeError.Wrap(err)
	}
	defer tx.Rollback()

	createNodeString := `INSERT
	INTO node_reputation (node_name) values (?);`

	insertStmt, err := tx.Prepare(createNodeString)
	if err != nil {
		return CreateNodeError.Wrap(err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(nodeName)
	if err != nil {
		return CreateNodeError.Wrap(err)
	}

	return tx.Commit()
}

type nodeFeature struct {
	nodeName          string
	feature           string
	goodRecall        float64
	badRecall         float64
	featureCounter    float64
	weightDenominator float64
	cumulativeSum     float64
	reputation        float64
}

func selectNodeFeature(db *sql.DB, nodeName string, feature string) (nodeFeature, error) {
	var res nodeFeature

	stmt := selectFeatureStmt(feature, nodeName)
	rows, err := db.Query(stmt)
	if err != nil {
		return res, SelectError.Wrap(err)
	}
	defer rows.Close()

	res, err = selectedFeaturesToNodeRecord(rows)
	if err != nil {
		return res, SelectError.Wrap(err)
	}

	res.nodeName = nodeName
	res.feature = feature

	err = rows.Err()
	if err != nil {
		return res, SelectError.Wrap(err)
	}

	return res, nil
}

//
func getRep(db *sql.DB, nodeName string) ([]nodeFeature, error) {
	var res []nodeFeature
	updateRes := func(feature string) (nodeFeature, error) {
		var newRes nodeFeature
		newRes, err := selectNodeFeature(db, nodeName, feature)
		if err != nil {
			return newRes, SelectError.Wrap(err)
		}
		return newRes, nil
	}

	for _, feature := range proto.Feature_name {
		update, err := updateRes(feature)
		if err != nil {
			SelectError.Wrap(err)
		}
		res = append(res, update)
	}
	return res, nil
}

func matchRepOrderStmt(features []proto.Feature, notIn []string) string {
	var exclude []string

	for _, not := range notIn {
		exclude = append(exclude, fmt.Sprintf(`'%s'`, not))
	}

	var ordered []string

	for _, feature := range features {
		ordered = append(ordered,
			fmt.Sprintf("%s_%s",
				feature.String(), proto.BetaStateCols_CURRENT_REPUTATION.String()))
	}

	selectNodesStmt := fmt.Sprintf(`SELECT node_name
	FROM node_reputation
	WHERE node_name NOT IN (%s)
	ORDER BY %s`, strings.Join(exclude, ","), strings.Join(ordered, ","))

	return selectNodesStmt
}

//
func matchRepOrder(db *sql.DB, features []proto.Feature, notIn []string) ([]string, error) {
	rows, err := db.Query(matchRepOrderStmt(features, notIn))
	if err != nil {
		return nil, SelectError.Wrap(err)
	}
	defer rows.Close()

	var res []string

	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			return nil, SelectError.Wrap(err)
		}

		res = append(res, s)
	}

	return res, nil

}

func selectAllBetaStateStmt() string {
	res := "SELECT"
	fromWhere := `FROM node_reputation
	WHERE node_name = ?`

	var repState []string

	for _, feature := range proto.Feature_name {
		for _, state := range proto.BetaStateCols_name {
			repState = append(repState, fmt.Sprintf("%s_%s", feature, state))
		}
	}

	joined := strings.Join(repState, ",")

	res = res + joined + fromWhere

	return res
}

//
func updateNodeRecord(db *sql.DB, nodeName string, feature proto.Feature, value proto.UpdateRepValue) error {
	node, err := selectNodeFeature(db, nodeName, feature.String())
	if err != nil {
		return UpdateError.Wrap(err)
	}
	betaRes := rep.Beta(node.badRecall, node.goodRecall, node.weightDenominator, node.featureCounter, node.cumulativeSum, updateToFloat(value))
	newSum := node.cumulativeSum + betaRes.Reputation
	newCount := node.featureCounter + 1

	tx, err := db.Begin()
	if err != nil {
		return UpdateError.Wrap(err)
	}
	defer tx.Rollback()

	updateStringRep := updateFeatureRepStmt(nodeName, feature.String(), proto.BetaStateCols_CURRENT_REPUTATION.String(), updateToFloat(value))

	_, err = tx.Exec(updateStringRep)
	if err != nil {
		return UpdateError.Wrap(err)
	}

	updateStringSum := updateFeatureRepStmt(nodeName, feature.String(), proto.BetaStateCols_CUMULATIVE_SUM_REPUTATION.String(), newSum)

	_, err = tx.Exec(updateStringSum)
	if err != nil {
		return UpdateError.Wrap(err)
	}

	updateStringCount := updateFeatureRepStmt(nodeName, feature.String(), proto.BetaStateCols_FEATURE_COUNTER.String(), newCount)

	_, err = tx.Exec(updateStringCount)
	if err != nil {
		return UpdateError.Wrap(err)
	}

	return tx.Commit()
}

func updateNodeParameters(db *sql.DB, nodeName string, feature string, parameter proto.Parameter, parameterValue float64) error {
	tx, err := db.Begin()
	if err != nil {
		return UpdateError.Wrap(err)
	}
	defer tx.Rollback()

	updateParamString := fmt.Sprintf(`UPDATE node_reputation
	 SET %s_%s = %.4f
	 WHERE node_name = '%s';`, feature, parameter.String(), parameterValue, nodeName)

	_, err = tx.Exec(updateParamString)
	if err != nil {
		return UpdateError.Wrap(err)
	}

	return tx.Commit()
}

func updateNodeState(db *sql.DB, nodeName string, feature string, state proto.BetaStateCols, stateValue proto.UpdateRepValue) error {
	tx, err := db.Begin()
	if err != nil {
		return UpdateError.Wrap(err)
	}
	defer tx.Rollback()

	updateParamString := fmt.Sprintf(`UPDATE node_reputation
	 SET %s_%s = %.4f
	 WHERE node_name = '%s';`, feature, state.String(), updateToFloat(stateValue), nodeName)

	_, err = tx.Exec(updateParamString)
	if err != nil {
		return UpdateError.Wrap(err)
	}

	return tx.Commit()
}

// assumtion one row per node id
func selectedFeaturesToNodeRecord(rows *sql.Rows) (nodeFeature, error) {
	var res nodeFeature

	for rows.Next() {
		err := rows.Scan(
			&res.badRecall,
			&res.weightDenominator,
			&res.goodRecall,
			&res.featureCounter,
			&res.cumulativeSum,
			&res.reputation,
		)
		if err != nil {
			return res, SelectError.Wrap(err)
		}
	}

	return res, nil
}

func updateFeatureRepStmt(nodeName string, feature string, state string, value float64) string {
	update := `UPDATE node_reputation
	SET last_seen = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'),`

	return fmt.Sprintf(`%s
		%s_%s = %.4f
		WHERE node_name = '%s';`, update, feature, state, value, nodeName)
}

func selectFeatureStmt(feature string, nodeName string) string {
	var cols []string

	for _, state := range proto.Parameter_name {
		cols = append(cols, fmt.Sprintf("%s_%s", feature, state))
	}

	for _, params := range proto.BetaStateCols_name {
		cols = append(cols, fmt.Sprintf("%s_%s", feature, params))
	}

	return fmt.Sprintf(`SELECT 
		%s FROM node_reputation WHERE node_name = '%s';`,
		strings.Join(cols, ",\n"), nodeName)
}

func updateToFloat(val proto.UpdateRepValue) float64 {
	res := float64(0)

	switch val {
	case 0:
		res = float64(-1)
	case 1:
		res = float64(-0.5)
	case 2:
		res = float64(-0.25)
	case 3:
		res = float64(0)
	case 4:
		res = float64(0.25)
	case 5:
		res = float64(0.5)
	case 6:
		res = float64(1)
	}

	return res
}
