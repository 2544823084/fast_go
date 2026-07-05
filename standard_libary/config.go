package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

type Config struct {
	URL            string `json:"url"`
	TimeoutSeconds int    `json:"timeout_seconds"`
	Output         string `json:"output"`
	LogLevel       string `json:"log_level"`
}

func (c *Config) Validate() error {
	if strings.TrimSpace(c.URL) == "" {
		return errors.New("缺少必要字段: url")
	}
	parsed, err := url.Parse(c.URL)
	if err != nil {
		return fmt.Errorf("无效的 url: %w", err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("url 必须使用 http 或 https 协议")
	}
	if parsed.Host == "" {
		return errors.New("无效的 url: 缺少主机名")
	}

	if c.TimeoutSeconds <= 0 {
		return errors.New("timeout_seconds 必须大于 0")
	}

	if strings.TrimSpace(c.Output) == "" {
		return errors.New("缺少必要字段: output")
	}

	if strings.TrimSpace(c.LogLevel) == "" {
		return errors.New("缺少必要字段: log_level")
	}

	return nil
}

func (c *Config) Timeout() time.Duration {
	return time.Duration(c.TimeoutSeconds) * time.Second
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("JSON 格式错误: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func printConfig(cfg *Config) {
	fmt.Printf("URL: %s\n", cfg.URL)
	fmt.Printf("超时: %s\n", cfg.Timeout())
	fmt.Printf("输出: %s\n", cfg.Output)
	fmt.Printf("日志级别: %s\n", cfg.LogLevel)
}
