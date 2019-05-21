/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabric

import "github.com/hyperledger/fabric/peer/node"

// Peer fabric peer
type Peer interface {
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
