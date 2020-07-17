package dbops

import (
	"errors"
	"fmt"
	pool "foxWebTool/db/mysqlFunc"
	crypto "foxWebTool/defs/encryption"
	// "log"
)

// 注册入参定义
type RegisterParams struct {
	Name       string
	Passwd     string
	Email      string
	CreateTime int64
	Lastlogin  int64
}

// 注册出参定义
type RegisterRet struct {
	Ret int
	Msg string
}

// 注册
func Register(param RegisterParams) (RegisterRet, error) {

	// ----------------------- 事务 start -------------------------------------------

	var db = pool.Link()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return RegisterRet{Ret: 0, Msg: fmt.Sprintf("%s", err)}, err
	}
	db_prepare, err := tx.Prepare("insert into users(name, passwd, email, create_time, lastlogin) values (?, ?, ?, ?, ?)")
	defer db_prepare.Close()
	if err != nil {
		tx.Rollback()
		return RegisterRet{Ret: 0, Msg: fmt.Sprintf("%s", err)}, err
	}
	db_exec_ret, err := db_prepare.Exec(param.Name, crypto.Generate(param.Passwd), param.Email, param.CreateTime, param.Lastlogin)
	if err != nil {
		tx.Rollback()
		return RegisterRet{Ret: 0, Msg: fmt.Sprintf("%s", err)}, err
	}
	// 判断是否插入成功
	row_num, err := db_exec_ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		return RegisterRet{Ret: 0, Msg: fmt.Sprintf("%s", err)}, err
	}
	if row_num > 0 {
		tx.Commit()
		return RegisterRet{Ret: 1, Msg: ""}, nil
	} else {
		tx.Rollback()
		return RegisterRet{Ret: 0, Msg: "数据库插入数量err"}, errors.New("DATABASE INSERT RESULTS ROWS LESS")
	}

	// ----------------------- 事务 end -------------------------------------------

}
