package dal

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz_demo/biz/dal/mysql"
	"hertz_demo/biz/dal/postgres"
	"hertz_demo/biz/dal/sqlite"
	"hertz_demo/utils/config"
)

func Init() {
	dbType := config.Cfg.Db.Type

	hlog.Infof("当前数据库为%s", dbType)

	switch dbType {
	case "mysql":
		mysql.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database)
	case "postgres":
		postgres.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database)
	case "sqlite3":
		sqlite.Init(config.Cfg.Db.Database)
	}

}
