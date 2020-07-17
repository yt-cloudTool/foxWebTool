/* *****************************************************************************
							接收文件切片
***************************************************************************** */
package section_upload

import (
	"fmt"
	"io"
	// "net/http"
	"mime/multipart"
	"os"
	"errors"
	// "strings"
)

// 临时文件目录结构
// ->.cache/
// 	 ->upload_tmp/
//         ->timestamp_filename_sectionNum_uid/
//		   ->sectionInd_sectionSize

/*
	1. (->api) 文件接收第一步 -> 预先准备
*/
// ( 预处理 ) 创建上传时需要用到的目录及文件
// Uid         string // 文件唯一id, 文件的每个切片都要携带且uid相同
// Timestamp   int64  // 时间戳 文件上传时间,每个切片携带的Timestamp相同
// FileName    string // 文件名
// Section_num int64  // 文件切片数量
func PretreatTmp(pre_info *PretreatInfo) (PretreatReturn, error) {

	// **************************************************
	//		   1.判断上传临时目录
	// **************************************************
	_, err := os.Stat(tmpDirPath)
	if err != nil {
		// 如果目录不存在则创建
		if os.IsNotExist(err) {
			err := os.MkdirAll(tmpDirPath, os.ModePerm)
			if err != nil {
				fmt.Println("SECTION_UPLOAD receive.create err1 =>", err)
				return PretreatReturn{Code: 0, TmpDir: ""}, err
			}
		}
	}

	// 如果以上路径(1)存在或已创建则

	// **************************************************
	//		   2.判断文件临时目录
	// **************************************************

	// 上传文件临时目录 按uid建立目录 (里面为这个文件的切片)

	var tmpFileDir = fmt.Sprintf("%s%s%d%s%s%s%d%s%s", tmpDirPath, "/", pre_info.Timestamp, "_", pre_info.FileName, "_", pre_info.Section_num, "_", pre_info.Uid)

	_, err = os.Stat(tmpFileDir)
	if err != nil {
		// 如果目录不存在则创建
		if os.IsNotExist(err) {
			err := os.MkdirAll(tmpFileDir, os.ModePerm)
			if err != nil {
				fmt.Println("SECTION_UPLOAD receive.create err2 =>", err)
				return PretreatReturn{Code: -1, TmpDir: "" }, err
			}
		}
	}
	
	// **************************************************
	// 设置记录活动 (以后sectionNum改成: 剩余应传切片个数)
	// **************************************************
	err = Activer.Set(pre_info.Uid, pre_info.Section_num)
	if err != nil {
		return PretreatReturn{Code: -2, TmpDir: "" }, err
	}

	// 如果以上路径(2)存在或已创建则

	// 返回创建成功状态及路径
	return PretreatReturn{Code: 1, TmpDir: fmt.Sprintf("%s", tmpFileDir)}, nil
}

/*
	2. (->api) 文件接收第二步 -> 并发上传
*/
// 接收操作
func Receive(rec_info *ReceiveInfo, file multipart.File) (int64, error) {
	fmt.Println("receive_split =>")
	// 存到服务器的目标路径
	var target_path = rec_info.TargetPath
	// 目标文件名 -> 切片编号_切片大小
	var target_filename = fmt.Sprintf("%d%s%d", rec_info.Section_index, "_", rec_info.Section_size)

	// 创建临时文件 (覆盖)
	tmp_file, err := os.Create(target_path + "/" + target_filename)
	if err != nil {
		return -1, err
	}

	_, err = io.Copy(tmp_file, file)
	if err != nil {
		return -2, err
	}
	
	// **************************************************
	// 设置记录活动 已记录的活动数减1
	// **************************************************
	if Activer.HasKey(rec_info.Uid) == false {
		return -3, errors.New("no uid storage at activer")
	}
	arctiverGet64, err := Activer.Get(rec_info.Uid)
	arctiverGet64 = arctiverGet64 - 1
	if arctiverGet64 < 0 {
		return -4, errors.New("receive number are over")
	}
	err = Activer.Set(rec_info.Uid, arctiverGet64)
	if err != nil {
		return -5, err
	}

	// 返回成功
	return 1, nil
}
