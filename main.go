package main

import (
	"context"
	"flag"
	"fmt"
	"singo/bootstrap"
	"singo/config"
	"singo/server"
	"singo/util"
)

func main() {
	ctx := context.Background()

	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
	}
	flag.Parse()

	// 初始化
	bootstrap.Init(ctx, *env)

	util.Log().Info("测试")

	// 装载路由
	cfg := config.GetConfig()
	r := server.NewRouter()
	r.Run(fmt.Sprintf(":%s", cfg.GetString("gin.port")))
}
