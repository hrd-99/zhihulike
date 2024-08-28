package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenOptions 结构体，用于存储生成 token 的选项
type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	// Token 结构体，用于存储生成的 token
	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

// BuildTokens 函数，用于生成 token
func BuildTokens(opt TokenOptions) (Token, error) {
	var token Token
	// 获取当前时间，并减去一分钟
	now := time.Now().Add(-time.Minute).Unix()
	// 生成 access token
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return token, err
	}
	// 将生成的 token 和过期时间存储到 token 结构体中
	token.AccessToken = accessToken
	token.AccessExpire = now + opt.AccessExpire

	return token, nil
}

// genToken 函数，用于生成 token
func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	// 创建一个 MapClaims 类型的变量，用于存储 token 的 payload
	claims := make(jwt.MapClaims)
	// 设置 token 的过期时间
	claims["exp"] = iat + seconds
	// 设置 token 的签发时间
	claims["iat"] = iat
	// 将 payloads 中的数据存储到 token 的 payload 中
	for k, v := range payloads {
		claims[k] = v
	}
	// 创建一个 token，并设置签名方法为 HS256
	token := jwt.New(jwt.SigningMethodHS256)
	// 将 payload 存储到 token 中
	token.Claims = claims

	// 使用 secretKey 对 token 进行签名，并返回签名后的 token
	return token.SignedString([]byte(secretKey))
}
