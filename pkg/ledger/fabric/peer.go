/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabric

import (
	"strings"

	"github.com/hyperledger/fabric/peer/node"
	"github.com/spf13/viper"
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
	viper.SetEnvPrefix(node.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	node.InitCmd(nil, nil)
}
