#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

clone git github.com/docker/go-plugins-helpers 2bce64a6080c74e6eb388aab8da36067cd33ca08
clone git github.com/docker/go-connections 64a666f8c9ca9539fe05fc36fab9ccc257bcbc55
clone git github.com/Sirupsen/logrus v0.8.7 # logrus is a common dependency among multiple deps
clone git github.com/opencontainers/runc 3d8a20bb772defc28c355534d83486416d1719b4

clean

mv vendor/src/* vendor/
