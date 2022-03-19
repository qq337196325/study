package models

import (
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mdb *Mongodb

type Mongodb struct {
	Client *mongo.Client
	User   *mongo.Collection
}

func (db *Mongodb) NewMongo() {
	mongodb_ip, _ := config.String("mongodb_ip")
	param := fmt.Sprintf("mongodb://" + mongodb_ip + ":27017")
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

var RedisExpire = 3600 //缓存有效期
var Pool *redis.Pool

func NewRedis() {
	//连接redis
	//c, err := redis.Dial("tcp", "42.192.124.106:6379",
	//	redis.DialPassword("admin123"),
	//)

	Pool = &redis.Pool{
		//MaxIdle:     8,   // 最大空闲连接数
		//MaxActive:   0,   // 最大连接数，0表示没有限制
		//IdleTimeout: 100, // 空闲超时时间
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			redis_ip, _ := config.String("redis_ip")
			c, err := redis.Dial("tcp", redis_ip+":6379")
			if err != nil {
				fmt.Println("conn redis failed,", err)
				return nil, err
			}
			if _, err = c.Do("AUTH", "admin123"); err != nil {
				c.Close()
				fmt.Println("conn redis password failed,", err)
				return nil, err
			}
			fmt.Println("连接成功")
			return c, nil
		},
	}

	// 从池里获取连接
	rc := Pool.Get()
	key := "redis.key"
	_, err := rc.Do("Set", key, "1", "EX", RedisExpire)
	rc.Do("Set", "aaa", "1")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	//fmt.Println("redis conn success", Pool)
	//defer c.Close()
}
