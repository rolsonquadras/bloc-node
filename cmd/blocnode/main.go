/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/trustbloc/bloc-node/cmd/blocnode/fabpeer"
)

type nodeCLI interface {
	ExecuteCMD() error
}

func main() {
	var cli nodeCLI
	var err error
	cli, err = fabpeer.NewFabPeerCLI()
	if err != nil {
		panic(err)
	}

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if cli.ExecuteCMD() != nil {
		os.Exit(1)
	}
}
