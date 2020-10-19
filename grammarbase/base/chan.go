package grammarbase

import "fmt"

/**
 * @description: channel 是 Go 语言在语言级别提供的 goroutine 间的通信方式，
 * 我们可以使用 channel 在多个 goroutine 之间传递消息。channel是进程内的通信方式。
 * @author: Rover
 * @date: 2020/10/19 10:30 上午
 */

func ChanBase() {
	ch := make(chan int)
	// 以下代码执行会报错：fatal error: all goroutines are asleep - deadlock!
	// 这是因为会在ch <- 100阻塞，知道ch的值被读出来
	/*ch <- 100
	v := <- ch
	fmt.Println("从cahn读取的v:", v)*/

	// 改写以上,将ch <- 100异步出来
	go func() {
		ch <- 100
	}()

	// 同样地，从channel中读取数据，如果channel之前没有写入数据，也会导致阻塞，直到channel中被写入数据为止
	v := <-ch
	fmt.Println("从chan读取的v:", v)
}

func UserChan() {
	ch := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		// 使用内置函数 make() 定义一个channel
		ch[i] = make(chan int)
		go add(i, ch[i])
	}

	for _, c := range ch {
		// ok.判断chan是否已经关闭， false 则表示 ch 已经被关闭
		v, ok := <-c
		if ok {
			fmt.Println("value:", v)
		}
		close(c)
	}

	fmt.Println("全部执行完毕")

}

func add(num int, ch chan int) {
	ch <- num
}
