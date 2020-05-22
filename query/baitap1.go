package query

import (
	"fmt"
)
type Students struct{
	ID_SV string
	Surname string
	Name string
	Birthday string
	Phone string
  }
func Baitap1() {
db :=ConnectMysql()
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
}