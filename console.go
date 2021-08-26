package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 往终端打印
type ConsoleLogger struct {
	level Level
}

// NewConsoleLogger 文件日志结构体的构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &ConsoleLogger{
		level: logLevel,
	}
	return fl
}

// 将公用的记录日志的功能封装成一个单独的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志
	// 日志格式 [时间][文件：行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		nowStr, fileName, line, funcName, logLevelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg) // 利用fmt包将msg字符串写入f.file文件中
}

// Debug 方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info 方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warn 方法
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error 方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// Fatal 方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

// Close 操作系统的标准输入不需要关闭
func (f *ConsoleLogger) Close() {
}
