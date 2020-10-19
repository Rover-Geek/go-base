package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"go-base/config"
	"time"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/9/1 2:10 下午
 */

var Db *sql.DB

func MysqlCfg() {
	username, password, network, server, database, dbPort :=
		config.GetCfg("mysql.username"),
		config.GetCfg("mysql.password"),
		config.GetCfg("mysql.network"),
		config.GetCfg("mysql.server"),
		config.GetCfg("mysql.database"),
		config.GetCfg("mysql.dbPort")

	connection := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, dbPort, database)
	log.Infof(connection)

	DB, err := sql.Open("mysql", connection)
	if err != nil {
		fmt.Print("数据库初始化连接失败", err)
		panic(err)
	}

	DB.SetMaxIdleConns(100)
	DB.SetConnMaxLifetime(100 * time.Second)

	Db = *&DB
}
