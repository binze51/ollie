package pkg

// JWT Parse处理登录请求主体的身份合法性验证，供接下来使用casbin处理请求的验证

// 权限设计部分：身份验证(jwt) 和 请求鉴权(casbin)

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	account "ollie/kitex_gen/account"

	jwtV5 "github.com/golang-jwt/jwt/v5"
)

const jwtKey = "rXUAVdbdE0xthK74R1jbwcHghLdXlQUP"

type JwtClaims struct {
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	jwtV5.RegisteredClaims
}

func CreateJwtToken(claims *JwtClaims) (*account.LoginInfo, error) {
	// 请求token生成
	expiresAt := time.Now().Add(1 * time.Hour)
	claims.RegisteredClaims.ExpiresAt = jwtV5.NewNumericDate(expiresAt)
	jwtToken := jwtV5.NewWithClaims(jwtV5.SigningMethodHS512, claims)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}
	jwtKeyByte := []byte(jwtKey)
	tokenBtye := []byte(token)
	encryptedData, err := encrypt(jwtKeyByte, tokenBtye)
	if err != nil {
		return nil, err
	}
	encryptedDataString := base64.StdEncoding.EncodeToString(encryptedData)

	// 刷新token生成
	refleshExpiresAt := expiresAt.Add(2 * time.Hour)
	refleshJwtClaims := new(JwtClaims)
	refleshJwtClaims.RegisteredClaims.ExpiresAt = jwtV5.NewNumericDate(refleshExpiresAt)
	refleshjwtToken := jwtV5.NewWithClaims(jwtV5.SigningMethodHS512, refleshJwtClaims)
	refleshtoken, err := refleshjwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}
	refleshtokenBtye := []byte(refleshtoken)
	refleshencryptedData, err := encrypt(jwtKeyByte, refleshtokenBtye)
	if err != nil {
		return nil, err
	}
	refleshencryptedDataString := base64.StdEncoding.EncodeToString(refleshencryptedData)

	NewJwtPayload := &account.LoginInfo{
		UserName: claims.UserName,
		Phone:    claims.Phone,

		ReqToken:  encryptedDataString,
		ExpiresAt: uint64(expiresAt.Unix()),

		RefleshToken:     refleshencryptedDataString,
		RefleshExpiresAt: uint64(refleshExpiresAt.Unix()),
	}
	return NewJwtPayload, nil
}

func RefreshJwtToken(token string) (*account.LoginInfo, error) {
	RefreshJwtClaims, err := parseToken(token)
	if err != nil {
		return nil, err
	}
	// 创建新的ReqToken
	return CreateJwtToken(RefreshJwtClaims)
}

func parseToken(token string) (*JwtClaims, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	decryptedData, err := decrypt([]byte(jwtKey), encryptedData)
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

func encrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
