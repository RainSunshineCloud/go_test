package weblog

import (
	"log"
	"os"
	"time"
	"fmt"
)

var last_time int64
var loger *log.Logger
var log_file *os.File

//获取记录loger句柄
func Loger(path string) *log.Logger {

	now_time := time.Now().Unix()
	if now_time-last_time <= 86400 {
		return loger
	}

	if log_file != nil {
		log_file.Close()
	}
	fmt.Print(path);
	last_time = now_time
	log_file = MkFileOrGetFile(path , time.Now().Format("20180102") + ".log")
	loger = log.New(log_file, "test_", log.Ltime|log.Lshortfile)

	return loger
}

//创建或获取文件句柄
func MkFileOrGetFile(path string,file_name string) *os.File {
	var err error
	mkDir(path);

	path = path + file_name;
	if _, err = os.Stat(path); err == nil {
		if log_file, err := os.OpenFile(path, os.SEEK_END, os.ModeAppend); err == nil {
			return log_file
		}
		panic("打开文件出错")
	}

	if os.IsExist(err) {
		panic("获取文档出错，可能是没有权限")
	}

	if log_file, err := os.Create(path); err == nil {
		return log_file
	}

	panic("创建文档出错")
}


func mkDir (path string) {
	var err error;

	if _, err = os.Stat(path); err == nil {
		return;
	}

	if os.IsExist(err) {
		panic("获取文档出错，可能是没有权限")
	}

	if err := os.MkdirAll(path,os.ModePerm); err != nil {
		panic("创建文件夹出错")
	}
}

func Close() {
	log_file.Close()
	loger = nil
}
