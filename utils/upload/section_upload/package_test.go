package section_upload

import (
	// "fmt"
	"testing"
	// "io/ioutil"
	// "strings"
	// "os"
	// "time"
)

// func TestCreateTmp(t *testing.T) {
// 	v := PretreatTmp(&PretreatInfo{Uid: "afsdfw", Timestamp: 1169936063})
// 	fmt.Println("v =>", v)
// }

func TestCombineCheck(t *testing.T) {
	// fileList, err := ioutil.ReadDir("/home/stu/.cache/upload_tmp/1594429997445_测试文件_5_100")
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// }

	// for _, v := range fileList {
	// 	f, err := os.OpenFile("/home/stu/.cache/upload_tmp/1594429997445_测试文件_5_100/" + v.Name(), os.O_RDWR, 0666)
	// 	defer f.Close()
	// 	if err != nil {
	// 		fmt.Println("file open err=>", err)
	// 	}
	// 	content, _ := ioutil.ReadAll(f)
	// 	fmt.Println("file inner =>", string(content))
	// }

	/* Uid         string
	Timestamp   int64
	FileName    string
	FileSize    int64
	Section_num int6*/
	
	// v := CombineCheck(CombineCheckInfo{
	// 	Uid:         "100",
	// 	Timestamp:   1594429997445,
	// 	FileName:    "测试文件",
	// 	FileSize:    69501,
	// 	Section_num: 5,
	// })

	// fmt.Println("v =>", v)
	
	// v, err := Combine(CombineInfo{
	// 		Timestamp: 1594429997445,
	// 		FileName: "测试文件",
	// 		FileSize: 20480,
	// 		Section_num: 5,
	// 		Uid: "100",
	// })
	// fmt.Println("v =>", v, "err =>", err)
}

func Testinit (t *testing.T) {
	
}