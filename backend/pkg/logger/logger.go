package logger

import (
	"log"
	"os"
)

var std = log.New(os.Stdout, "", log.LstdFlags)

// Info 输出 INFO 日志。
func Info(v ...any) { std.Println(append([]any{"[INFO]"}, v...)...) }

// Error 输出 ERROR 日志。
func Error(v ...any) { std.Println(append([]any{"[ERROR]"}, v...)...) }
