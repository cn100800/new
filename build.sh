#!/bin/bash

go install

if [[ -z $GOPATH ]]; then
    echo "GOPATH is not exists"
    exit
fi

${GOPATH}/bin/new

VERSION=$(date +%Y.%m.%d)".release"

OLD_VERSION=$(git tag |sort -Vr |head -1)

sed -i '' "s/${VERSION}/${OLD_VERSION}/g" conf.go

git tag $VERSION

git push origin $VERSION
