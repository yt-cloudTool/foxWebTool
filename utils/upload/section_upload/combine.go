/* *****************************************************************************
							组合文件切片
***************************************************************************** */
package section_upload

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"errors"
	"strings"
)

// 临时文件目录结构
// ->.cache/
// 	 ->upload_tmp/
//         ->timestamp_filename_sectionNum_uid/
//		   ->sectionInd_sectionSize

// 检查临时文件是否已存储完毕(检查是否可以合并)
func CombineCheck(cbc_info CombineCheckInfo) (int64, error) {
	// 参数
	var timestamp = cbc_info.Timestamp
	var filename = cbc_info.FileName
	// var filesize = cbc_info.FileSize
	var sectionNum = cbc_info.Section_num
	var uid = cbc_info.Uid

	var tmpFileDirPath = fmt.Sprintf("%s%s%d%s%s%s%d%s%s", tmpDirPath, "/", timestamp, "_", filename, "_", sectionNum, "_", uid)
	// 1. 判断临时文件切片目录是否存在
	_, err := os.Stat(tmpFileDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return -1, err
		}
	}

	// 2. 检查临时文件切片目录中切片数量是否足够
	fileList, err := ioutil.ReadDir(tmpFileDirPath)
	if err != nil {
		return -2, err
	}
	if int64(len(fileList)) < int64(cbc_info.Section_num) {
		return -3, errors.New("文件切片不足")
	}

	// 3. 检查每个切片是否下载完整
	for _, v := range fileList {
		var split_name_arr = strings.Split(v.Name(), "_")
		if len(split_name_arr) == 2 {
			var v_section_ind, _ = strconv.ParseInt(split_name_arr[0], 10, 64)
			var v_section_size, _ = strconv.ParseInt(split_name_arr[1], 10, 64)
			
			fmt.Println("inner 文件大小预期 =>", v_section_size, v.Size())

			// 检测文件实际大小和文件额定大小是否相等
			// 如果不等
			if v_section_size != int64(v.Size()) {
				return -4, errors.New(fmt.Sprintf("%s%d", "切片大小不符预期", v_section_ind))
			}
		}
	}
	
	return 1, nil
}

// 组合切片
func Combine(cb_info CombineInfo) (int64, error) {
	// 参数
	var timestamp = cb_info.Timestamp
	var filename = cb_info.FileName
	var filesize = cb_info.FileSize
	var sectionNum = cb_info.Section_num
	var uid = cb_info.Uid
	
	// **************************************************
	// 设置记录活动 删除
	// **************************************************
	defer Activer.Del(uid)

	// 目标临时文件目录路径
	var tmpFileDirPath = fmt.Sprintf("%s%s%d%s%s%s%d%s%s", tmpDirPath, "/", timestamp, "_", filename, "_", sectionNum, "_", uid)
	// 目标组合后存储文件路径
	var fileStorePath = targetDirPath + "/" + filename + "_" + uid

	// 1. 获取文件切片列表
	tmpfileList, err := ioutil.ReadDir(tmpFileDirPath)
	if err != nil {
		return -1, err
	}
	
	// 2. 创建目标文件
	_, err = os.Create(fileStorePath)
	if err != nil {
		return -2, err
	}
	
	// 3. 组合文件切片
	tarF_p, err := os.OpenFile(fileStorePath, os.O_RDWR, 0666)
	defer tarF_p.Close()
	if err != nil {
		return -3, err
	}
	// 追加每个文件切片
	for _, v := range tmpfileList {
		
		tmpF, err := os.OpenFile(tmpFileDirPath + "/" + v.Name(), os.O_RDONLY, 0666)
		if err != nil {
			return -4, err
		}
		
		fBuffer := make([]byte, v.Size())
		var step = 0
		for {
			step = step + 1
			tailN, err := tmpF.Read(fBuffer)
			if err != nil {
				return -5, err
			}
			if tailN == 0 {
				break
			}
			// 查找目标文件末尾
			fTail, err := tarF_p.Seek(0, os.SEEK_END)
			if err != nil {
				return -6, err
			}
			// 写入到目标文件
			_, err = tarF_p.WriteAt(fBuffer, fTail)
			if err != nil {
				return -7, err
			}
			
			fBuffer = fBuffer[:0]
		}
		tmpF.Close()	
	}
	
	// 4. 判断组合后文件大小是否等于额定大小
	fileInfo, err := os.Stat(fileStorePath)
	if err != nil {
		return -7, err
	}
	fmt.Println("判断大小 =>", fileInfo.Size(), int64(filesize))
	if int64(fileInfo.Size()) != int64(filesize) {
		return -8, errors.New("文件大小和额定大小不同")
	}
	
	// 5. 删除此文件临时目录
	err = os.RemoveAll(tmpFileDirPath)
	if err != nil {
		return -9, err
	} 
	
	return 1, nil
}
