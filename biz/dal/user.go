package dal

import (
	"errors"
	"github.com/hertz-contrib/jwt"
	"golang.org/x/crypto/bcrypt"
	"hertz_demo/biz/dbmodel"
)

var jwtMiddleware *jwt.HertzJWTMiddleware

func CreateUser(users []*dbmodel.User) error {
	return DB.Create(users).Error
}

func DeleteUser(userId int64) error {
	return DB.Where("id = ?", userId).Delete(&dbmodel.User{}).Error
}

func UpdateUser(user *dbmodel.User) error {
	return DB.Updates(user).Error
}

func UserLogin(username string, password string) (string, error) {
	var user dbmodel.User

	// 根据用户名查找用户
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("用户名不存在")
	}

	// 验证密码（假设密码使用 bcrypt 加密存储）
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 创建 JWT Token
	token, _, err := jwtMiddleware.TokenGenerator(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}
