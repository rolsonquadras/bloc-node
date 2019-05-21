#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
set -e



# Packages to exclude
PKGS=`go list github.com/trustbloc/bloc-node/... 2> /dev/null | \
                                                 grep -v /mocks | \
                                                 grep -v /api | \
                                                 grep -v /protos`
echo "Running unit tests..."
FABRIC_SAMPLECONFIG_PATH="src/github.com/trustbloc/bloc-node/pkg/ledger/fabric/testutil/sampleconfig" go test $PKGS -count=1 -race -coverprofile=coverage.txt -covermode=atomic  -p 1 -timeout=10m