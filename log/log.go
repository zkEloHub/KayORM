package log

import (
	"io"
	"log"
	"os"
	"sync"
)

// Lshortfile => 显示文件名和代码行号
// 不同层级日志显示不同的颜色

var (
	errorLog = log.New(os.Stdout, "\033[31m[error ]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods (暴露 Error，Errorf，Info，Infof 4个方法)
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
// iota: 0, (**counter** in const)
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if level > ErrorLevel {
		errorLog.SetOutput(io.Discard)
	}

	if level > InfoLevel {
		errorLog.SetOutput(io.Discard)
	}
}


