#!/bin/bash

# go install

# if [[ -z $GOPATH ]]; then
#     echo "GOPATH is not exists"
#     exit
# fi

# ${GOPATH}/bin/news

SUFFIX=".release"

OUTPUT_DIR="output"

RELEASE_DIR="release"

ETC_DIR="etc"

GZ="gz"
XZ="xz"

VERSION=$(date +%Y.%m.%d)${SUFFIX}
OLD_VERSION=$(git tag |sort -Vr |head -1)

echo "===>当前版本是:"$OLD_VERSION

sed -i '' "s/${OLD_VERSION}/${VERSION}/g" ${ETC_DIR}/conf.go

echo "===>新版本是:"$VERSION

make clean

make build

make build-linux

mkdir -p $RELEASE_DIR
rm -vrf $RELEASE_DIR/*
echo "===>开发环境打包:"
tar -cvzf $RELEASE_DIR/$VERSION.tar.gz -C $OUTPUT_DIR/$GZ .
shasum -a 256 $RELEASE_DIR/$VERSION.tar.gz
echo "===>最终发布包"
tar -cvjf $RELEASE_DIR/$VERSION.tar.xz -C $OUTPUT_DIR/$XZ .

if [[ $1 == "" ]]; then
    read -p "输入提交内容" message
    if [[ $message == "" ]]; then
        exit
    fi
else
    message=$1
fi

git add .
git commit -m $message
git push
git tag $VERSION
git push origin $VERSION
