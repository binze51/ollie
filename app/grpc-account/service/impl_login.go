package service

import (
	"context"
	"fmt"

	"ollie/app/grpc-account/pkg"
	"ollie/kitex_gen/account"
	"ollie/pkg/respstatus"
)

// LoginQR 开放接口
func (h *ServiceImpl) LoginQR(ctx context.Context, req *account.GetLoginQRRequest) (resp *account.GetLoginQRResponse, err error) {
	toekn, err := getAccessTokenByCode(req.Code, req.RedirectUrl)
	if err != nil {
		resp.Status = respstatus.BuildStatusResp(AccountGetFeiShuCodeErr)
		return
	}
	fmt.Println(toekn.AccessToken)
	info, err := getUserinfo(toekn.AccessToken)
	if err != nil {
		resp.Status = respstatus.BuildStatusResp(AccountGetFeiShuUserInfoErr)
		return
	}
	jwtClaims := &pkg.JwtClaims{
		UserName: info.Name,
		Phone:    info.Mobile,
	}
	jwtPayload, err := pkg.CreateJwtToken(jwtClaims)
	resp.Info = &account.LoginInfo{
		UserName:  jwtPayload.UserName,
		Phone:     jwtPayload.Phone,
		ReqToken:  jwtPayload.ReqToken,
		ExpiresAt: jwtPayload.ExpiresAt,

		RefleshToken:     jwtPayload.RefleshToken,
		RefleshExpiresAt: jwtPayload.RefleshExpiresAt,
	}
	resp.Status = respstatus.BuildStatusResp(respstatus.Success)
	return
}

// RefreshJwtToken 开放接口
func (h *ServiceImpl) RefreshJwtToken(ctx context.Context, req *account.RefreshJwtTokenRequest) (resp *account.RefreshJwtTokenResponse, err error) {
	jwtPayload, err := pkg.RefreshJwtToken(req.RefleshToken)
	if err != nil {
		resp.Status = respstatus.BuildStatusResp(AccountRefreshTokenErr)
		return
	}
	resp.Info = &account.LoginInfo{
		UserName:  jwtPayload.UserName,
		Phone:     jwtPayload.Phone,
		ReqToken:  jwtPayload.ReqToken,
		ExpiresAt: jwtPayload.ExpiresAt,

		RefleshToken:     jwtPayload.RefleshToken,
		RefleshExpiresAt: jwtPayload.RefleshExpiresAt,
	}
	resp.Status = respstatus.BuildStatusResp(respstatus.Success)
	return
}
