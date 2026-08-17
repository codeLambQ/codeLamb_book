package config

import "os"

// Config 应用配置。
type Config struct {
	Server ServerConfig
}

// ServerConfig HTTP 服务配置。
type ServerConfig struct {
	Address string
}

// Load 从环境变量加载配置，未设置时使用默认值。
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDR", ":8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
