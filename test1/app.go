package main

import (
	"fmt"
	"time"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/28 11:10 下午
 */

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Print("睡眠结束")

	} ()

	fmt.Print("》》》》》")

	time.Sleep(6 * time.Second)

}