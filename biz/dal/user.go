package dal

import (
	"errors"
	"fmt"
	"hertz_demo/biz/dbmodel"

	"gorm.io/gorm"
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
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在或已被删除")
		}
		return err
	}

	return DB.Delete(&user).Error
}

// GetUserByID 根据用户 ID 获取用户信息
func GetUserByID(userId int) (*dbmodel.User, error) {
	var user dbmodel.User
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在时返回 nil
		}
		return nil, err // 其他错误
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *dbmodel.User) error {
	return DB.Updates(user).Error
}

// GetUserList 获取用户列表（分页）
func GetUserList(pageSize int, offset int, username string) ([]*dbmodel.User, int64, error) {
	var users []*dbmodel.User
	var total int64

	// 查询总数
	if err := DB.Model(&dbmodel.User{}).Where("username LIKE ?", "%"+username+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := DB.Model(&dbmodel.User{}).Where("username LIKE ?", "%"+username+"%").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
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
