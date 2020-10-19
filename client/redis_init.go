package client

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"go-base/config"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/9/1 2:12 下午
 */

var Redis *redis.Conn

func RedisInit() {
	_, _, server, port :=
		config.GetCfg("redis.username"),
		config.GetCfg("redis.password"),
		config.GetCfg("redis.server"),
		config.GetCfg("redis.port")

	c, err := redis.Dial("tcp", fmt.Sprintf("%v:%v", server, port))

	if err != nil {
		log.Infof("redis初始化失败，", err)
		panic(err)
	}
	Redis = &c
}
