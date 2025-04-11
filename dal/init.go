package dal

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz_demo/config"
	"hertz_demo/dal/mysql"
	"hertz_demo/dal/postgres"
	"hertz_demo/dal/sqlite"
)

func Init() {
	hlog.Infof("当前数据库为%s", config.Cfg.Db.Type)
	switch config.Cfg.Db.Type {
	case "mysql":
		mysql.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database)
	case "postgres":
		postgres.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database)
	case "sqlite3":
		sqlite.Init(config.Cfg.Db.Database)
	}

}
