package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)

	//defer close(messages)
	//defer cancel()

	select {
	case <-ctx.Done():
		//如果WithTimeout参数为：context.TODO()  会先执行这里(先执行外面的再执行里面的)
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
}
