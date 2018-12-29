package main

import (
	"net/http"
	"log"
	"time"
	"os"
	"path"
	"os/exec"
	"weblog"
)

//路由
var router map[string]func(r *http.Request) string = map[string]func(r *http.Request) string {
	"/":home,
}

var exec_dir string
var last_time int64;
var loger *log.Logger;

func home(r *http.Request) string {
	return r.UserAgent()
}

func logtest (r *http.Request) string {
	return "sdkf"
}

func main() {
	for route,handle := range router {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			loger = weblog.Loger(exec_dir);
			loger.Printf("访问路由开始%s time => %s,params=> %v",route,time.Now().Format("20181210 00:00:00"),r);
			handle(r);
			loger.Printf("访问路由结束%s time => %s,params=> %v",route,time.Now().Format("20181210 00:00:00"),r);
		})
	}

	loger.Print(http.ListenAndServe(":8080", nil))
	loger.Output(2,"web关闭");
	weblog.Close();
}

func init() {
	var err error;
	//获取路径
	exec_dir,err = exec.LookPath(os.Args[0]);
	if  err != nil {
		panic("找不到程序");
	}
	exec_dir = path.Clean(path.Dir(exec_dir)) + "/logs/";
	loger = weblog.Loger(exec_dir);
	loger.Output(2,"web启动");
}

