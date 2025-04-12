package bootstrao

import (
	"gorm.io/gorm"
	"hertz_demo/biz/dbmodel"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&dbmodel.User{},
	)
}
