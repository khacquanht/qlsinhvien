package main

import (
	"qlsv/parseJson"
  "qlsv/query"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
 parseJson.ParseJson()

 query.Baitap1()

 query.Baitap2()

 query.Baitap3()
}



