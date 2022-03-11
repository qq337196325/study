package main

import (
	_ "study/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

