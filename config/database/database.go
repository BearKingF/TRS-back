package database

import (
	"TRS/config/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.pass")
	port := config.Config.GetString("database.port")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.name")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
		//Logger: logger.Default.LogMode(logger.Info), //执行每条数据库相关语句都输出Debug()日志
	})

	if err != nil {
		log.Fatal("Database connect failed: ", err)
	}

	// 自动建表（已有不会覆盖）
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("Database migrate failed: ", err)
	}

	DB = db
}
