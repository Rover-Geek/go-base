package api

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"go-base/client"
	"go-base/entity"
	"go-base/mysql"
	"net/http"
	"strconv"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 11:58 上午
 */

const (
	success         = "success"
	successCode int = 200
	failed          = "failed"
	failedCode  int = 500
)

type (
	// 响应
	Resp struct {
		Code int
		Msg  string
		Data interface{}
	}

	// 分页查询
	Page struct {
		Start int
		Size  int
	}
)

var db = client.Db

// 服务健康检查接口
func check(c echo.Context) error {
	return c.String(http.StatusOK, success)
}

func saveRover(c echo.Context) error {
	var req entity.Rover
	// 绑定POST请求参数
	c.Bind(&req)

	_, err := mysql.Save(&req)
	if err != nil {
		return err
	}
	return nil
}

func getRoverById(c echo.Context) error {
	// 获取url参数
	id := c.Param("id")
	var idInt, err = strconv.Atoi(id)
	if err != nil {
		log.Infof("%v转string失败", id)
		return nil
	}
	log.Infof("getRoverById executing, id:%v", id)
	rover := mysql.QueryById(idInt)

	return c.JSON(http.StatusOK, rover)
}

func getRoverById2(c echo.Context) error {
	// 获取url参数
	id := c.QueryParam("id")
	var idInt, err = strconv.Atoi(id)
	if err != nil {
		log.Infof("%v转string失败", id)
		return c.JSON(http.StatusOK, &Resp{failedCode, "id转换失败", nil})
	}

	log.Infof("getRoverById2 executing, id:%v", id)
	rover := mysql.QueryById(idInt)

	return c.JSON(http.StatusOK, rover)
}
