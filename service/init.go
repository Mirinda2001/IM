package service

import (
	"IM/model"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	drivername := "mysql"
	DsName := "root:123456@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivername, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	// 是否显示SQL语句
	DbEngin.ShowSQL(true)
	// 数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)
	// 自动User 创建表
	DbEngin.Sync2(new(model.User))
	fmt.Println("Db init Ok")
}
