package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type Students struct{
  ID_SV string
  Surname string
  Name string
  Birthday string
  Phone string
}
type Address struct{
  ID_SV string
  Apartment_number string
  Ward string
  District string
  City string
}
type Mark struct{
  ID_SV string
  Math float64
  Phisical float64
  Chemistry float64
}
func main(){
  db, err := sql.Open("mysql", "root:Quan@123@tcp(127.0.0.1:3306)/QLSV")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  err = db.Ping()
  if err != nil{
    panic(err.Error())
  } 
  results, err := db.Query("SELECT* FROM QLSV.Students WHERE YEAR(Birthday) = 1996")
  if err != nil{
    panic(err.Error())
  }
  fmt.Println("Thông tin các sinh viên sinh năm 1996")
  for results.Next(){
    var students Students
    err = results.Scan(&students.ID_SV, &students.Surname, &students.Name, &students.Birthday, &students.Phone)
    if err != nil{
      panic(err.Error())
    }
    
    fmt.Println(students)
  }
  fmt.Println("\n")
  var district ="Quận 7" 
  results, err = db.Query("SELECT * FROM QLSV.Students INNER JOIN QLSV.Address ON QLSV.Students.ID_SV = QLSV.Address.ID_SV AND District = ?",district)
  if err != nil{
    panic(err.Error())
  }
  fmt.Println("Thông tin các sinh viên có địa chỉ ở Quận 7")
  for results.Next(){
    var students Students
    var address Address
    err = results.Scan(&students.ID_SV, &students.Surname, &students.Name, &students.Birthday, &students.Phone, &address.ID_SV,
       &address.Apartment_number, &address.Ward, &address.District, &address.City)
    if err != nil{
      panic(err.Error())
    }
    
    fmt.Println(students,address)
  }
  fmt.Println("\n")
  results, err = db.Query("SELECT* FROM QLSV.Students AS Student LEFT JOIN QLSV.Mark AS Mark ON Student.ID_SV = Mark.ID_SV WHERE (Mark.Math + Mark.Chemistry + Mark.Physical)/3 >=8")
   
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



