package grammarbase

import (
	"fmt"
	"time"
)

/**
 * @description: defer，函数最后的执行流程，类似于Java的finally
 * @author: Rover
 * @date: 2020/8/30 10:29 下午
 */

func DeferTest() {
	start := time.Now()
	// 会立刻对函数中引用的外部参数进行拷贝，所以 time.Since(startedAt)
	// 的结果不是在 main 函数退出之前计算的，而是在 defer 关键字调用时计算的
	defer fmt.Println("1方法耗时:%v", time.Since(start))

	defer func() {
		fmt.Println("2方法耗时:%v", time.Since(start))
	}()

	num := 0
	defer func(num int) {
		fmt.Println("num:", num)
	}(num) // num是值拷贝，num = 100并不影响最终defer中的num = 0的输出

	num = 100

	time.Sleep(1 * time.Second)

	// 退出进程时，defer不会执行
	//os.Exit(1)
}
