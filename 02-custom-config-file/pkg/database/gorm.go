package database

import (
	"database/sql"

	"study.dubbogo/02-custom-config-file/pkg/conf"

	"github.com/apache/dubbo-go/common/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	sqlDB *sql.DB
)

// InitMySQL init gorm.DB
func InitMySQL(cfg *conf.AppConfig) (db *gorm.DB, err error) {
	mysqlConf := mysql.Config{
		DSN:                       cfg.DSN, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}
	gormConfig := isLogOn(cfg.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConf), gormConfig); err != nil {
		logger.Error("database:opens database failed", err)
		return
	}

	if sqlDB, err = db.DB(); err != nil {
		logger.Error("database:get DB() failed", err)
		return
	}
	// GORM 使用 database/sql 维护连接池
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) // 设置打开数据库连接的最大数量
	return
}

// isLogOn gorm 日志配置，mod为true，表示开启日志
func isLogOn(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   gormLog.Default.LogMode(gormLog.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				// Whether the data table name is in plural form,default:false for plural table's name
				SingularTable: true,
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   gormLog.Default.LogMode(gormLog.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	}
	return
}
