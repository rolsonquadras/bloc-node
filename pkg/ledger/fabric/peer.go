/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabric

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/hyperledger/fabric/peer/common"
	"github.com/hyperledger/fabric/peer/node"
)

// Peer fabric peer
type Peer interface {
	// Init peer
	Init()
	// Start peer
	Start() error
}

type peer struct {
}

// NewPeer return fabric peer
func NewPeer() Peer {
	return peer{}
}

func (l peer) Start() error {
	return node.Start()
}

func (l peer) Init() {
	// For environment variables.
	viper.SetEnvPrefix(common.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	common.InitCmd(nil, nil)
}
