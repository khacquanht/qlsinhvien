package main

import (  
    "fmt"
    "encoding/json"
)
type Address struct{
    Apartment_number string
    Wards string
    District string
    City string
}
type Mark struct{
    Math float64
    Physical float64
    Chemistry float64
}
type Student struct {
    Surname string
    Name string
    Birthday string
    Address  Address
    Phone string
    Mark  Mark
}
func main(){
  studentJson  := `[{"surname": "Lê Khắc",
  "name": "Quân",
  "birthday": "28/08/1996",
  "address": {
    "apartment_number": "",
    "ward": "Thạch Văn",
    "district": "Thạch Hà",
    "city": "Hà Tĩnh"
  },
  "phone": "0983536038",
  "mark": {
    "math": 8,
    "physical": 7,
    "chemistry": 8
  }
},
{"surname": "Nguyễn Thị",
  "name": "Bình",
  "birthday": "7/11/1995",
  "address": {
    "apartment_number": "28 Lâm Văn Bền",
    "ward": "Phường 8",
    "district": "Quận 7",
    "city": "Tp. Hồ Chí Minh"
  },
  "phone": "0978456258",
  "mark": {
    "math": 8.5,
    "physical": 7.2,
    "chemistry": 8.8
  }
},
{"surname": "Trần Văn",
  "name": "Sơn",
  "birthday": "17/05/1996",
  "address": {
    "apartment_number": "14 Thống nhất",
    "ward": "Phường 11",
    "district": "Quận Gò Vấp",
    "city": "Tp.Hồ Chí Minh"
  },
  "phone": "0917856954",
  "mark": {
    "math": 8.5,
    "physical": 9.5,
    "chemistry": 8.8
  }
},
{"surname": "Trịnh Duy",
  "name": "Khánh",
  "birthday": "20/10/1996",
  "address": {
    "apartment_number": "32 Nguyễn Thị Thập",
    "ward": "Phường 4",
    "district": "Quận 7",
    "city": "Tp.Hồ Chí Minh"
  },
  "phone": "0988890099",
  "mark": {
    "math": 8,
    "physical": 7.5,
    "chemistry": 7
  }
}]`
                  
  var students []Student
  json.Unmarshal([]byte (studentJson), &students)
    fmt.Println("\n", students)
    
} 


