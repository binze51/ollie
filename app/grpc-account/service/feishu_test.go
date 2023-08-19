package service

import (
	"fmt"
	"testing"
)

func TestFeishu(t *testing.T) {
	// https://open.feishu.cn/document/common-capabilities/sso/web-application-sso/web-app-overview
	toekn, err := getAccessTokenByCode("428v7b6c4caa48dcaba564fb1157f4f6", "http://172.25.4.241:4000/login")
	if err != nil {
		panic(err)
	}
	fmt.Println(toekn.AccessToken)
	info, err := getUserinfo(toekn.AccessToken)
	if err != nil {
		panic(err)
	}

	fmt.Println(*info)
}
