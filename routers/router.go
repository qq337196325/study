package routers

import (
	"study/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/get", &controllers.MainController{}, "get:Get")

	//上下文
	beego.Router("/withtimeout", &controllers.MainController{}, "get:Withtimeout")

	//锁
	beego.Router("/testsync", &controllers.MainController{}, "get:Testsync")

	//生产消息
	beego.Router("/producer", &controllers.MainController{}, "get:Producer")

	//MongoDB操作
	beego.Router("/mongo/install", &controllers.MongoController{}, "get:Install")
	beego.Router("/mongo/update", &controllers.MongoController{}, "get:Update")
	beego.Router("/mongo/find", &controllers.MongoController{}, "get:Find")
}
