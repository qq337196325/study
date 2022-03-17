package controllers

import (
	"fmt"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func (c *MainController) Testsync() {
	log := logs.NewLogger()

	var countGuard sync.Mutex
	var tint int64
	for i := 0; i < 100; i++ {
		go func() {
			countGuard.Lock()
			tint = time.Now().UnixNano() / 1e6
			fmt.Printf("循环：%d\n", tint)
			countGuard.Unlock()
		}()
	}

	log.Debug("批量新增33失败")
	c.Ctx.WriteString("批量新增33ff数据成功\n")
}
