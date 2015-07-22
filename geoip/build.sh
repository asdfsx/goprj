#!/bin/sh

REPO_PATH="geoip"

export GOPATH=${PWD}/gopath

rm -rf $GOPATH
mkdir -p $GOPATH/src
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)
echo $(go env)

go build -o bin/geoipserver ${REPO_PATH}
go build -o bin/geoipclient ${REPO_PATH}/geoclient
rm -rf $GOPATH
