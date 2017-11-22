package logger

import (
	"fmt"
	"time"
)

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_purple
	log_erro    = "[ERRO]"
	log_success = "[SUCCESS]"
	log_warning = "[WARNING]"
	log_info    = "[INFO]"
	log_system  = "[SYSTEM]"
)

func print(color uint8, tip string, format string, a ...interface{}) {
	prefix := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, tip)
	logContent := time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
	fmt.Println(logContent, fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	print(color_red, log_erro, format, a...)
}

func Success(format string, a ...interface{}) {
	print(color_green, log_success, format, a...)
}

func Warning(format string, a ...interface{}) {
	print(color_yellow, log_warning, format, a...)
}

func Info(format string, a ...interface{}) {
	print(color_blue, log_info, format, a...)
}

func System(format string, a ...interface{}) {
	print(color_purple, log_system, format, a...)
}
