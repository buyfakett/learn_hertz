package dal

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hertz_demo/biz/dbmodel"
)

func CreateUser(users []*dbmodel.User) error {
	return DB.Create(users).Error
}

func IsUsernameExists(username string) (bool, error) {
	var count int64
	err := DB.Model(&dbmodel.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func DeleteUser(userId int) error {
	var user dbmodel.User
	err := DB.First(&user, "id = ? AND deleted_at IS NULL", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在或已被删除")
		}
		return err
	}

	return DB.Delete(&user).Error
}

func UpdateUser(user *dbmodel.User) error {
	return DB.Updates(user).Error
}

func UserLogin(username string) (*dbmodel.User, error) {
	var user dbmodel.User

	// 根据用户名查找用户
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}
