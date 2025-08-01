package log

import (
	"os"
)

type exitFunc func(int)

// Tracef 使用格式化字符串记录跟踪级别的日志
func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.Logf(TraceLevel, format, args...)
}

// Debugf 使用格式化字符串记录调试级别的日志
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logf(DebugLevel, format, args...)
}

// Infof 使用格式化字符串记录信息级别的日志
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logf(InfoLevel, format, args...)
}

// Warnf 使用格式化字符串记录警告级别的日志
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logf(WarnLevel, format, args...)
}

// Warningf 是 Warnf 的别名，使用格式化字符串记录警告级别的日志
func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf 使用格式化字符串记录错误级别的日志
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logf(ErrorLevel, format, args...)
}

// Fatalf 使用格式化字符串记录致命错误级别的日志，并退出程序
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Logf(FatalLevel, format, args...)
	logger.Exit(1)
}

// Panicf 使用格式化字符串记录panic级别的日志
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.Logf(PanicLevel, format, args...)
}

// Trace 记录跟踪级别的日志
func (logger *Logger) Trace(args ...interface{}) {
	logger.Log(TraceLevel, args...)
}

// Debug 记录调试级别的日志
func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(DebugLevel, args...)
}

// Info 记录信息级别的日志
func (logger *Logger) Info(args ...interface{}) {
	logger.Log(InfoLevel, args...)
}

// Warn 记录警告级别的日志
func (logger *Logger) Warn(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

// Warning 是 Warn 的别名，记录警告级别的日志
func (logger *Logger) Warning(args ...interface{}) {
	logger.Warn(args...)
}

// Error 记录错误级别的日志
func (logger *Logger) Error(args ...interface{}) {
	logger.Log(ErrorLevel, args...)
}

// Fatal 记录致命错误级别的日志，并退出程序
func (logger *Logger) Fatal(args ...interface{}) {
	logger.Log(FatalLevel, args...)
	logger.Exit(1)
}

// Panic 记录panic级别的日志
func (logger *Logger) Panic(args ...interface{}) {
	logger.Log(PanicLevel, args...)
}

// Traceln 记录跟踪级别的日志并自动添加换行符
func (logger *Logger) Traceln(args ...interface{}) {
	logger.Logln(TraceLevel, args...)
}

// Debugln 记录调试级别的日志并自动添加换行符
func (logger *Logger) Debugln(args ...interface{}) {
	logger.Logln(DebugLevel, args...)
}

// Infoln 记录信息级别的日志并自动添加换行符
func (logger *Logger) Infoln(args ...interface{}) {
	logger.Logln(InfoLevel, args...)
}

// Warnln 记录警告级别的日志并自动添加换行符
func (logger *Logger) Warnln(args ...interface{}) {
	logger.Logln(WarnLevel, args...)
}

// Warningln 是 Warnln 的别名，记录警告级别的日志并自动添加换行符
func (logger *Logger) Warningln(args ...interface{}) {
	logger.Warnln(args...)
}

// Errorln 记录错误级别的日志并自动添加换行符
func (logger *Logger) Errorln(args ...interface{}) {
	logger.Logln(ErrorLevel, args...)
}

// Fatalln 记录致命错误级别的日志并自动添加换行符，然后退出程序
func (logger *Logger) Fatalln(args ...interface{}) {
	logger.Logln(FatalLevel, args...)
	logger.Exit(1)
}

// Panicln 记录panic级别的日志并自动添加换行符
func (logger *Logger) Panicln(args ...interface{}) {
	logger.Logln(PanicLevel, args...)
}

// Exit 执行清理处理程序并退出程序
func (logger *Logger) Exit(code int) {
	runHandlers()
	if logger.exitFunc == nil {
		logger.exitFunc = os.Exit
	}
	logger.exitFunc(code)
}

// GetLevel 返回日志记录器的当前日志级别
func (logger *Logger) GetLevel() Level {
	return logger.level()
}
