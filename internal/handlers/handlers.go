package handlers

import (
	"net/http"
	"time"

	"lol-rank/internal/analysis"
	"lol-rank/internal/cache"
	"lol-rank/internal/config"
	"lol-rank/internal/lcu"

	"github.com/gin-gonic/gin"
)

var lcuClient *lcu.Client

// InitHandlers 初始化处理器
func InitHandlers(port int) {
	lcuClient = lcu.NewClient(port)
	cache.InitCache()
}

// GetGameStatus 获取当前游戏状态
func GetGameStatus(c *gin.Context) {
	// 检查缓存
	if config.GlobalConfig.Cache.Enabled {
		if cached, exists := cache.Get("game_status"); exists {
			c.JSON(http.StatusOK, cached)
			return
		}
	}

	status, err := lcuClient.GetGameStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get game status: " + err.Error(),
		})
		return
	}

	// 设置缓存
	if config.GlobalConfig.Cache.Enabled {
		cache.Set("game_status", status, time.Duration(config.GlobalConfig.Cache.TTL)*time.Second)
	}

	c.JSON(http.StatusOK, status)
}

// GetRankedStats 获取排位数据
func GetRankedStats(c *gin.Context) {
	// 检查缓存
	if config.GlobalConfig.Cache.Enabled {
		if cached, exists := cache.Get("ranked_stats"); exists {
			c.JSON(http.StatusOK, cached)
			return
		}
	}

	stats, err := lcuClient.GetRankedStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get ranked stats: " + err.Error(),
		})
		return
	}

	// 设置缓存
	if config.GlobalConfig.Cache.Enabled {
		cache.Set("ranked_stats", stats, time.Duration(config.GlobalConfig.Cache.TTL)*time.Second)
	}

	c.JSON(http.StatusOK, stats)
}

// GetRecentMatches 获取最近比赛记录
func GetRecentMatches(c *gin.Context) {
	// 检查缓存
	if config.GlobalConfig.Cache.Enabled {
		if cached, exists := cache.Get("recent_matches"); exists {
			c.JSON(http.StatusOK, cached)
			return
		}
	}

	matches, err := lcuClient.GetRecentMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recent matches: " + err.Error(),
		})
		return
	}

	// 设置缓存
	if config.GlobalConfig.Cache.Enabled {
		cache.Set("recent_matches", matches, time.Duration(config.GlobalConfig.Cache.TTL)*time.Second)
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchAnalysis 获取比赛分析
func GetMatchAnalysis(c *gin.Context) {
	// 检查缓存
	if config.GlobalConfig.Cache.Enabled {
		if cached, exists := cache.Get("match_analysis"); exists {
			c.JSON(http.StatusOK, cached)
			return
		}
	}

	matches, err := lcuClient.GetRecentMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recent matches: " + err.Error(),
		})
		return
	}

	// 分析比赛数据
	analysis := analysis.AnalyzeMatches(matches)

	// 设置缓存
	if config.GlobalConfig.Cache.Enabled {
		cache.Set("match_analysis", analysis, time.Duration(config.GlobalConfig.Cache.TTL)*time.Second)
	}

	c.JSON(http.StatusOK, analysis)
}
