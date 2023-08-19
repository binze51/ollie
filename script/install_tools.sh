#!/bin/bash

if ! command -v buf &> /dev/null
then
    go install github.com/bufbuild/buf/cmd/buf@v1.17.0
fi

if ! command -v protoc-gen-go &> /dev/null
then
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
fi

if ! command -v protoc-gen-gotag &> /dev/null
then
    go install github.com/srikrsna/protoc-gen-gotag@v0.6.2
fi

if ! command -v kitex &> /dev/null
then
    go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.6.1
fi