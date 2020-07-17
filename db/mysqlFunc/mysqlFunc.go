package mysqlFunc

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var mysql_cfg = mysql.NewConfig()

func init() {
	mysql_cfg.User = "root"
	mysql_cfg.Passwd = "root"
	mysql_cfg.Net = "tcp"
	mysql_cfg.Addr = "localhost:3306"
	mysql_cfg.DBName = "forum"

}

func Link() *sql.DB {
	db, err := sql.Open("mysql", mysql_cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
