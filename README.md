# kitex-grpc mono repo 微服务

采用kitex的grpc/proto作为微服务技术选型，服务治理全部交给cncf的istio
所有grpc服务都此仓库，
proto接口文件在独立仓库，

grpc服务的pb生成使用buf gen 来读取本仓库根目proto目录下的proto子仓库内容

## 部署拓朴
todo

## mono repo结构

```shell
├── app                       #app目录定义为 微服务目录：grpc-开头 为grpc服务，http-开头 为restful服务
│   ├── grpc-account              #grpc-account服务：账户服务
│   │   ├── config.yaml
│   │   ├── main.go
│   │   ├── pkg
│   │   │   └── jwt.go
│   │   ├── rbac_model.conf
│   │   ├── readme.md
│   │   └── service
│   │       ├── const_def.go
│   │       ├── errno.go
│   │       ├── feishu.go
│   │       ├── feishu_test.go
│   │       ├── impl_health.go
│   │       ├── impl_login.go
│   │       └── service.go
│   ├── grpc-authx                #grpc-authx服务：私有api统一鉴权服务，sidecar 网关模式部署
│   │   ├── config.yaml
│   │   ├── main.go
│   │   ├── rbac_model.conf
│   │   ├── readme.md
│   │   └── service
│   │       ├── authn.go
│   │       ├── authz.go
│   │       ├── const_def.go
│   │       ├── envoy_check.go
│   │       ├── errno.go
│   │       └── service.go
│   └── grpc-release             #grpc-release服务：gitops发布服务
│       ├── config.yml
│       ├── main.go
│       ├── readme.md
│       └── service
│           ├── const_def.go
│           ├── errno.go
│           ├── health.go
│           ├── jobci.go
│           ├── service.go
│           └── webops.go
├── pkg                  #pkg目录为公共pkg包：casbin鉴权，viper配置，db连接，event 事件消息，logger日志，respstatus响应状态码，shutdown延迟优雅停服...
│   ├── authz
│   │   └── enforcer.go
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── postsql.go
│   ├── event
│   │   └── nats.go
│   ├── logger
│   │   └── logger.go
│   ├── respstatus
│   │   ├── err.go
│   │   └── resp.go
│   ├── shutdown
│   │   └── graceful.go
│   └── utils
│       ├── async.go
│       ├── cast
│       │   └── cast.go
├── proto                        # proto目录为proto3 idl gitsubmodle reop (http/grpc接口 集中定义管理)。除了_common和google公共目录，其他目录都为对应服务的proto3目录
│   ├── _common                     #_common目录为公共的proto message定义 
│   │   ├── gotag
│   │   │   └── options.proto
│   │   ├── jwt
│   │   │   └── wechat.proto
│   │   ├── respstatus
│   │   │   └── resp.proto
│   │   └── validtor
│   │       ├── fix_length.txt
│   │       └── validtor.proto
│   ├── account                    #account目录为 `账户服务` proto3定义（基于proto功能拆分3种proto文件）
│   │   ├── service.proto             #service.proto文件：grpc和http 接口签名定义
│   │   ├── t_info.proto              #t_info.proto文件：数据库表结构定义
│   │   └── vo.proto                  #vo.proto文件：服务里message内嵌值定义
│   ├── authx                      #authx目录为 `鉴权服务` proto3定义（基于proto功能拆分3种proto文件）
│   │   ├── service.proto
│   │   ├── t_info.proto
│   │   └── vo.proto
│   ├── google                     #googlem目录为google api 相关proto 结构定义
│   │   ├── api
│   │   │   ├── annotations.proto
│   │   │   ├── client.proto
│   │   │   ├── field_behavior.proto
│   │   │   ├── http.proto
│   │   │   ├── httpbody.proto
│   │   │   ├── resource.proto
│   │   │   └── routing.proto
│   │   └── protobuf
│   │       ├── any.proto
│   │       ├── api.proto
│   │       ├── compiler
│   │       ├── descriptor.proto
│   │       ├── duration.proto
│   │       ├── empty.proto
│   │       ├── field_mask.proto
│   │       ├── source_context.proto
│   │       ├── struct.proto
│   │       ├── timestamp.proto
│   │       ├── type.proto
│   │       └── wrappers.proto
│   └── release                  #release目录为 `发布服务` proto3定义（基于proto功能拆分3种proto文件）
│       ├── service.proto
│       ├── t_info.proto
│       └── vo.proto
└── script                       #环境脚本工具
|    ├── install_tools.sh      
|    └── start.sh
├── buf.gen.tag.yaml
├── buf.gen.validator.yaml
├── buf.gen.yaml
├── buf.work.yaml
├── buf.yaml
├── Dockerfile                 #二阶段容器镜像构建
├── go.mod
├── go.sum
├── Makefile                   #构建工具集合
├── README.md
```

TODO
- [] proto目录repo里 自动化gitsubmodle同名分支切换
- [] 补齐基于istio网关部署拓朴
