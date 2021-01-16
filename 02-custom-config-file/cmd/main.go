package main

import (
	gxlog "github.com/dubbogo/gost/log"

	"study.dubbogo/02-custom-config-file/pkg/conf"
)

func main() {
	dataConf, err := conf.Init()
	if err != nil {
		gxlog.CError("main:dataconfig init failed,err is %v", err)
		return
	}
	gxlog.CInfo("main:dataconfig init success,dataconfig:%v", dataConf)
}
