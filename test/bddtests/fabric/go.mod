// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/bloc-node/test/bddtests/fabric

replace github.com/hyperledger/fabric => github.com/trustbloc/fabric-mod v0.0.0-20190605152521-6547615cb978

replace github.com/hyperledger/fabric/extensions => github.com/trustbloc/fabric-mod/extensions v0.0.0-20190605152521-6547615cb978

replace github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric => github.com/trustbloc/fabric-sdk-go-ext/fabric v0.0.0-20190528182243-b95c24511993

require (
	github.com/DATA-DOG/godog v0.7.13
	github.com/hyperledger/fabric v2.0.0-alpha+incompatible
	github.com/spf13/viper v1.0.2
	github.com/trustbloc/fabric-peer-test-common v0.0.0-20190530161629-b93925e57103
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 // indirect
)
