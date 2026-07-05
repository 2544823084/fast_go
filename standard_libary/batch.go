package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type batchResult struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}

func loadURLs(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("URL 列表为空")
	}
	return urls, nil
}

func batchFetch(urls []string, timeout time.Duration) []batchResult {
	results := make(chan batchResult, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results <- fetchOne(u, timeout)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	collected := make([]batchResult, 0, len(urls))
	for result := range results {
		collected = append(collected, result)
	}
	return collected
}

func fetchOne(url string, timeout time.Duration) batchResult {
	start := time.Now()
	result := batchResult{URL: url}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result.Duration = time.Since(start)
		result.Err = fmt.Errorf("创建请求失败: %w", err)
		return result
	}

	resp, err := http.DefaultClient.Do(req)
	result.Duration = time.Since(start)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			result.Err = fmt.Errorf("请求超时: 超过 %s", timeout)
		} else {
			result.Err = fmt.Errorf("请求失败: %w", err)
		}
		return result
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	result.StatusCode = resp.StatusCode
	return result
}

func printBatchResults(results []batchResult) {
	fmt.Printf("%-40s %8s %10s %s\n", "URL", "状态码", "耗时", "错误")
	fmt.Println(strings.Repeat("-", 80))
	for _, r := range results {
		status := "-"
		if r.StatusCode > 0 {
			status = fmt.Sprintf("%d", r.StatusCode)
		}
		errMsg := ""
		if r.Err != nil {
			errMsg = r.Err.Error()
		}
		fmt.Printf("%-40s %8s %10s %s\n", r.URL, status, r.Duration.Round(time.Millisecond), errMsg)
	}
}

func runBatchFetch(file string, timeoutSec int) error {
	if timeoutSec <= 0 {
		return fmt.Errorf("timeout 必须大于 0")
	}

	urls, err := loadURLs(file)
	if err != nil {
		return err
	}

	timeout := time.Duration(timeoutSec) * time.Second
	results := batchFetch(urls, timeout)
	printBatchResults(results)
	return nil
}
