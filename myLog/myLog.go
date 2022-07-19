package myLog

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)


var Logger  *log.Logger

func init() {
	// 配置日志
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	fileName := strconv.FormatInt(time.Now().Unix(), 10)
	writer3, err := os.OpenFile(fileName + ".txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	Logger = log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
}


