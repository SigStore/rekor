#!/bin/bash
set -ex
testdir=$(dirname "$0")

docker-compose up -d

# Node
nodedir=${testdir}/node
go run ./cmd/cli/ upload \
    --artifact $(< ${nodedir}/url) --sha $(< ${nodedir}/sha) \
    --signature=${nodedir}/sig  --public-key=${nodedir}/key