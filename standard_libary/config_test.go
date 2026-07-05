package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeTempFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("写入测试文件失败: %v", err)
	}
	return path
}

func TestLoadConfigValid(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "config.json", `{
		"url": "https://example.com",
		"timeout_seconds": 3,
		"output": "./data/out.txt",
		"log_level": "info"
	}`)

	cfg, err := loadConfig(path)
	if err != nil {
		t.Fatalf("loadConfig() 错误: %v", err)
	}
	if cfg.URL != "https://example.com" {
		t.Fatalf("URL = %q, want https://example.com", cfg.URL)
	}
	if cfg.Timeout() != 3*time.Second {
		t.Fatalf("Timeout() = %v, want 3s", cfg.Timeout())
	}
}

func TestLoadConfigInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "bad.json", `{invalid`)

	_, err := loadConfig(path)
	if err == nil {
		t.Fatal("期望 JSON 格式错误")
	}
	if !strings.Contains(err.Error(), "JSON 格式错误") {
		t.Fatalf("错误信息 = %q, want 包含 JSON 格式错误", err.Error())
	}
}

func TestLoadConfigMissingFields(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "missing.json", `{
		"url": "https://example.com",
		"timeout_seconds": 3,
		"log_level": "info"
	}`)

	_, err := loadConfig(path)
	if err == nil {
		t.Fatal("期望缺少 output 时返回错误")
	}
	if !strings.Contains(err.Error(), "output") {
		t.Fatalf("错误信息 = %q, want 包含 output", err.Error())
	}
}

func TestConfigValidateTimeout(t *testing.T) {
	cfg := &Config{
		URL:            "https://example.com",
		TimeoutSeconds: 0,
		Output:         "./out.txt",
		LogLevel:       "info",
	}
	if err := cfg.Validate(); err == nil {
		t.Fatal("期望 timeout_seconds <= 0 时返回错误")
	}
}

func TestConfigValidateInvalidURL(t *testing.T) {
	cfg := &Config{
		URL:            "ftp://example.com",
		TimeoutSeconds: 3,
		Output:         "./out.txt",
		LogLevel:       "info",
	}
	if err := cfg.Validate(); err == nil {
		t.Fatal("期望非 http/https URL 返回错误")
	}
}
