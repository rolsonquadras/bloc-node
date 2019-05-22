#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

ARG GO_VER
ARG ALPINE_VER
ARG FABRIC_PEER_EXT_IMAGE
ARG FABRIC_PEER_EXT_TAG

FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
RUN apk add --no-cache \
	gcc \
	musl-dev \
	git \
	libtool \
	bash \
	make;
ADD . $GOPATH/src/trustbloc/bloc-node
WORKDIR $GOPATH/src/trustbloc/bloc-node
ARG GO_TAGS
ARG GOPROXY
RUN make GO_TAGS=${GO_TAGS} GOPROXY=${GOPROXY} blocnode-cli



FROM ${FABRIC_PEER_EXT_IMAGE}:${FABRIC_PEER_EXT_TAG}
COPY --from=golang /go/src/trustbloc/bloc-node/.build/bin /usr/local/bin
CMD ["blocnode","ledger","start"]