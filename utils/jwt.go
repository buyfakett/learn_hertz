package utils

import (
	"fmt"
	"hertz_demo/utils/config"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Cfg.Jwt.Secret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Cfg.Jwt.ExpireTime) * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Cfg.Server.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 类型断言 + 校验 token 有效性
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		// 检查 issuer 是否匹配
		if claims.Issuer != config.Cfg.Server.Name {
			return nil, fmt.Errorf("issuer 不匹配")
		}

		// 检查是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("token 已过期")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("token 不合法")
}

func GetUsernameFromContext(c *app.RequestContext) (string, error) {
	claimsInterface, exists := c.Get("claims")
	if !exists {
		return "", fmt.Errorf("未找到用户信息")
	}

	claims, ok := claimsInterface.(*Claims)
	if !ok {
		return "", fmt.Errorf("用户信息格式错误")
	}

	return claims.Username, nil
}
