package service

import (
	"context"

	"ollie/app/grpc-account/pkg"
	"ollie/kitex_gen/account"
	"ollie/pkg/respstatus"
)

// LoginQR 开放接口
func (h *ServiceImpl) LoginQR(ctx context.Context, req *account.GetLoginQRRequest) (resp *account.GetLoginQRResponse, err error) {
	resp = new(account.GetLoginQRResponse)
	accessToken, err := getAccessTokenByCode(req.Code, req.RedirectUrl)
	if err != nil {
		resp.Status = respstatus.ErrStatusResp(AccountGetFeiShuCodeErr)
		return
	}
	info, err := getUserinfo(accessToken.AccessToken)
	if err != nil {
		resp.Status = respstatus.ErrStatusResp(AccountGetFeiShuUserInfoErr)
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
	resp.Status = respstatus.ErrStatusResp(nil)
	return
}

// RefreshJwtToken 开放接口
func (h *ServiceImpl) RefreshJwtToken(ctx context.Context, req *account.RefreshJwtTokenRequest) (resp *account.RefreshJwtTokenResponse, err error) {
	resp = new(account.RefreshJwtTokenResponse)
	jwtPayload, err := pkg.RefreshJwtToken(req.RefleshToken)
	if err != nil {
		resp.Status = respstatus.ErrStatusResp(AccountRefreshTokenErr)
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
	resp.Status = respstatus.ErrStatusResp(nil)
	return
}
