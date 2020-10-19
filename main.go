package main

import (
	log "github.com/sirupsen/logrus"
	"go-base/api"
	"go-base/client"
	"go-base/config"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 1:06 上午
 */

func init() {
	config.Init()

	log.Infof("中间件连接初始化开始...")
	client.MysqlCfg()
	client.RedisInit()
	log.Infof("中间件连接初始化结束.")
}

func main() {

	api.Server()
}
