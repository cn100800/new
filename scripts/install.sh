#!/bin/bash set -o errexit
# go mod vendor
wget -cq $(curl -sSf https://api.github.com/repos/cn100800/news/releases | grep release.tar.xz | awk 'NR==2 {printf $2"\n"}' | tr -d \") -O - | sudo tar -xjC /usr/local/bin/
