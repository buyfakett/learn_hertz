package dal

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm/logger"
	"hertz_demo/biz/dal/mysql"
	"hertz_demo/biz/dal/postgres"
	"hertz_demo/biz/dal/sqlite"
	"hertz_demo/bootstrao"
	"hertz_demo/utils/config"
)

func Init() {
	dbType := config.Cfg.Db.Type

	hlog.Infof("当前数据库为%s", dbType)

	var gormLogger logger.Interface
	if config.Cfg.Server.LogLevel != "debug" {
		gormLogger = logger.Default.LogMode(logger.Error) // 只有错误日志
	} else {
		gormLogger = logger.Default.LogMode(logger.Info) // 输出信息级别的日志
	}

	switch dbType {
	case "mysql":
		db := mysql.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database, gormLogger)
		err := bootstrao.Migrate(db)
		if err != nil {
			return
		}
	case "postgres":
		db := postgres.Init(config.Cfg.Db.User, config.Cfg.Db.Password, config.Cfg.Db.Host, config.Cfg.Db.Port, config.Cfg.Db.Database, gormLogger)
		err := bootstrao.Migrate(db)
		if err != nil {
			return
		}
	case "sqlite3":
		db := sqlite.Init(config.Cfg.Db.Database, gormLogger)
		err := bootstrao.Migrate(db)
		if err != nil {
			return
		}
	}

}
