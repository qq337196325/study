package models

import (
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mdb *Mongodb

type Mongodb struct {
	Client *mongo.Client
	User   *mongo.Collection
}

func (db *Mongodb) NewMongo() {
	param := fmt.Sprintf("mongodb://127.0.0.1:27017")
	clientOptions := options.Client().ApplyURI(param)

	// 建立客户端连接
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}

	// 检查连接情况
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}

	//指定要操作的数据集
	db.User = client.Database("testmongo").Collection("user")
	db.Client = client
	Mdb = db
}

type Mysql struct {
	Client *mongo.Client
	User   *mongo.Collection
}

type Users struct {
	Id         int32 `orm:"auto"`
	Name       string
	Email      string
	Date       time.Time `orm:"type(date)"`
	Datetime   time.Time `orm:"type(datetime)"`
	Createtime int32
	Updatetime int32
	Deletetime int32
}

func (u *Users) TableName() string {
	return "user"
}

func (m *Mysql) NewMysql() {

	//获取配置信息
	mysql_ip, _ := config.String("mysql_ip")
	mysql_user, _ := config.String("mysql_user")
	mysql_password, _ := config.String("mysql_password")
	fmt.Println(mysql_ip)

	//注册model与连接数据库
	orm.RegisterModel(new(Users))
	orm.RegisterDataBase("default", "mysql", mysql_user+":"+mysql_password+"@tcp("+mysql_ip+")/beegoorm?charset=utf8")

	// 创建一个 Ormer
	Ormer := orm.NewOrm()
	user := &Users{
		Name:       "333",
		Email:      "333",
		Date:       time.Now(),
		Datetime:   time.Now(),
		Createtime: 333,
		Updatetime: 333,
		Deletetime: 33,
	}
	Ormer.Insert(user)
}
