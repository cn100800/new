#!/bin/bash

# go install

# if [[ -z $GOPATH ]]; then
#     echo "GOPATH is not exists"
#     exit
# fi

# ${GOPATH}/bin/new

SUFFIX=".release"

OUTPUT_DIR="output"

RELEASE_DIR="release"

VERSION=$(date +%Y.%m.%d)${SUFFIX}

OLD_VERSION=$(git tag |sort -Vr |head -1)

sed -i '' "s/${VERSION}/${OLD_VERSION}/g" conf.go

make build

tar -cvzf $RELEASE_DIR/$VERSION.tar.gz -C $OUTPUT_DIR .

git tag $VERSION

git push origin $VERSION
