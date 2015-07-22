#!/bin/sh

ORG_PATH="geoip"
REPO_PATH="${ORG_PATH}/geoip"

export GOPATH=${PWD}/gopath

rm -rf $GOPATH/src/${REPO_PATH}
mkdir -p $GOPATH/src/${ORG_PATH}
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)
echo $(go env)

go build -o bin/geoip ${REPO_PATH}
