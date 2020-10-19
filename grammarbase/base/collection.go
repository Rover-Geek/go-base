package grammarbase

import (
	"fmt"
	"go-base/entity"
)

/**
 * @description: 集合相关
 * @author: Rover
 * @date: 2020/8/30 10:25 下午
 */

func Slice() []*entity.Rover {
	arr := make([]*entity.Rover, 0)

	for i := 0; i < 2; i++ {
		rover := entity.Rover{
			Id:   i,
			Name: fmt.Sprintf("name-%v", i),
			Pass: fmt.Sprintf("pass-%v", i),
		}
		arr = append(arr, &rover)
	}
	return arr
}

func MapInt() map[string]int {
	m := make(map[string]int, 2)

	for i := 0; i < 5; i++ {
		m[fmt.Sprintf("v-%v", i)] = i
	}
	return m
}
