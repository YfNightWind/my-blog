package mylog

import (
	"fmt"
)

func Init(serverName string) {
	_ = SetLogger(AdapterFile, fmt.Sprintf(`{"filename":"logs/%v.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`, serverName))
	_ = SetLogger(AdapterConsole)
	EnableFuncCallDepth(true)
	SetLogFuncCallDepth(3)
	SetShowFilePathLevel(LevelError, LevelWarning, LevelDebug)
	Async(1e3)
}
