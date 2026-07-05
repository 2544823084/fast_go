package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `gobox - Go 标准库练习工具箱

用法:
  gobox <command> [options]

可用命令:
  text    文本统计
  config  JSON 配置读取
  fetch   HTTP 请求
  batch   并发批量请求

使用 "gobox <command> -h" 查看子命令帮助。
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "text":
		runText(os.Args[2:])
	case "config":
		runConfig(os.Args[2:])
	case "fetch":
		runFetch(os.Args[2:])
	case "batch":
		runBatch(os.Args[2:])
	case "-h", "--help", "help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "未知命令: %q\n\n", os.Args[1])
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
}

func runText(args []string) {
	fs := flag.NewFlagSet("text", flag.ExitOnError)
	file := fs.String("file", "", "文本文件路径")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: gobox text -file <path>\n\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *file == "" {
		fmt.Fprintln(os.Stderr, "错误: 必须指定 -file 参数")
		fs.Usage()
		os.Exit(1)
	}
	stats, err := analyzeTextFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
	printTextStats(stats)
}

func runConfig(args []string) {
	fs := flag.NewFlagSet("config", flag.ExitOnError)
	file := fs.String("file", "", "JSON 配置文件路径")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: gobox config -file <path>\n\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *file == "" {
		fmt.Fprintln(os.Stderr, "错误: 必须指定 -file 参数")
		fs.Usage()
		os.Exit(1)
	}
	fmt.Printf("config: file=%s\n", *file)
}

func runFetch(args []string) {
	fs := flag.NewFlagSet("fetch", flag.ExitOnError)
	config := fs.String("config", "", "JSON 配置文件路径")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: gobox fetch -config <path>\n\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *config == "" {
		fmt.Fprintln(os.Stderr, "错误: 必须指定 -config 参数")
		fs.Usage()
		os.Exit(1)
	}
	fmt.Printf("fetch: config=%s\n", *config)
}

func runBatch(args []string) {
	fs := flag.NewFlagSet("batch", flag.ExitOnError)
	file := fs.String("file", "", "URL 列表文件路径")
	timeout := fs.Int("timeout", 5, "每个请求的超时时间（秒）")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: gobox batch -file <path> [-timeout 秒数]\n\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *file == "" {
		fmt.Fprintln(os.Stderr, "错误: 必须指定 -file 参数")
		fs.Usage()
		os.Exit(1)
	}
	fmt.Printf("batch: file=%s timeout=%ds\n", *file, *timeout)
}
