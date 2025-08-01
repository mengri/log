# 日志库 - Log

一个功能强大、高性能的 Go 语言日志库，支持结构化日志记录、多传输器、文件日志轮转和灵活的格式化。

## 特性

- 🚀 **高性能**: 使用对象池和原子操作优化性能
- 📝 **结构化日志**: 支持带字段的结构化日志记录
- 🎯 **多级别**: 支持 Trace, Debug, Info, Warn, Error, Fatal, Panic 七个日志级别
- 🔄 **多传输器**: 支持同时输出到多个目标（文件、控制台等）
- 📁 **文件轮转**: 内置文件日志轮转，支持按时间分割和自动清理
- 🎨 **自定义格式**: 灵活的日志格式化器
- 🔍 **调用栈追踪**: 可选的调用者信息记录
- 🛡️ **线程安全**: 完全并发安全

## 安装

```bash
go get gitlab.ncader.com/common/log
```

## 快速开始

### 基本使用

```go
package main

import "gitlab.ncader.com/common/log"

func main() {
    // 基本日志记录
    log.Info("应用程序启动")
    log.Debug("调试信息")
    log.Warn("警告信息")
    log.Error("错误信息")
    
    // 格式化日志
    log.Infof("用户 %s 登录成功", "张三")
    log.Errorf("连接数据库失败: %v", err)
    
    // 带字段的结构化日志
    log.WithFields(log.Fields{
        "user_id":   123,
        "operation": "login",
        "ip":        "192.168.1.1",
    }).Info("用户操作日志")
}
```

### 设置日志级别

```go
import "gitlab.ncader.com/common/log"

func main() {
    // 设置全局日志级别
    log.SetLevel(log.DebugLevel)  // 只输出 Debug 级别及以上的日志
}
```

### 自定义Logger

```go
import (
    "os"
    "gitlab.ncader.com/common/log"
)

func main() {
    // 创建自定义传输器
    transporter := &log.Transporter{}
    transporter.SetOutput(os.Stdout)
    transporter.SetLevel(log.InfoLevel)
    transporter.SetFormatter(&log.LineFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
    
    // 创建自定义Logger
    logger := log.NewLogger(transporter, true, "myapp")
    logger.SetPrefix("[MyApp] ")
    
    logger.Info("这是自定义logger的消息")
}
```

## 文件日志配置

### 使用文件日志

```go
import (
    "time"
    "gitlab.ncader.com/common/log"
    "gitlab.ncader.com/common/log/filelog"
)

func main() {
    // 文件日志配置
    config := filelog.Config{
        Dir:    "/var/log/myapp",     // 日志目录
        File:   "app.log",            // 日志文件名
        Expire: 7 * 24 * time.Hour,   // 保留7天
        Period: filelog.Day,          // 按天分割
    }
    
    // 创建文件传输器
    fileTransporter, err := filelog.NewTransporter(config)
    if err != nil {
        panic(err)
    }
    
    // 重置全局日志传输器
    log.Reset(fileTransporter)
    
    log.Info("这条日志将写入文件")
}
```

### 初始化配置

```go
import (
    "gitlab.ncader.com/common/log/init"
)

func main() {
    // 使用配置结构体
    config := &init.ErrorLogConfig{
        LogDir:    "./logs",
        FileName:  "error.log",
        LogLevel:  "info",
        LogExpire: "7d",
        LogPeriod: "day",
    }
    
    // 应用配置
    log.Init(config)
}
```

## 核心概念

### 日志级别

```go
const (
    PanicLevel Level = iota  // 最高级别，记录后会panic
    FatalLevel              // 致命错误，记录后程序退出
    ErrorLevel              // 错误级别
    WarnLevel               // 警告级别
    InfoLevel               // 信息级别
    DebugLevel              // 调试级别
    TraceLevel              // 最详细的追踪级别
)
```

### 传输器（Transporter）

传输器负责将日志条目输出到特定的目标：

- **Transporter**: 基础传输器，可输出到任何 `io.Writer`
- **Complex**: 复合传输器，支持同时输出到多个传输器
- **FileTransporter**: 文件传输器，支持文件轮转和清理

### 格式化器（Formatter）

- **LineFormatter**: 行格式化器，输出可读的文本格式日志

## API 参考

### 全局函数

```go
// 基本日志记录
func Debug(args ...interface{})
func Info(args ...interface{})
func Warn(args ...interface{})
func Error(args ...interface{})
func Fatal(args ...interface{})
func Panic(args ...interface{})

// 格式化日志记录
func DebugF(format string, args ...interface{})
func Infof(format string, args ...interface{})
func Warnf(format string, args ...interface{})
func Errorf(format string, args ...interface{})
func Fatalf(format string, args ...interface{})

// 结构化日志
func WithFields(fields Fields) Builder

// 配置函数
func SetLevel(level Level)
func SetPrefix(prefix string)
func Reset(transports ...EntryTransporter)
func Close()
```

### Logger 类型

```go
type Logger struct {
    // ...
}

func NewLogger(transporter EntryTransporter, reportCaller bool, packageName string) *Logger
func (logger *Logger) SetPrefix(prefix string)
func (logger *Logger) SetTransporter(transporter EntryTransporter)
func (logger *Logger) WithFields(fields Fields) Builder
func (logger *Logger) IsLevelEnabled(level Level) bool
```

### Builder 模式

```go
type Builder interface {
    Logln(level Level, args ...interface{})
    Log(level Level, args ...interface{})
    Logf(level Level, format string, args ...interface{})
    WithError(err error) Builder
    WithField(key string, value interface{}) Builder
    WithFields(fields Fields) Builder
}
```

## 高级用法

### 多传输器配置

```go
import (
    "os"
    "gitlab.ncader.com/common/log"
    "gitlab.ncader.com/common/log/filelog"
)

func setupLogging() {
    // 控制台传输器
    consoleTransporter := &log.Transporter{}
    consoleTransporter.SetOutput(os.Stdout)
    consoleTransporter.SetLevel(log.DebugLevel)
    consoleTransporter.SetFormatter(&log.LineFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
    
    // 文件传输器
    fileConfig := filelog.Config{
        Dir:    "./logs",
        File:   "app.log",
        Expire: 7 * 24 * time.Hour,
        Period: filelog.Day,
    }
    fileTransporter, _ := filelog.NewTransporter(fileConfig)
    
    // 使用复合传输器
    log.Reset(consoleTransporter, fileTransporter)
}
```

### 自定义格式化器

```go
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
    return []byte(fmt.Sprintf("[%s] %s: %s\n", 
        entry.Time.Format("15:04:05"), 
        strings.ToUpper(entry.Level.String()), 
        entry.Message)), nil
}

// 使用自定义格式化器
transporter.SetFormatter(&CustomFormatter{})
```

### 错误处理和链式调用

```go
log.WithFields(log.Fields{
    "module": "database",
    "action": "connect",
}).WithError(err).Error("数据库连接失败")
```

## 性能优化

该日志库采用了以下性能优化策略：

1. **对象池**: 使用 `sync.Pool` 复用 `EntryBuilder` 对象
2. **原子操作**: 日志级别使用原子操作，避免锁竞争
3. **延迟格式化**: 只有在需要输出时才进行格式化
4. **级别检查**: 在格式化前先检查日志级别

## 许可证

此项目使用内部许可证。

## 贡献

欢迎提交问题和功能请求到项目仓库。