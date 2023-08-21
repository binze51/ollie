.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

## 服务生成命令 ##
gen:  ## [生成proto指定服务代码], example: `make gen service=release`
	# buf breaking --against '.git#branch=master,subdir=proto'
	@kitex -module ollie -type protobuf -I proto/  ./proto/$(service)/service.proto
	@buf generate --template buf.gen.tag.yaml --path proto/$(service)
	@buf generate --template buf.gen.validator.yaml --path proto/$(service)
	# @go mod tidy

build_all_pb: ## [生成包含所有服务的pb自描述文件], example: `make build_all_pb`
	buf build -o all.pb

list_grpc_service: ## [列出本项目所有grpc服务信息], example: `make list_grpc_service`
	@buf build -o all.pb
	grpcurl -protoset all.pb list

## 本地操作proto repo submodle指令 ##
## submodle 操作可以是同名分支，默认是main分支

addproto: ## [保存本地proto修改], example: `make addproto`
	@git subtree add --prefix=proto git@gitlab.com:back_end9494529/proto.git master

pullproto: ## [更新远端proto到本地], example: `make pullproto`
	@git subtree pull --prefix=proto git@gitlab.com:back_end9494529/proto.git master

pushproto: ## [提交本地proto到远端], example: `make pushproto`
	@git subtree push --prefix=proto git@gitlab.com:back_end9494529/proto.git master 


## 本地构建容器镜像 ##

VERSION := $(shell git describe --tags --always)
# VERSION:=v1.1.1
GIT_COMMIT := $(shell git rev-parse --short HEAD)
REGISTRY := <你的imagehub地址>
### image build push ###
imagebuildpublish: ## [PODMAN BUIDL AND PUSH] ,example: `make imagebuildpublish service=authx`
	@echo 'publish $(VERSION) to $(REGISTRY)'
	@podman build -f Dockerfile --no-cache --build-arg SERVICE=$(service) --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(REGISTRY)/ollie-$(service):$(VERSION) .
	@podman push $(REGISTRY)/ollie-$(service):$(VERSION)