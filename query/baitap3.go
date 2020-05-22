package query

import (
	"fmt"
)
type Mark struct{
	ID_SV string
	Math float64
	Phisical float64
	Chemistry float64
  }
func Baitap3(){
	db := ConnectMysql()
	results, err := db.Query("SELECT* FROM QLSV.Students AS Student LEFT JOIN QLSV.Mark AS Mark ON Student.ID_SV = Mark.ID_SV WHERE (Mark.Math + Mark.Chemistry + Mark.Physical)/3 >=8")
	 
	if err != nil{
	  panic(err.Error())
	}
	fmt.Println("Thông tin các sinh viên được chọn vào lớp Cử Nhân Tài Năng")
	for results.Next(){
	  var students Students
	  var mark Mark
	  err = results.Scan(&students.ID_SV, &students.Surname, &students.Name, &students.Birthday, &students.Phone, &mark.ID_SV,
		 &mark.Math, &mark.Phisical, &mark.Chemistry)
	  if err != nil{
		panic(err.Error())
	  }
	  
	  fmt.Println(students,mark)
	}
  }
