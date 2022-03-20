package main

import (
	"fmt"
	"time"
)

//测试channel
//链接地址： https://juejin.cn/post/6844904070667321357
func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {

		//Ticker是一个周期触发定时的计时器，它会按照一个时间间隔往channel发送系统当前时间，而channel的接收者可以以固定的时间间隔从channel中读取事件。
		ticker := time.NewTicker(1 * time.Second)

		for _ = range ticker.C {
			fmt.Println("time.Second:", time.Second)
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)

	//close函数是一个内建函数， 用来关闭channel，这个channel要么是双向的， 要么是只写的（chan<- Type）。
	//这个方法应该只由发送者调用， 而不是接收者。
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
