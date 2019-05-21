/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package startcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// peer
type peer interface {
	// Init peer
	Init()
	// Start peer
	Start() error
}

// Cmd return start cmd
func Cmd(p peer) *cobra.Command {
	return newCmd(p)
}

func newCmd(p peer) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts the node.",
		Long:  `Starts a node that interacts with the network.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return fmt.Errorf("trailing args detected")
			}
			// Parsing of the command line is done so silence cmd usage
			cmd.SilenceUsage = true
			return start(p)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) { initCmd(p) },
	}
}

func initCmd(p peer) {
	p.Init()
}

func start(p peer) error {
	// start ledger
	return p.Start()
}
