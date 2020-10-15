package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"singo/cache"
	"singo/config"
	"singo/model"
	"singo/util"
)

// Init 初始化配置项
func Init(ctx context.Context, env string) {
	config.Init(env)
	cfg := config.GetConfig()

	gin.SetMode(cfg.GetString("gin.mode"))

	// 设置日志级别
	util.BuildLogger(cfg.GetString("log.level"))

	// 读取翻译文件
	if err := LoadLocales("bootstrap/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(cfg.GetString("mysql.dsn"))
	cache.Redis(ctx)
}
