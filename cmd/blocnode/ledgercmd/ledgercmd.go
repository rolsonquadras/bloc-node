/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ledgercmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trustbloc/bloc-node/cmd/blocnode/ledgercmd/startcmd"
)

const (
	ledgerFuncName = "ledger"
	ledgerCmdDes   = "Operate a ledger node: start."
)

// peer
type peer interface {
	// Init peer
	Init()
	// Start peer
	Start() error
}

// Cmd returns the cobra command for ledger
func Cmd(p peer) *cobra.Command {
	nodeCmd := &cobra.Command{
		Use:   ledgerFuncName,
		Short: fmt.Sprint(ledgerCmdDes),
		Long:  fmt.Sprint(ledgerCmdDes),
	}
	nodeCmd.AddCommand(startcmd.Cmd(p))
	return nodeCmd
}
