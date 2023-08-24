package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type tokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint64 `json:"expires_in"`
}

type userInfo struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	OpenID  string `json:"open_id"`
	UnionID string `json:"union_id"`
	Mobile  string `json:"mobile"`
}

func getUserinfo(feishuAccessTokengo string) (*userInfo, error) {
	url := "https://passport.feishu.cn/suite/passport/oauth/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+feishuAccessTokengo)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := new(userInfo)
	err = json.Unmarshal(body, user)
	if user.Name == "" {
		err = errors.New(string(body))
		return nil, err
	}
	return user, err
}

func getAccessTokenByCode(code, redirectURL string) (*tokenInfo, error) {
	url := "https://passport.feishu.cn/suite/passport/oauth/token"
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
			viper.GetString("app.feishu.app_id"), viper.GetString("app.feishu.app_secret"), code, redirectURL)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	token := new(tokenInfo)
	err = json.Unmarshal(body, token)
	if token.AccessToken == "" {
		err = errors.New(string(body))
		return nil, err
	}
	return token, err
}
