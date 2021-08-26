package mylogger

import (
	"path"
	"runtime"
)

// 存放一些共用的工具函数
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	// 从filename（文件全路径）中剥离出文件名
	fileName = path.Base(fileName)
	// 根据pc拿到函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}
