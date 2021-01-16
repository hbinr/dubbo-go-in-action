package main

import (
	gxlog "github.com/dubbogo/gost/log"
	"study.dubbogo/02-custom-config-file/pkg/cache"
	"study.dubbogo/02-custom-config-file/pkg/database"

	"study.dubbogo/02-custom-config-file/pkg/conf"
)

func main() {
	dataConf, err := conf.Init()
	if err != nil {
		gxlog.CError("main:dataconfig init failed,err is %v", err)
		return
	}
	gxlog.CInfo("main:dataconfig init success,dataconfig:%v", dataConf)

	gormDB, err := database.InitMySQL(dataConf)
	if err != nil {
		gxlog.CError("main:mysql init failed,err is %v", err)
		return
	}
	gxlog.CInfo("main:gormDB init success,gormDB:%v", gormDB)

	redisClient, err := cache.InitRedis(dataConf)
	if err != nil {
		gxlog.CError("main:redis init failed,err is %v", err)
		return
	}
	gxlog.CInfo("main:redis Client init success,redisClient:%v", redisClient)

}
