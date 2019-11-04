# Copyright (c) 2019 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#!/bin/sh

set -e
set -x

# compile for all architectures
GOOS=linux   CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/amd64/vela   github.com/go-vela/cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/arm64/vela   github.com/go-vela/cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm   go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/arm/vela     github.com/go-vela/cli
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/windows/amd64/vela github.com/go-vela/cli
GOOS=darwin  CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/darwin/amd64/vela  github.com/go-vela/cli

# tar binary files prior to upload
tar -cvzf release/vela_linux_amd64.tar.gz   -C release/linux/amd64   vela
tar -cvzf release/vela_linux_arm64.tar.gz   -C release/linux/arm64   vela
tar -cvzf release/vela_linux_arm.tar.gz     -C release/linux/arm     vela
tar -cvzf release/vela_windows_amd64.tar.gz -C release/windows/amd64 vela
tar -cvzf release/vela_darwin_amd64.tar.gz  -C release/darwin/amd64  vela

# generate shas for tar files
sha256sum release/*.tar.gz > release/vela_checksums.txt
