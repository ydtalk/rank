package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	LCU      LCUConfig      `mapstructure:"lcu"`
	Cache    CacheConfig    `mapstructure:"cache"`
	Analysis AnalysisConfig `mapstructure:"analysis"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type LCUConfig struct {
	Port    int    `mapstructure:"port"`
	BaseURL string `mapstructure:"base_url"`
}

type CacheConfig struct {
	Enabled bool `mapstructure:"enabled"`
	TTL     int  `mapstructure:"ttl"`
}

type AnalysisConfig struct {
	MatchHistoryLimit int `mapstructure:"match_history_limit"`
	MinGamesForStats  int `mapstructure:"min_games_for_stats"`
}

var GlobalConfig Config

// LoadConfig 加载配置文件
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&GlobalConfig)
}
