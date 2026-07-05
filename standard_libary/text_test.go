package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzeTextFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "article.txt")
	content := "Go is great.\n学习 Go 很有趣。\nGo go go!\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("写入测试文件失败: %v", err)
	}

	stats, err := analyzeTextFile(path)
	if err != nil {
		t.Fatalf("analyzeTextFile() 错误: %v", err)
	}
	if stats.Lines != 3 {
		t.Fatalf("Lines = %d, want 3", stats.Lines)
	}
	if stats.Words != 12 {
		t.Fatalf("Words = %d, want 12", stats.Words)
	}
	if stats.Chars != len([]rune(content)) {
		t.Fatalf("Chars = %d, want %d", stats.Chars, len([]rune(content)))
	}
	if stats.WordFreq["go"] != 5 {
		t.Fatalf(`WordFreq["go"] = %d, want 5`, stats.WordFreq["go"])
	}
}

func TestAnalyzeTextFileEmpty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatalf("写入测试文件失败: %v", err)
	}

	stats, err := analyzeTextFile(path)
	if err != nil {
		t.Fatalf("analyzeTextFile() 错误: %v", err)
	}
	if stats.Lines != 0 || stats.Words != 0 || stats.Chars != 0 {
		t.Fatalf("空文件统计应为 0, got lines=%d words=%d chars=%d", stats.Lines, stats.Words, stats.Chars)
	}
}

func TestAnalyzeTextFileNotFound(t *testing.T) {
	_, err := analyzeTextFile(filepath.Join(t.TempDir(), "missing.txt"))
	if err == nil {
		t.Fatal("期望文件不存在时返回错误")
	}
}

func TestTokenizeChinese(t *testing.T) {
	words := tokenize("你好世界")
	if len(words) != 4 {
		t.Fatalf("len(words) = %d, want 4", len(words))
	}
	want := []string{"你", "好", "世", "界"}
	for i, word := range words {
		if word != want[i] {
			t.Fatalf("words[%d] = %q, want %q", i, word, want[i])
		}
	}
}

func TestTopWords(t *testing.T) {
	freq := map[string]int{
		"go":   5,
		"rust": 2,
		"java": 3,
	}
	top := topWords(freq, 2)
	if len(top) != 2 {
		t.Fatalf("len(top) = %d, want 2", len(top))
	}
	if top[0].word != "go" || top[0].count != 5 {
		t.Fatalf("top[0] = %+v, want go:5", top[0])
	}
}
