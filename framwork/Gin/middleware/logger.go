package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Init 将 Gin 日志同时输出到控制台和 logs/gin.log，返回的 Closer 需在程序退出前关闭。
func Init() (io.Closer, error) {
	if err := os.MkdirAll("logs", 0o755); err != nil {
		return nil, fmt.Errorf("create logs dir: %w", err)
	}

	f, err := os.OpenFile("logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, f)
	return f, nil
}

// Logger 记录请求方法、路径、状态码、耗时和客户端 IP。
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Fprintf(gin.DefaultWriter,
			"[GIN] %s | %3d | %13v | %15s | %-7s %s\n",
			start.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
