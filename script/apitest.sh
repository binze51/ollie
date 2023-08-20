#!/bin/bash

#测试loginQR接口
grpcurl -plaintext -protoset all.pb -d @ localhost:7001 account.AccountService/LoginQR <<EOM
{
  "code": "428v7b6c4caa48dcaba564fb1157f4f6",
  "redirect_url": "http://172.25.4.241:4000/login"
}
EOM

