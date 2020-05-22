package query

import (
	"fmt"
)
type Address struct{
	ID_SV string
	Apartment_number string
	Ward string
	District string
	City string
  }
func Baitap2()  {
db := ConnectMysql()
var district ="%Quận 7%" 
results, err := db.Query("SELECT * FROM QLSV.Students INNER JOIN QLSV.Address ON QLSV.Students.ID_SV = QLSV.Address.ID_SV AND District LIKE ?",district)
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
}