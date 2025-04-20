package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"gorm.io/gorm"
	"hertz_demo/biz/dbmodel"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("your-secret"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*dbmodel.User); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &dbmodel.User{
				Model: gorm.Model{
					ID: uint(claims["id"].(float64)),
				},
			}
		},
	})
}
