/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabric

import (
	"io/ioutil"
	"os"
	"testing"

	xtestutil "github.com/hyperledger/fabric/extensions/testutil"
	"github.com/hyperledger/fabric/msp/mgmt/testtools"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	//setup extension test environment
	_, _, destroy := xtestutil.SetupExtTestEnv()

	code := m.Run()
	destroy()
	os.Exit(code)

}

func TestPeerStart(t *testing.T) {
	defer viper.Reset()
	g := NewGomegaWithT(t)

	tempDir, err := ioutil.TempDir("", "peerstart")
	g.Expect(err).NotTo(HaveOccurred())
	defer func() { require.NoError(t, os.RemoveAll(tempDir)) }()

	viper.Set("peer.address", "localhost:6051")
	viper.Set("peer.listenAddress", "0.0.0.0:6051")
	viper.Set("peer.chaincodeListenAddress", "0.0.0.0:6052")
	viper.Set("peer.fileSystemPath", tempDir)
	viper.Set("chaincode.executetimeout", "30s")
	viper.Set("chaincode.mode", "dev")

	require.NoError(t, msptesttools.LoadMSPSetupForTesting())

	go func() {
		fabpeer := NewPeer()
		require.NoError(t, fabpeer.Start(), "expected to successfully start peer")
	}()

	grpcProbe := func(addr string) bool {
		c, err := grpc.Dial(addr, grpc.WithBlock(), grpc.WithInsecure())
		if err == nil {
			require.NoError(t, c.Close())
			return true
		}
		return false
	}
	g.Eventually(grpcProbe("localhost:6051")).Should(BeTrue())
}
