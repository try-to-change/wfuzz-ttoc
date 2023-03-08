package log

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// 在当前目录下创建 log 文件，若文件已存在则追加内容
	file, err := os.OpenFile("../out/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Can't create log file: ", err)
	}
	logger = log.New(file, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogError(err error) {
	logger.Println(err)
}
