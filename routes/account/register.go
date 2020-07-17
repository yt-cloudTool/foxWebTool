package account

import (
	"encoding/json"
	"fmt"
	dbops_account "foxWebTool/db/dbops/account"
	datas "foxWebTool/defs/dataStruct"

	// crypto "foxProject/defs/encryption"
	// "log"
	"net/http"
	"time"
)

/*
	注册 (POST)
	参数 ->
		name
		passwd
*/
func Register(res http.ResponseWriter, req *http.Request) {
	// 判断请求方法
	var req_method = req.Method
	if req_method != "POST" {
		json, _ := json.Marshal(datas.RetJson{Ret: 0, Msg: "请求方法错误"})
		fmt.Fprintf(res, string(json))
		return
	}

	// 获取参数
	var name = req.FormValue("name")
	var passwd = req.FormValue("passwd")
	var email = req.FormValue("email")

	// 判断参数
	if passwd == "" || name == "" || email == "" {
		json, _ := json.Marshal(datas.RetJson{Ret: 0, Msg: "用户名密码不能为空"})
		fmt.Fprintf(res, string(json))
		return
	}

	// 当前时间
	var cur_time = time.Now().UnixNano() / 1e6 // 毫秒

	// 存入库
	DB_ret, err := dbops_account.Register(dbops_account.RegisterParams{
		Name:       name,
		Passwd:     passwd,
		Email:      email,
		CreateTime: cur_time,
		Lastlogin:  cur_time,
	})
	if err != nil {
		fmt.Println("dbret =>", err)
		json, _ := json.Marshal(datas.RetJson{Ret: DB_ret.Ret, Msg: DB_ret.Msg})
		fmt.Fprintf(res, string(json))
		return
	} else {
		// 返回正确结果
		json, _ := json.Marshal(datas.RetJson{Ret: 1, Msg: "ok"})
		fmt.Fprintf(res, string(json))
		return
	}
}
