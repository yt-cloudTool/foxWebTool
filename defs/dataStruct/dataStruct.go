package dataStruct

// 请求返回
type RetJson struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data",omitempty`
}

// *************************************************************************
//					upload
// *************************************************************************
// ( 上传 ) 接收文件描述信息 返回数据
type RetJsonUploadDes struct {
	Uid string		`json:"uid"`
	Ret int64		`json:"ret"`
	Msg string		`json:"msg"`
	Dir string		`json:"dir"`
	Timestamp int64 `json:"timestamp"`
}

// *************************************************************************
//					account
// *************************************************************************
// ( 注册 ) 请求返回中的数据
type RetJsonRegister struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Email  string `json:"email"`
}

