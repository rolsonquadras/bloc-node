/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockpeer struct {
	err error
}

func (m *mockpeer) Start() error {
	return m.err
}

func TestBlocNodeCMD(t *testing.T) {
	mockpeer := &mockpeer{}
	cmd := newBlocNodeCLICmd(mockpeer)
	cmd.SetArgs([]string{})
	require.NoError(t, cmd.Execute())
}
