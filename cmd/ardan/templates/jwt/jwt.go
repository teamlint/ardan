package jwt

import (
	"borderless/pom/app/model"
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/teamlint/ardan/config"
)

// NewToken 生成JWT Token
func NewToken(userID string, loginType model.LoginType) *jwtlib.Token {

	token := jwtlib.New(jwtlib.GetSigningMethod(config.Get("JWT", "Method").String("HS256")))
	// claims
	claims := jwtlib.MapClaims{
		"id":   userID,
		"type": loginType,
	}
	exp := config.Get("JWT", "Exp").String("72h")
	dur, err := time.ParseDuration(exp)
	if err == nil {
		claims["exp"] = time.Now().Add(dur).Unix()
	}
	token.Claims = claims
	return token
}

// SignedString 获取Token签名字符串
func SignedString(token *jwtlib.Token) string {
	tokenString, err := token.SignedString([]byte(config.Get("JWT", "Secret").String("")))
	if err != nil {
		return ""
	}
	return tokenString
}
