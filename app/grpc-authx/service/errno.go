package service

import "ollie/pkg/respstatus"

const (

	// 鉴权中心 错误码3w+
	AURHX_ERROR_NO = 30000
)

var authxErrErr = respstatus.NewErrNo(AURHX_ERROR_NO, "Wrong username or password")
