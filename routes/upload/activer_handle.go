// ===============================================================================
//								调试
// ===============================================================================

package upload_concurrent

import (
	"net/http"
	suplod "foxWebTool/utils/upload/section_upload"
	// global_id "foxWebTool/utils/global_id"
	ds "foxWebTool/defs/dataStruct"
	"encoding/json"
	// "mime/multipart"
	"fmt"
	// "time"
	"strconv"
	// "io"
)

// **********************************************************************
//				设置
// **********************************************************************
func SetActiver (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 0, Msg: "请求方法错误",})
		res.Write(ret_json)
		return
	}
	
	// 参数
	var key    = req.FormValue("key")
	var val, _ = strconv.ParseInt(req.FormValue("val"), 10, 64)
	
	err := suplod.Activer.Set(key, val)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -1, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	
	ret_json, _ := json.Marshal(ds.RetJson {Ret: 1, Msg: "ok"})
	res.Write(ret_json)
	return	
}

// **********************************************************************
//				获取
// **********************************************************************
func GetActiver (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 0, Msg: "请求方法错误",})
		res.Write(ret_json)
		return
	}
	
	// 参数
	var key = req.FormValue("key")
	
	retInt64, err := suplod.Activer.Get(key)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -1, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	
	ret_json, _ := json.Marshal(ds.RetJson {Ret: 1, Msg: "ok", Data: retInt64})
	res.Write(ret_json)
	return	
}

// **********************************************************************
//				删除
// **********************************************************************
func DelActiver (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 0, Msg: "请求方法错误",})
		res.Write(ret_json)
		return
	}
	
	// 参数
	var key = req.FormValue("key")
	
	suplod.Activer.Del(key)
	retInt64 := suplod.Activer.Len()
	ret_json, _ := json.Marshal(ds.RetJson {Ret: 1, Msg: "ok", Data: retInt64})
	res.Write(ret_json)
	return	
}