package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TextStats struct {
	Lines    int
	Words    int
	Chars    int
	WordFreq map[string]int
}

func analyzeTextFile(path string) (*TextStats, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	stats := &TextStats{
		Chars:    utf8.RuneCountInString(string(data)),
		WordFreq: make(map[string]int),
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		stats.Lines++
		for _, word := range tokenize(scanner.Text()) {
			stats.Words++
			stats.WordFreq[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func tokenize(line string) []string {
	var words []string
	var current strings.Builder

	emit := func() {
		if current.Len() == 0 {
			return
		}
		words = append(words, strings.ToLower(current.String()))
		current.Reset()
	}

	for _, r := range line {
		switch {
		case unicode.Is(unicode.Han, r):
			emit()
			words = append(words, string(r))
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			current.WriteRune(unicode.ToLower(r))
		default:
			emit()
		}
	}
	emit()

	return words
}

type wordCount struct {
	word  string
	count int
}

func topWords(freq map[string]int, n int) []wordCount {
	items := make([]wordCount, 0, len(freq))
	for word, count := range freq {
		items = append(items, wordCount{word: word, count: count})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].count != items[j].count {
			return items[i].count > items[j].count
		}
		return items[i].word < items[j].word
	})

	if len(items) > n {
		items = items[:n]
	}
	return items
}

func printTextStats(stats *TextStats) {
	fmt.Printf("行数: %d\n", stats.Lines)
	fmt.Printf("单词数: %d\n", stats.Words)
	fmt.Printf("字符数: %d\n", stats.Chars)

	fmt.Println("出现次数最多的单词:")
	for i, item := range topWords(stats.WordFreq, 10) {
		fmt.Printf("  %d. %s (%d)\n", i+1, item.word, item.count)
	}
}
