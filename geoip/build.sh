#!/bin/sh

REPO_PATH="geoip"

export GOPATH=${PWD}/gopath

rm -rf $GOPATH
mkdir -p $GOPATH/src
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)
echo $(go env)

go build -o bin/geoip ${REPO_PATH}
rm -rf $GOPATH
