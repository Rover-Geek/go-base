package api

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"go-base/config"
	"strconv"
)

/**
 * @description: echo是go语言开发的一种高性能，可扩展，轻量级的web框架
 * @author: Rover
 * @date: 2020/8/30 12:33 下午
 */

func Server() {
	// 实例化echo对象。
	e := echo.New()

	e.POST("/v1/rover/save", saveRover)
	e.GET("/v1/rover/:id", getRoverById)
	e.GET("/v1/rover/get", getRoverById2)

	port := config.GetCfg("port").(int)
	p := ":" + strconv.Itoa(port)

	err := e.Start(p)

	if err != nil {
		log.Infof("服务启动失败，", err)
	}

}
