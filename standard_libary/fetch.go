package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func fetchFromConfig(configPath string) error {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout())
	defer cancel()

	if err := fetchURL(ctx, cfg); err != nil {
		return err
	}

	fmt.Printf("请求成功，已保存到: %s\n", cfg.Output)
	return nil
}

func fetchURL(ctx context.Context, cfg *Config) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.URL, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("请求超时: 超过 %s", cfg.Timeout())
		}
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP 状态码错误: %d %s", resp.StatusCode, resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(cfg.Output), 0o755); err != nil {
		return fmt.Errorf("创建输出目录失败: %w", err)
	}

	out, err := os.Create(cfg.Output)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("写入响应内容失败: %w", err)
	}

	return nil
}
