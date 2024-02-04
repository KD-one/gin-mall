package cache

import (
	"context"
	"gin-mall/pkg/util"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient   *redis.Client
	RedisAddr     string
	RedisPassword string
	RedisDbName   string
)

func Init() {
	loadRedis()

	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		util.LogrusObj.Errorf("Redis出错！！: %v", err)
		panic(err)
	}
	RedisClient = client
}

// 加载redis配置文件(如果使用conf中的变量，会导致循环导入包)
func loadRedis() {
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		util.LogrusObj.Errorf("加载配置文件失败: %v", err)
		return
	}
	RedisAddr = file.Section("redis").Key("addr").String()
	RedisPassword = file.Section("redis").Key("password").String()
	RedisDbName = file.Section("redis").Key("dbName").String()
}
