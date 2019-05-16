#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Supported Targets:
#
# all:          runs unit and integration tests
# lint:         runs linters
# checks:       runs code checks
# unit-test:    runs unit tests
#

export GO111MODULE = on

# Local variables used by makefile
DEV_IMAGES = $(shell docker images dev-* -q)

checks: version license lint

lint:
	@scripts/check_lint.sh

license: version
	@scripts/check_license.sh

all: checks unit-test

unit-test: checks docker-thirdparty
	@scripts/unit.sh

version:
	@scripts/check_version.sh

docker-thirdparty:
	docker pull couchdb:2.2.0

clean-images:
	@echo "Stopping all containers, pruning containers and images, deleting dev images"
ifneq ($(strip $(CONTAINER_IDS)),)
	@docker stop $(CONTAINER_IDS)
endif
	@docker system prune -f
ifneq ($(strip $(DEV_IMAGES)),)
	@docker rmi $(DEV_IMAGES) -f
endif

.PHONY: checks lint license all unit-test version clean-images docker-thirdparty