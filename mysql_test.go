package main

import (
	"fmt"
	"go-base/client"
	_ "go-base/client"
	"go-base/config"
	"go-base/entity"
	"go-base/mysql"
	"testing"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/9/1 11:12 下午
 */

func init() {
	config.Init()
	mysql.Init()
}

func TestSelectUser(t *testing.T) {

	r := entity.Rover{
		Name: "name-001",
		Pass: "pass-001",
	}

	mysql.Save(&r)

	rover := new(entity.Rover)
	row := client.Db.QueryRow("select * from rover where id = ?", 1)

	err2 := row.Scan(&rover.Id, &rover.Name, &rover.Pass)
	if err2 != nil {
		panic("查询失败")
	}
	fmt.Println(rover.Name)

	rover1 := mysql.QueryById(22)
	fmt.Println(rover1.Name)
}
