/* *****************************************************************************
							用于配置文件上传
***************************************************************************** */
package section_upload

import (
	"os/user"
)

// 临时文件设置为用户主目录
var cur_user, _ = user.Current()
var homeDir = cur_user.HomeDir

// 存储临时文件的目录
var tmpDirPath string = homeDir + "/.cache/upload_tmp"
// 文件目标路径
var targetDirPath string = homeDir + "/serverUpload"

// 同时上传任务数限制
var TaskLimitNum int64 = 1024
// 每个上传任务切片大小限制 (byte)
var TaskCountNum int64 = 4 * 1024 * 1024