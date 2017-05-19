package model

import (
	"maid-data-log/model/schema"
	"maid-data-log/utils"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitConnect() {
	config := utils.GetConfig()
	var connectErr error
	engine, connectErr = xorm.NewEngine(config.Db.Driver, config.Db.Connect)
	if connectErr != nil {
		log.Fatalln(connectErr)
		return
	}
	if engine.Ping() != nil {
		log.Fatalln("数据库连接超时...")
	} else {
		log.Printf("数据库连接成功")
	}
	err := engine.Sync2(new(schema.Log))
	if err != nil {
		log.Fatalln("数据库同步表格失败")
	}
}

func GetDBConect() *xorm.Engine {
	return engine
}
