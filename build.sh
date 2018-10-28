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

VERSION=$(date +%Y.%m.%d)${SUFFIX}
OLD_VERSION=$(git tag |sort -Vr |head -1)

echo "===>当前版本是:"$OLD_VERSION

sed -i '' "s/${OLD_VERSION}/${VERSION}/g" ${ETC_DIR}/conf.go

echo "===>新版本是:"$VERSION

make build

mkdir -p $RELEASE_DIR
echo "===>开发环境打包:"
tar -cvzf $RELEASE_DIR/$VERSION.tar.gz -C $OUTPUT_DIR news
echo "===>最终发布包"
tar -cvzf $RELEASE_DIR/$VERSION"_linux".tar.gz -C $OUTPUT_DIR news_linux

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
