package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-base/client"
	"go-base/entity"
	"log"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 12:24 上午
 */

var db *sql.DB

func Init() {
	db = client.Db
}

func Save(rover *entity.Rover) (int64, error) {
	result, err := db.Exec("insert into rover (name, pass) values (?, ?)", rover.Name, rover.Pass)
	if err != nil {
		log.Printf("数据库新增数据失败，%v", err)
		return -1, err
	}
	// 获取刚刚新增的自增ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("数据库新增数据失败，%s", err)
		return -1, err
	}
	log.Printf("数据库新增返回的ID:%v", id)

	return id, nil
}

// 查询单条数据
func QueryById(id int) *entity.Rover {
	rover := new(entity.Rover)
	row := db.QueryRow("select * from rover where id = ?", id)
	err := row.Scan(&rover.Id, &rover.Name, &rover.Pass)

	if err != nil {
		log.Printf("查询数据失败，id:%s,%v", id, err)
	}
	return rover
}

func QueryList(id int) []*entity.Rover {
	arr := make([]*entity.Rover, 0)
	rows, err := db.Query("select * from rover where id > ?", id)
	if err != nil {
		log.Printf("多行数据查询失败,id:%v, 错误:%v", id, err)
	}

	for rows.Next() {
		rover := new(entity.Rover)
		// 数据解析失败,id:0, 错误:sql: Scan error on column index 0, name "id": destination not a pointer
		// 此时需要将rover.Id =》 &rover.Id
		if err := rows.Scan(&rover.Id, &rover.Name, &rover.Pass); err != nil {
			log.Printf("数据解析失败,id:%v, 错误:%v", rover.Id, err)
		}
		arr = append(arr, rover)
	}
	return arr
}
