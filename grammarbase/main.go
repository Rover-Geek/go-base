package main

import (
	"fmt"
	"go-base/entity"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 11:16 上午
 */

func main() {
	fmt.Println(">>>>>>>>>>>>>go基础语法学习开始>>>>>>>>>>>>>>")
	slice := slice()

	// 切片遍历
	for value := range slice {
		fmt.Println(value)
	}

	// map遍历
	m := mapInt()
	for k, v := range m {
		fmt.Print(k, "=", v, ", ")
	}
}

func slice() []*entity.Rover {
	arr := make([]*entity.Rover, 0)

	for i:= 0; i < 2; i ++ {
		rover := entity.Rover{
			Id: i,
			Name: fmt.Sprintf("name-%v", i),
			Pass: fmt.Sprintf("pass-%v", i),
		}
		arr = append(arr, &rover)
	}
	return arr
}

func mapInt() map[string] int {
	m := make(map[string] int, 2)

	for i:= 0; i < 5; i ++ {
		m[fmt.Sprintf("v-%v", i)] = i
	}
	return m
}