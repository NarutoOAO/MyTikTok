package main

import (
	"MyTikTok/conf"
	util "MyTikTok/pkg/utils"
	"MyTikTok/repository/db/dao"
	"MyTikTok/router"
)

func main() {
	conf.Init()
	dao.InitMySQL()
	//cache.InitCache()
	util.InitLog()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
