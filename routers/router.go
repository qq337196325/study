package routers

import (
	"study/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/get", &controllers.MainController{}, "get:Get")

	//Grpc客户端
	beego.Router("/GrpcClient", &controllers.MainController{}, "get:GrpcClient")

	//上下文
	beego.Router("/withtimeout", &controllers.MainController{}, "get:Withtimeout")

	//锁
	beego.Router("/testsync", &controllers.MainController{}, "get:Testsync")

	//生产消息
	beego.Router("/producer", &controllers.MainController{}, "get:Producer")

	//MongoDB操作
	nsBase := beego.NewNamespace("/mongo",
		beego.NSRouter("/install", &controllers.MongoController{}, "get:Install"),
		beego.NSRouter("/update", &controllers.MongoController{}, "get:Update"),
		beego.NSRouter("/find", &controllers.MongoController{}, "get:Find"),
	)
	beego.AddNamespace(nsBase)

	//Redis操作
	reBase := beego.NewNamespace("/redis",
		beego.NSRouter("/string", &controllers.RedisController{}, "get:Srt"),
		beego.NSRouter("/hash", &controllers.RedisController{}, "get:Hash"),
		beego.NSRouter("/list", &controllers.RedisController{}, "get:List"),
		beego.NSRouter("/set", &controllers.RedisController{}, "get:Set"),
		beego.NSRouter("/sset", &controllers.RedisController{}, "get:Sset"),
	)
	beego.AddNamespace(reBase)
}
