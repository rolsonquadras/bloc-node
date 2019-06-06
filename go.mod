// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/bloc-node

require (
	github.com/hyperledger/fabric v2.0.0-alpha+incompatible
	github.com/hyperledger/fabric/extensions v0.0.0
	github.com/onsi/gomega v1.5.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v0.0.0-20150908122457-1967d93db724
	github.com/stretchr/testify v1.3.0
	google.golang.org/grpc v1.20.1
)

replace github.com/hyperledger/fabric => github.com/trustbloc/fabric-mod v0.0.0-20190605152521-6547615cb978

replace github.com/hyperledger/fabric/extensions => github.com/trustbloc/fabric-peer-ext/mod/peer v0.0.0-20190606010932-60cd4152d74a

replace github.com/trustbloc/fabric-peer-ext => github.com/trustbloc/fabric-peer-ext v0.0.0-20190606010932-60cd4152d74a
