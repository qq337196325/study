package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func (c *MainController) Withtimeout() {
	log := logs.NewLogger()

	// WithTimeout设置超时时间，printNum2函数将ctx传下去，超过设定时间会自动调用cancel
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	printNum2(ctx)

	fmt.Println("结束222", ctx)
	log.Debug("批量新增33失败")
	c.Ctx.WriteString("批量新增33ff数据成功\n")
}

func printNum2(ctx context.Context) {

	//要打印的数字
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束")
			return
		default:
			fmt.Println("数字: ", n)
			n++
		}
		time.Sleep(time.Second)
	}
}
