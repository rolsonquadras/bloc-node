/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabpeer

import (
	"strings"

	"github.com/hyperledger/fabric/peer/chaincode"
	"github.com/hyperledger/fabric/peer/channel"
	"github.com/hyperledger/fabric/peer/clilogging"
	"github.com/hyperledger/fabric/peer/common"
	"github.com/hyperledger/fabric/peer/lifecycle"
	"github.com/hyperledger/fabric/peer/node"
	"github.com/hyperledger/fabric/peer/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FabCLI fabric cli
type FabCLI interface {
	ExecuteCMD() error
}

type cli struct {
	cmd *cobra.Command
}

// NewFabPeerCLI return fab peer cli
func NewFabPeerCLI() (FabCLI, error) {
	mainCmd := &cobra.Command{Use: "peer"}

	if err := setFlags(mainCmd); err != nil {
		return nil, err
	}

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(node.Cmd())
	mainCmd.AddCommand(chaincode.Cmd(nil))
	mainCmd.AddCommand(clilogging.Cmd(nil))
	mainCmd.AddCommand(channel.Cmd(nil))
	mainCmd.AddCommand(lifecycle.Cmd())

	return &cli{cmd: mainCmd}, nil
}

func setFlags(cmd *cobra.Command) error {
	// For environment variables.
	viper.SetEnvPrefix(common.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	mainFlags := cmd.PersistentFlags()

	mainFlags.String("logging-level", "", "Legacy logging level flag")
	if err := viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level")); err != nil {
		return err
	}
	if err := mainFlags.MarkHidden("logging-level"); err != nil {
		return err
	}
	return nil
}

func (c cli) ExecuteCMD() error {
	return c.cmd.Execute()
}
