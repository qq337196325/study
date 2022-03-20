package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"study/models"
	"time"
)

type RedisController struct {
	beego.Controller
}

//字符串操作
func (c *RedisController) Srt() {
	log := logs.NewLogger()
	Pool := models.Pool
	rc := Pool.Get()
	defer rc.Close()

	//EX 为KEY设置秒级过期时间；
	_, err := rc.Do("SET", "test:srt1", "测试字符村222", "EX", 10)
	if err != nil {
		log.Debug("新增失败")
		return
	}

	//读取数据
	srt, err := rc.Do("GET", "test:srt1")
	if err != nil {
		log.Debug("新增失败")
		return
	}
	fmt.Println("test:srt1的值为：", string(srt.([]byte)))

	//批量设置
	_, err = rc.Do("MSET", "test:msrt1", "test1", "test:msrt2", "test2", "test:msrt3", "test3")
	if err != nil {
		log.Debug("批量新增失败")
		return
	}

	//批量获取
	resp, err := rc.Do("MGET", "test:msrt1", "test:msrt2", "test:msrt3")
	re, err := redis.Values(resp, err)
	if err != nil {
		log.Debug("批量新增失败")
		return
	}
	//自定义获取值
	var str01, str02 string
	redis.Scan(re, &str01, &str02)
	fmt.Println("批量获取的值：", str01, str02)

	//解析全部获取到得值
	strs, err := redis.Strings(resp, err)
	fmt.Println("批量获取的值map：", strs)

	//解析全部获取到得值
	maps, err := redis.StringMap(resp, err)
	fmt.Println("批量获取的值map：", maps)

	c.Ctx.WriteString("批量新增数据成功\n")
}

//哈希操作
func (c *RedisController) Hash() {
	log := logs.NewLogger()
	Pool := models.Pool
	rc := Pool.Get()
	defer rc.Close()

	//int64转字符串
	hkey := "test:hash:" + strconv.FormatInt(time.Now().Unix(), 10)
	//设置哈希；
	hash := map[string]string{"test1": "111test1", "test2": "2222test2"}
	_, err := rc.Do("HMSET", redis.Args{}.Add(hkey).AddFlat(hash)...)
	if err != nil {
		log.Debug("新增失败")
		return
	}

	//批量获取
	resp, err := rc.Do("HGETALL", hkey)
	if err != nil {
		log.Debug("批量新增失败")
		return
	}
	hashs, err := redis.StringMap(resp, err)

	fmt.Println("批量获取的值map：", hashs)
	log.Debug("新增失败")
	c.Ctx.WriteString("批量新增数据成功\n")
}

//列表操作
func (c *RedisController) List() {
	log := logs.NewLogger()
	Pool := models.Pool
	rc := Pool.Get()
	defer rc.Close()

	//int64转字符串
	hkey := "test:hash:" + strconv.FormatInt(time.Now().Unix(), 10)
	//设置哈希；
	//keys := []string{"nn", "mm"}
	_, err := rc.Do("lpush", "testpush", hkey)
	if err != nil {
		log.Debug("新增失败")
		return
	}

	c.Ctx.WriteString("批量新增数据成功\n")
}

//集合操作
func (c *RedisController) Set() {
	log := logs.NewLogger()

	log.Debug("新增失败")
	c.Ctx.WriteString("批量新增数据成功\n")
}

//有序集合操作
func (c *RedisController) Sset() {
	log := logs.NewLogger()

	log.Debug("新增失败")
	c.Ctx.WriteString("批量新增数据成功\n")
}
