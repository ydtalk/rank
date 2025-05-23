package lcu

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	LCUBaseURL = "https://127.0.0.1:%d"
)

type Client struct {
	client *resty.Client
	port   int
}

// NewClient 创建一个新的 LCU 客户端
func NewClient(port int) *Client {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// 从 lockfile 获取认证信息
	authToken := getAuthToken()
	client.SetHeader("Authorization", fmt.Sprintf("Basic %s", authToken))

	return &Client{
		client: client,
		port:   port,
	}
}

// getAuthToken 从 lockfile 获取认证令牌
func getAuthToken() string {
	// 获取 League of Legends 安装路径
	riotGamesPath := os.Getenv("PROGRAMFILES") + "\\Riot Games"
	lockfilePath := filepath.Join(riotGamesPath, "League of Legends", "lockfile")

	// 读取 lockfile
	content, err := os.ReadFile(lockfilePath)
	if err != nil {
		return ""
	}

	// 解析 lockfile 内容
	parts := strings.Split(string(content), ":")
	if len(parts) < 3 {
		return ""
	}

	// 返回认证令牌
	return parts[3]
}

// GetGameStatus 获取当前游戏状态
func (c *Client) GetGameStatus() (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := c.client.R().
		SetResult(&result).
		Get(fmt.Sprintf(LCUBaseURL+"/lol-gameflow/v1/session", c.port))

	return result, err
}

// GetRankedStats 获取排位数据
func (c *Client) GetRankedStats() (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := c.client.R().
		SetResult(&result).
		Get(fmt.Sprintf(LCUBaseURL+"/lol-ranked/v1/current-ranked-stats", c.port))

	return result, err
}

// GetRecentMatches 获取最近比赛记录
func (c *Client) GetRecentMatches() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	_, err := c.client.R().
		SetResult(&result).
		Get(fmt.Sprintf(LCUBaseURL+"/lol-match-history/v1/products/lol/current-summoner/matches", c.port))

	return result, err
}
