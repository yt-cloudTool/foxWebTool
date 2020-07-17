package section_upload

/* *************************************************************************************
							receive
************************************************************************************* */
// 文件目录预处理信息格式
type PretreatInfo struct {
	Uid         string // 文件唯一id, 文件的每个切片都要携带且uid相同
	Timestamp   int64  // 时间戳 文件上传时间,每个切片携带的Timestamp相同
	FileName    string // 文件名
	Section_num int64  // 文件切片数量
}

// 切片预处理返回数据格式
type PretreatReturn struct {
	Code   int64  // 返回状态代码 0失败 1成功
	TmpDir string // 创建的临时目录路径
}

// 接收文件及切片信息格式
type ReceiveInfo struct {
	Uid           string // 文件唯一id, 文件的每个切片都要携带且uid相同
	Timestamp     int64  // 时间戳 文件上传时间,每个切片携带的Timestamp相同
	FilePath      string // 前端文件传来的文件路径包含文件名
	TargetPath    string // 前端文件存储到服务器的目标路径
	FileName      string // 文件名
	FileSize      int64  // 文件大小
	Section_size  int64  // 文件切片大小
	Section_index int64  // 当前文件切片的序号 序号从0开始 用于给文件所有切片编号
	Section_num   int64  // 文件切片数量
}

/* *************************************************************************************
							combine
************************************************************************************* */
// 合并检查接收参数
type CombineCheckInfo struct {
	Uid         string
	Timestamp   int64
	FileName    string
	FileSize    int64
	Section_num int64
}

// 合并操作接收参数
type CombineInfo struct {
	Uid         string
	Timestamp   int64
	FileName    string
	FileSize    int64
	Section_num int64
	Section_size int64
	TargetPath string // 文件合并后放到此路径
}