/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/spf13/cobra"
	ledgercmd "github.com/trustbloc/bloc-node/cmd/blocnode/ledgercmd"
	"github.com/trustbloc/bloc-node/pkg/ledger/fabric"
)

// peer
type peer interface {
	// Start peer
	Start() error
}

func newBlocNodeCLICmd(p peer) *cobra.Command {
	mainCmd := &cobra.Command{
		Use: "blocnode",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	mainCmd.AddCommand(ledgercmd.Cmd(p))

	return mainCmd
}

func main() {
	// make it configurable
	p := fabric.NewPeer()
	if newBlocNodeCLICmd(p).Execute() != nil {
		os.Exit(1)
	}
}
