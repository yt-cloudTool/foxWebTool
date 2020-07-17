package upload_concurrent

import (
	"net/http"
	suplod "foxWebTool/utils/upload/section_upload"
	ds "foxWebTool/defs/dataStruct"
	global_id "foxWebTool/utils/global_id"
	"encoding/json"
	// "mime/multipart"
	"fmt"
	"time"
	"strconv"
	// "io"
)

// 接收文件描述信息
/*
	1.接收client的文件描述信息 { filename 文件名, sectionnum 文件切片数 }
	2.生成uuid
	3.返回uuid,目标路径给client
*/
func FileDescribe (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 0, Msg: "请求方法错误",})
		res.Write(ret_json)
		return
	}
	
	// 客户端传来的描述信息
	// Uid         string // 文件唯一id, 文件的每个切片都要携带且uid相同
	// Timestamp   int64  // 时间戳 文件上传时间,每个切片携带的Timestamp相同
	// FileName    string // 文件名
	// Section_num int64  // 文件切片数量
	sectionNum, err := strconv.ParseInt(req.FormValue("sectionnum"), 10, 64)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJsonUploadDes { Ret: -1, Msg: fmt.Sprintf("%s",err) })
		res.Write(ret_json)
		return
	}
	
	// 生成uid
	var fuid = global_id.Generate()
	// 操作时间
	var timestamp = time.Now().UnixNano() / 1e6
	
	var clientDatainfo = suplod.PretreatInfo{
		Uid: 			fuid,
		Timestamp: 		timestamp,
		FileName:		req.FormValue("filename"),
		Section_num:	sectionNum,
	}
	
	// 上传用临时文件初始化
	pretreatRet, err := suplod.PretreatTmp(&clientDatainfo)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJsonUploadDes { Ret: -2, Msg: fmt.Sprintf("%s",err), })
		res.Write(ret_json)
		return
	}

	// 返回
	ret_json, _ := json.Marshal(ds.RetJsonUploadDes { Uid: fuid, Ret: 1, Msg: "ok", Dir: pretreatRet.TmpDir, Timestamp: timestamp })
	res.WriteHeader(http.StatusOK)
	res.Write(ret_json)
}

// 接收文件切片 | 合并文件
/*
参数
{
	uid
	timestamp
	dir
	filename
	filesize
	sectonsize
	sectionindex
	sectionnum
	file(文件)	
}
*/
func FileReceive (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 0, Msg: "请求方法错误",})
		res.Write(ret_json)
		return
	}
	
	// 文件切片数据
	formFile, _, err := req.FormFile("file")
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -1, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	
	// 获取参数
	form_uid				:= req.PostFormValue("uid")
	form_timestamp, _		:= strconv.ParseInt(req.PostFormValue("timestamp"), 10, 64)
	form_dir 				:= req.PostFormValue("dir")
	form_filename 			:= req.PostFormValue("filename")
	form_filesize, _ 		:= strconv.ParseInt(req.PostFormValue("filesize"), 10, 64)
	form_sectionsize, _ 	:= strconv.ParseInt(req.PostFormValue("sectionsize"), 10, 64)
	form_sectionindex, _ 	:= strconv.ParseInt(req.PostFormValue("sectionindex"), 10, 64)
	form_sectionnum, _ 		:= strconv.ParseInt(req.PostFormValue("sectionnum"), 10, 64)
	
	// 接收文件切片描述
	// Uid           string // 文件唯一id, 文件的每个切片都要携带且uid相同
	// Timestamp     int64  // 时间戳 文件上传时间,每个切片携带的Timestamp相同
	// FilePath      string // 前端文件传来的文件路径包含文件名
	// TargetPath    string // 前端文件存储到服务器的目标路径
	// FileName      string // 文件名
	// FileSize      int64  // 文件大小
	// Section_size  int64  // 文件切片大小
	// Section_index int64  // 当前文件切片的序号 序号从0开始 用于给文件所有切片编号
	// Section_num   int64  // 文件切片数量
	var clientDatainfo = suplod.ReceiveInfo {
		Uid: 			form_uid,
		Timestamp: 		form_timestamp,
		TargetPath:		form_dir,
		FileName: 		form_filename,
		FileSize: 		form_filesize,
		Section_size: 	form_sectionsize,
		Section_index: 	form_sectionindex,
		Section_num: 	form_sectionnum,
	}
	
	_, err = suplod.Receive(&clientDatainfo, formFile)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -2, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	
	fmt.Println("接收参数 =>", form_sectionsize, form_sectionindex)
	
	// ===========================================================================
	//					判断是否接收完毕
	// ===========================================================================
	// **************************************************
	// 设置记录活动 (以后sectionNum改成: 剩余应传切片个数)
	// **************************************************
	activerInt64, err := suplod.Activer.Get(form_uid)
	if err != nil {
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -5, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	if activerInt64 > 0 {
		// 返回结果
		ret_json, _ := json.Marshal(ds.RetJson {Ret: 2, Msg: "切片接收ok"})
		res.Write(ret_json)
		return
	}
	// ===========================================================================
	
	// 判断是否接收完毕
	_, err = suplod.CombineCheck(suplod.CombineCheckInfo {
		Uid: 			form_uid,
		Timestamp:		form_timestamp,
		FileName:		form_filename,
		FileSize:		form_filesize,
		Section_num: 	form_sectionnum,
	})
	fmt.Println("判断是否接收完毕 =>", err)
	if err != nil {
		// 返回结果
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -3, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	// 合并操作
	_, err = suplod.Combine(suplod.CombineInfo {
		Uid: 			form_uid,
		Timestamp:		form_timestamp,
		FileName:		form_filename,
		FileSize:		form_filesize,
		Section_num: 	form_sectionnum,
		TargetPath:		form_dir,
	})
	if err != nil {
		fmt.Println("文件合并失败!", err)
		ret_json, _ := json.Marshal(ds.RetJson {Ret: -4, Msg: fmt.Sprintf("%s", err)})
		res.Write(ret_json)
		return
	}
	
	// 返回结果
	ret_json, _ := json.Marshal(ds.RetJson {Ret: 1, Msg: "文件接收完毕"})
	res.Write(ret_json)
}