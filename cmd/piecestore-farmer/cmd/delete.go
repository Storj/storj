// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // sqlite driver
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a farmer node by ID",
	Long:  "Delete config and all data stored on node by node ID",
	RunE:  deleteNode,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// deleteNode deletes a farmer node by ID
func deleteNode(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errs.New("no id specified")
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	nodeID := args[0]

	_, configFile := SetConfigPath(home, nodeID)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return errs.New("Invalid node id. Config file does not exist")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// get folder for stored data
	piecestoreDir := viper.GetString("piecestore.dir")
	piecestoreDir = filepath.Join(piecestoreDir, fmt.Sprintf("store-%s", nodeID))

	// remove all folders and files stored on node
	if err := os.RemoveAll(piecestoreDir); err != nil {
		return err
	}

	// delete node config
	err = os.Remove(configFile)
	if err != nil {
		return err
	}

	zap.S().Infof("Deleted node: %s", nodeID)
	fmt.Printf("Deleted node: %s", nodeID)

	return nil
}
