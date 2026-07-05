# Go 标准库练习项目说明

## 项目名称

`gobox`：一个用 Go 标准库实现的小型命令行工具箱。

## 项目目标

通过一个连续的小项目练习 Go 标准库中最常用的包。项目不依赖第三方库，重点掌握文件读写、JSON 解析、命令行参数、HTTP 请求、超时控制、日志、并发和测试。

完成后，你应该能够熟悉这些标准库：

- `fmt`
- `os`
- `flag`
- `strings`
- `strconv`
- `time`
- `errors`
- `io`
- `bufio`
- `path/filepath`
- `encoding/json`
- `net/http`
- `context`
- `sync`
- `testing`
- `log/slog`

## 项目功能

`gobox` 提供几个子功能，每个功能对应一组标准库练习。

### 1. 文本统计

读取一个文本文件，统计：

- 文件总行数
- 单词总数
- 字符总数
- 出现次数最多的前 10 个单词

练习标准库：

- `os`
- `io`
- `bufio`
- `strings`
- `sort`

示例命令：

```bash
go run . text -file ./data/article.txt
```

### 2. JSON 配置读取

读取一个 JSON 配置文件，配置内容包括：

- HTTP 请求地址
- 超时时间
- 输出文件路径
- 日志级别

练习标准库：

- `os`
- `encoding/json`
- `time`
- `errors`

示例配置：

```json
{
  "url": "https://example.com",
  "timeout_seconds": 3,
  "output": "./data/response.txt",
  "log_level": "info"
}
```

示例命令：

```bash
go run . config -file ./config.json
```

### 3. HTTP 请求工具

根据配置文件里的 URL 发起 HTTP GET 请求，将响应内容保存到文件。

要求：

- 支持请求超时
- 检查 HTTP 状态码
- 请求失败时返回清晰错误
- 响应内容写入指定文件

练习标准库：

- `net/http`
- `context`
- `time`
- `os`
- `io`
- `fmt`

示例命令：

```bash
go run . fetch -config ./config.json
```

### 4. 日志记录

为程序添加结构化日志，记录：

- 程序启动
- 读取的配置文件路径
- HTTP 请求地址
- 请求耗时
- 输出文件路径
- 错误信息

练习标准库：

- `log/slog`
- `os`
- `time`

### 5. 并发批量请求

读取一个 URL 列表文件，并发请求多个地址，输出每个地址的状态码和耗时。

要求：

- 使用 goroutine 并发请求
- 使用 `sync.WaitGroup` 等待任务完成
- 使用 channel 收集结果
- 为每个请求设置超时

练习标准库：

- `sync`
- `context`
- `net/http`
- `time`
- `bufio`

示例命令：

```bash
go run . batch -file ./data/urls.txt -timeout 5
```

## 推荐目录结构

```text
gobox/
  main.go
  config.go
  text.go
  fetch.go
  batch.go
  logger.go
  config_test.go
  text_test.go
  data/
    article.txt
    urls.txt
  config.json
```

## 开发阶段

### 第一阶段：命令行入口

目标：

- 使用 `os.Args` 或 `flag` 解析命令
- 支持 `text`、`config`、`fetch`、`batch` 四个子命令
- 参数错误时输出帮助信息

验收标准：

- 不带参数运行时提示可用命令
- 输入未知命令时提示错误
- 每个子命令能解析自己的参数

### 第二阶段：文本统计

目标：

- 实现文本文件读取
- 统计行数、单词数和字符数
- 统计单词出现次数

验收标准：

- 文件不存在时返回错误
- 空文件可以正常处理
- 中文字符不会被错误截断

### 第三阶段：配置读取

目标：

- 定义 `Config` 结构体
- 从 JSON 文件读取配置
- 校验 URL、超时时间和输出路径

验收标准：

- JSON 格式错误时返回错误
- 缺少必要字段时返回错误
- 超时时间小于等于 0 时返回错误

### 第四阶段：HTTP 请求

目标：

- 使用配置发起 HTTP GET 请求
- 用 `context.WithTimeout` 控制超时
- 将响应内容写入文件

验收标准：

- 请求成功后生成输出文件
- 非 2xx 状态码返回错误
- 超时时能正常退出并提示原因

### 第五阶段：并发批量请求

目标：

- 从文件读取多个 URL
- 并发请求 URL
- 汇总每个 URL 的状态码、耗时和错误

验收标准：

- 所有 URL 都有结果输出
- 某个 URL 失败不会影响其他 URL
- 程序不会提前退出或卡住

### 第六阶段：测试

目标：

- 为配置解析写单元测试
- 为文本统计写单元测试
- 为 HTTP 请求逻辑写测试

练习标准库：

- `testing`
- `net/http/httptest`
- `os`
- `path/filepath`

验收标准：

```bash
go test ./...
```

能够通过所有测试。

## 进阶挑战

如果基础功能完成，可以继续扩展：

- 使用 `flag.FlagSet` 改造子命令参数解析
- 为批量请求增加最大并发数限制
- 将结果输出为 JSON 文件
- 使用 `httptest` 模拟成功、失败和超时接口
- 使用 `errors.Is` 或 `errors.As` 处理自定义错误
- 增加 `-verbose` 参数控制日志详细程度

## 学习建议

每完成一个阶段，先写一个最小可运行版本，再补充错误处理和测试。不要一开始追求结构完美，先让程序跑起来，然后逐步整理代码。

推荐节奏：

1. 先实现功能。
2. 再处理错误。
3. 最后补测试。

这样能更清楚地看到每个标准库包解决了什么问题。
