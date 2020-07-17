package test

import (
	"fmt"
	pool "foxWebTool/db/mysqlFunc"
	"log"
	"net/http"
)

// -----------------------------------------------------------------------------------
//							测试接口
// -----------------------------------------------------------------------------------
func Test(res http.ResponseWriter, req *http.Request) {
	var dblink = pool.Link()
	defer dblink.Close()
	rows, err := dblink.Query("select `id` from `users`")
	if err != nil {
		fmt.Fprintf(res, "db 	 query err")
	}

	var st string
	for rows.Next() {
		if err := rows.Scan(&st); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(st)
	}
	fmt.Fprintf(res, st)
}
