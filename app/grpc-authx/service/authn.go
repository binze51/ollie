package service

// JWT Parse处理登录请求主体的身份合法性验证，供接下来使用casbin处理请求的验证

// 权限设计部分：身份验证(jwt) 和 请求鉴权(casbin)

import (
	"encoding/base64"
	"fmt"

	"ollie/pkg/utils"

	jwtV5 "github.com/golang-jwt/jwt/v5"
)

const jwtKey = "rXUAVdbdE0xthK74R1jbwcHghLdXlQUP"

type JwtClaims struct {
	UID      string `json:"uid"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	jwtV5.RegisteredClaims
}

func (s *ServiceImpl) ParseToken(token string) (*JwtClaims, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	decryptedData, err := utils.Decrypt([]byte(jwtKey), encryptedData)
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwtV5.ParseWithClaims(string(decryptedData), &JwtClaims{}, func(t *jwtV5.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, fmt.Errorf("jwt token is not valid")
	}

	claims, ok := jwtToken.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse jwt token payload: %w", err)
	}
	return claims, nil
}
