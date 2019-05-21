#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Supported Targets:
#
# all:                        runs unit and integration tests
# clean:                      cleans the build area
# lint:                       runs linters
# checks:                     runs code checks
# unit-test:                  runs unit tests
# populate-fixtures:          populate crypto directory and channel configuration for bddtests
# crypto-gen:                 generates crypto directory
# channel-config-gen:         generates test channel configuration transactions and blocks
# bddtests:                   run bddtests
# blocnode-cli:               generate blocnode cli binary
# blocnode-docker:            generate blocnode image
#

ARCH=$(shell go env GOARCH)
GO_VER = $(shell grep "GO_VER" .ci-properties |cut -d'=' -f2-)
ALPINE_VER ?= 3.9
GO_TAGS ?=
export GO111MODULE = on

# Tool commands (overridable)
DOCKER_CMD ?= docker

# Local variables used by makefile
PROJECT_NAME   = bloc-node
CONTAINER_IDS  = $(shell docker ps -a -q)
DEV_IMAGES     = $(shell docker images dev-* -q)

# Fabric tools docker image (overridable)
FABRIC_TOOLS_IMAGE   ?= hyperledger/fabric-tools
FABRIC_TOOLS_VERSION ?= 2.0.0-alpha
FABRIC_TOOLS_TAG     ?= $(ARCH)-$(FABRIC_TOOLS_VERSION)

# Namespace for the blocnode image
DOCKER_OUTPUT_NS     ?= trustbloc
BLOC_NODE_IMAGE_NAME ?= bloc-node

# Fabric peer ext docker image (overridable)
FABRIC_PEER_EXT_IMAGE   ?= trustbloc/fabric-peer
FABRIC_PEER_EXT_VERSION ?= 0.1.0-snapshot-5e03e3c
FABRIC_PEER_EXT_TAG     ?= $(ARCH)-$(FABRIC_PEER_EXT_VERSION)


checks: version license lint

lint:
	@scripts/check_lint.sh

license: version
	@scripts/check_license.sh

all: checks unit-test bddtests

unit-test: checks docker-thirdparty
	@scripts/unit.sh

bddtests: clean checks populate-fixtures docker-thirdparty blocnode-docker
	@scripts/integration.sh

blocnode-cli: clean
	@echo "Building blocnode cli"
	@mkdir -p ./build/bin
	@go build -o .build/bin/blocnode github.com/trustbloc/bloc-node/cmd/blocnode

blocnode-docker:
	@docker build -f ./images/blocnode/fabric/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(BLOC_NODE_IMAGE_NAME):latest \
	--build-arg FABRIC_PEER_EXT_IMAGE=$(FABRIC_PEER_EXT_IMAGE) \
	--build-arg FABRIC_PEER_EXT_TAG=$(FABRIC_PEER_EXT_TAG) \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) \
	--build-arg GO_TAGS=$(GO_TAGS) \
	--build-arg GOPROXY=$(GOPROXY) .

crypto-gen:
	@echo "Generating crypto directory ..."
	@$(DOCKER_CMD) run -i \
		-v /$(abspath .):/opt/workspace/$(PROJECT_NAME) -u $(shell id -u):$(shell id -g) \
		$(FABRIC_TOOLS_IMAGE):$(FABRIC_TOOLS_TAG) \
		//bin/bash -c "FABRIC_VERSION_DIR=fabric /opt/workspace/${PROJECT_NAME}/scripts/generate_crypto.sh"

channel-config-gen:
	@echo "Generating test channel configuration transactions and blocks ..."
	@$(DOCKER_CMD) run -i \
		-v /$(abspath .):/opt/workspace/$(PROJECT_NAME) -u $(shell id -u):$(shell id -g) \
		$(FABRIC_TOOLS_IMAGE):$(FABRIC_TOOLS_TAG) \
		//bin/bash -c "FABRIC_VERSION_DIR=fabric/ /opt/workspace/${PROJECT_NAME}/scripts/generate_channeltx.sh"

populate-fixtures:
	@scripts/populate-fixtures.sh

version:
	@scripts/check_version.sh

docker-thirdparty:
	docker pull couchdb:2.2.0
	docker pull hyperledger/fabric-orderer:$(ARCH)-2.0.0-alpha

clean-images:
	@echo "Stopping all containers, pruning containers and images, deleting dev images"
ifneq ($(strip $(CONTAINER_IDS)),)
	@docker stop $(CONTAINER_IDS)
endif
	@docker system prune -f
ifneq ($(strip $(DEV_IMAGES)),)
	@docker rmi $(DEV_IMAGES) -f
endif

clean:
	rm -Rf ./test/bddtests/fabric/docker-compose.log
	rm -Rf ./test/bddtests/fabric/fixtures/fabric/channel
	rm -Rf ./test/bddtests/fabric/fixtures/fabric/crypto-config
	rm -Rf ./.build

.PHONY: clean checks lint license all unit-test version clean-images docker-thirdparty crypto-gen channel-config-gen populate-fixtures bddtests