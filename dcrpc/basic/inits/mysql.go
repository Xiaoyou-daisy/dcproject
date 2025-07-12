package inits

import (
	"dcproject/dcrpc/basic/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "zg6project:2003225zyh@tcp(14.103.243.149:3306)/zg6project?charset=utf8mb4&parseTime=True&loc=Local"
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据库连接成功")
	}

}
