package main

import (
	"fmt"
	"log"

	"lol-rank/internal/config"
	"lol-rank/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化处理器
	handlers.InitHandlers(config.GlobalConfig.LCU.Port)

	r := gin.Default()

	// 设置路由
	setupRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// API 路由组
	api := r.Group("/api")
	{
		// 获取当前游戏状态
		api.GET("/game-status", handlers.GetGameStatus)

		// 获取排位数据
		api.GET("/ranked-stats", handlers.GetRankedStats)

		// 获取最近比赛记录
		api.GET("/recent-matches", handlers.GetRecentMatches)

		// 获取比赛分析
		api.GET("/match-analysis", handlers.GetMatchAnalysis)
	}
}
