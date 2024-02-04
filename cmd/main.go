package main

import (
	"gin-mall/cache"
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/routes"
)

func main() {
	// 初始化配置文件和MySQL
	conf.Init()

	dao.Database(conf.ReadDB, conf.WriteDB)

	// 初始化Redis缓存
	cache.Init()

	// 加载路由
	routes.NweRouter()
}
