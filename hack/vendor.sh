#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

clone git github.com/docker/go-plugins-helpers c673bfc017b81d1dd62b922fdc8672a42dd28226
clone git github.com/docker/go-connections 64a666f8c9ca9539fe05fc36fab9ccc257bcbc55
clone git github.com/Sirupsen/logrus v0.8.7 # logrus is a common dependency among multiple deps
clone git github.com/opencontainers/runc 3d8a20bb772defc28c355534d83486416d1719b4

clean

mv vendor/src/* vendor/
