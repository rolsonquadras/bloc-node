/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package node

import (
	"fmt"

	"github.com/hyperledger/fabric/peer/node"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func startCmd() *cobra.Command {
	return nodeStartCmd
}

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		return serve()
	},
}

func serve() error {
	logger.Info("!!!! start bloc node")
	// start the fabric peer
	if err := node.Start(); err != nil {
		return errors.WithMessage(err, "start fabric peer failed")
	}

	// start bloc node components
	return nil
}
