package main

import (
	"fmt"
	"go-base/entity"
	grammarbase "go-base/grammarbase/base"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 11:16 上午
 */

func main() {
	fmt.Println(">>>>>>>>>>>>>go基础语法学习开始>>>>>>>>>>>>>>")

	var num int
	fmt.Println(num)  // 初始值 0
	fmt.Println(&num) // 0xc000096008

	fmt.Println("--------基础---------")

	slice := grammarbase.Slice()

	// 切片遍历
	for value := range slice {
		fmt.Println(value)
	}

	// map遍历
	m := grammarbase.MapInt()
	for k, v := range m {
		fmt.Print(k, "=", v, ", ")
	}

	fmt.Println("\n----------defer---------")

	// defer
	grammarbase.DeferTest()

	//
	rover := entity.Rover{Id: 1, Name: "Rover", Pass: "pass"}

	// 不建议使用这种实例化对象的方式，否则struct中新增了属性，这里就需要更改
	rover1 := entity.Rover{1, "Rover001", "pass001", entity.Dog{No: int64(100), Name: "大狗子"}}

	roverWithDog := entity.Rover{
		Id:  100,
		Dog: entity.Dog{No: int64(100), Name: "大狗子"},
	}

	fmt.Println(&rover)
	fmt.Println(&rover1)
	fmt.Println(&entity.Rover{Id: 3, Name: "Rover003", Pass: "pass003"})
	fmt.Println(&roverWithDog)

	fmt.Println("------------chan测试------------")

	grammarbase.ChanBase()
	grammarbase.UserChan()

}
