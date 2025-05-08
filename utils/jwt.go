package utils

import (
	"fmt"
	"hertz_demo/utils/config"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Cfg.Jwt.Secret)

type Claims struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT Token
//
// 参数:
//   - username: 用户名
//   - id: 用户id
//   - expTime: 可选参数，Token 过期时间（单位：小时）, 如果不传，则使用配置文件中的默认值
//
// 返回:
//   - string: 生成的 Token
//   - error: 错误信息（如果有）
func GenerateToken(userid uint, username string, expTime ...int) (string, error) {
	nowTime := time.Now()
	var expireHours int

	// 如果 expTime 有传参，则用第一个值；否则用默认值
	if len(expTime) > 0 {
		expireHours = expTime[0]
	} else {
		expireHours = config.Cfg.Jwt.ExpireTime
	}

	expireTime := nowTime.Add(time.Duration(expireHours) * time.Hour)

	userId := strconv.FormatUint(uint64(userid), 10)

	claims := Claims{
		userId,
		username,
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
