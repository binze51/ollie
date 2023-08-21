#!/bin/bash

#测试loginQR接口
grpcurl -plaintext -protoset all.pb -d @ localhost:7001 account.AccountService/LoginQR <<EOM
{
  "code": "e5ct5de8e1b240708b78b11651281984",
  "redirect_url": "http://172.25.4.241:4000/login"
}
EOM

