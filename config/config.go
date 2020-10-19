package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * @description: 中间件连接初始化
 * @author: Rover
 * @date: 2020/9/1 3:16 下午
 */

var env string

func Init() {

	env = GetEnv()
	cfgName := "app-" + env

	cfgPath := os.Getenv("GOPATH") + "/go-base/resource"

	viper.SetConfigName(cfgName)
	viper.AddConfigPath(cfgPath)

	viper.ReadInConfig()

	log.Infof("配置加载成功, env: %s", env)

	initLog()
}

func GetEnv() string {
	env = os.Getenv("env")
	if env == "" {
		return "dev"
	}
	return env
}

func GetCfg(key string) interface{} {
	return viper.Get(key)
}

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})

	filename := GetCfg("logger.filename").(string) + "-" + ".log"

	log.SetOutput(&lumberjack.Logger{
		Filename:  fmt.Sprintf(filename),
		MaxSize:   256,
		MaxAge:    7,
		LocalTime: true,
	})
}
