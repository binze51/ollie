package service

import "ollie/pkg/respstatus"

const (

	// 账户中心 错误码4w+
	RELEASE_ERROR_NO = 40000
)

var ReleaseErr = respstatus.NewErrNo(RELEASE_ERROR_NO, "Wrong username or password")
