package mysql

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 3:17 下午
 */
const (
	username        = "root"
	password        = "mysql123**"
	network         = "tcp"
	server          = "47.104.80.171"
	dbPort          = 3306
	database        = "db1"
)

func Init() sql.DB {
	connection := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",username, password, network, server, dbPort, database)
	log.Infof(connection)

	DB, err := sql.Open("mysql", connection)
	if err != nil {
		fmt.Print("数据库初始化连接失败", err)
	}
	DB.SetMaxIdleConns(100)
	DB.SetConnMaxLifetime(100 * time.Second)

	return *DB
}
