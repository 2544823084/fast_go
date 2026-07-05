package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFetchURLSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	}))
	defer server.Close()

	dir := t.TempDir()
	output := filepath.Join(dir, "response.txt")
	cfg := &Config{
		URL:            server.URL,
		TimeoutSeconds: 3,
		Output:         output,
		LogLevel:       "info",
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout())
	defer cancel()

	if err := fetchURL(ctx, cfg); err != nil {
		t.Fatalf("fetchURL() 错误: %v", err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("读取输出文件失败: %v", err)
	}
	if string(data) != "hello world" {
		t.Fatalf("输出内容 = %q, want hello world", string(data))
	}
}

func TestFetchURLNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	cfg := &Config{
		URL:            server.URL,
		TimeoutSeconds: 3,
		Output:         filepath.Join(t.TempDir(), "out.txt"),
		LogLevel:       "info",
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout())
	defer cancel()

	err := fetchURL(ctx, cfg)
	if err == nil {
		t.Fatal("期望非 2xx 状态码返回错误")
	}
	if !strings.Contains(err.Error(), "HTTP 状态码错误: 404") {
		t.Fatalf("错误信息 = %q, want 包含 HTTP 状态码错误: 404", err.Error())
	}
}

func TestFetchURLTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprint(w, "slow")
	}))
	defer server.Close()

	cfg := &Config{
		URL:            server.URL,
		TimeoutSeconds: 1,
		Output:         filepath.Join(t.TempDir(), "out.txt"),
		LogLevel:       "info",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := fetchURL(ctx, cfg)
	if err == nil {
		t.Fatal("期望超时时返回错误")
	}
	if !strings.Contains(err.Error(), "请求超时") {
		t.Fatalf("错误信息 = %q, want 包含 请求超时", err.Error())
	}
}
