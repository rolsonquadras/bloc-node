/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package startcmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockpeer struct {
	err error
}

func (m *mockpeer) Start() error {
	return m.err
}

func (m *mockpeer) Init() {
}

func TestStart(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		mockpeer := &mockpeer{}
		cmd := Cmd(mockpeer)
		cmd.SetArgs([]string{})
		require.NoError(t, cmd.Execute())
	})

	t.Run("test failure", func(t *testing.T) {
		errmsg := "peer start failed"
		mockpeer := &mockpeer{err: errors.New(errmsg)}
		cmd := Cmd(mockpeer)
		cmd.SetArgs([]string{})
		err := cmd.Execute()
		require.NotNil(t, err)
		require.Equal(t, errmsg, err.Error())
	})

	t.Run("test send extra args", func(t *testing.T) {
		errmsg := "peer start failed"
		mockpeer := &mockpeer{err: errors.New(errmsg)}
		cmd := Cmd(mockpeer)
		cmd.SetArgs([]string{"a"})
		require.Error(t, cmd.Execute())
	})

}
