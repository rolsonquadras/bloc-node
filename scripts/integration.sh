#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
set -e

echo "Running bloc-node integration tests..."
cd test/bddtests/fabric
go test -count=1 -v -cover . -p 1 -timeout=20m
cd $PWD