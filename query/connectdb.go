package query

import (
	"database/sql"
)
func ConnectMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Quan@123@tcp(127.0.0.1:3306)/QLSV")
	if err != nil {
	  panic(err.Error())
	}
	err = db.Ping()
	if err != nil{
	  panic(err.Error())
	} 
	return db
	
}