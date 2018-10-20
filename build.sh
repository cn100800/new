#!/bin/bash
go install
if [[ -z $GOPATH ]]; then
    echo "GOPATH is not exists"
    exit
fi
${GOPATH}/bin/new
