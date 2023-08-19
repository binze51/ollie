package service

import "ollie/pkg/respstatus"

const (

	// 账户中心 错误码2w+
	ACCOUNT_ERROR_NO = 20000
)

var (
	LoginErr = respstatus.NewErrNo(ACCOUNT_ERROR_NO, "Wrong username or password")

	AccountNotExistErr     = respstatus.NewErrNo(ACCOUNT_ERROR_NO+1, "User does not exists")
	AccountAlreadyExistErr = respstatus.NewErrNo(ACCOUNT_ERROR_NO+2, "User already exists")
)
