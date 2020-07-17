package main

import (
	"fmt"
	// pool "foxProject/db/mysqlFunc"
	// "log"
	"net/http"
	"os"
	"path"

	"foxWebTool/routes/account"
	"foxWebTool/routes/test"
	"foxWebTool/routes/upload"
)

type Http_handler struct{}

// 程序所在路径
var curAppPath string

func init () {
	path, _ :=  os.Getwd()
	curAppPath = path
	fmt.Println("curAppPath =>", curAppPath)
}

// -----------------------------------------------------------------------------------
//							routes
// -----------------------------------------------------------------------------------
// 返回路由
func (this Http_handler) ROUTES() {
	fmt.Println("静态资源 =>", path.Join(curAppPath, static_dir))
	// -------------------------- 静态资源 --------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(curAppPath, static_dir)))))
	
	// -------------------------- 基本 --------------------------------
	http.HandleFunc("/test", test.Test)
	http.HandleFunc("/account/register", account.Register) // 注册
	
	// -------------------------- 文件上传 -----------------------------
	http.HandleFunc("/upload_concurrent/filedescribe", upload_concurrent.FileDescribe) // 获取文件描述信息
	http.HandleFunc("/upload_concurrent/filereceive", upload_concurrent.FileReceive) // 接收文件切片
	
	// -------------------------- 上传限制查询 (调试) --------------------	
	http.HandleFunc("/upload_concurrent/set_activer", upload_concurrent.SetActiver) // 设置活动
	http.HandleFunc("/upload_concurrent/get_activer", upload_concurrent.GetActiver) // 获取活动
	http.HandleFunc("/upload_concurrent/del_activer", upload_concurrent.DelActiver) // 删除活动
}
