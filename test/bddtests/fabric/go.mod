// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/bloc-node/test/bddtests/fabric

replace github.com/hyperledger/fabric => github.com/trustbloc/fabric-mod v0.0.0-20190605152521-6547615cb978

replace github.com/hyperledger/fabric/extensions => github.com/trustbloc/fabric-mod/extensions v0.0.0-20190605152521-6547615cb978

replace github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric => github.com/trustbloc/fabric-sdk-go-ext/fabric v0.0.0-20190528182243-b95c24511993

require (
	github.com/cucumber/godog v0.8.1
	github.com/hyperledger/fabric v2.0.0-alpha+incompatible
	github.com/spf13/viper v1.0.2
	github.com/trustbloc/fabric-peer-test-common v0.1.2-0.20200213155832-06af5163b73f
)

go 1.13
