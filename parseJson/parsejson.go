package parseJson
import(
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Students struct{
	Surname string `json:"surname"`
	Name string `"json:name"`
	Birthday string `json:"birthday"`
	Address Address `json:"address"`
	Phone string `json:"phone"`
	Mark Mark `json:"mark"`
}
type Address struct{
	Apartment_number string `json:"apartment_number"`
	Ward string `json:"ward"`
	District string `json:"district"`
	City string `json:"city"`
}
type Mark struct{
	Math float64 `"json:math"`
	Phisical float64 `json:"phisical"`
	Chemistry float64 `json:"chemistry"`
}

func ParseJson() {
	data, err :=ioutil.ReadFile("student2.json")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully readfile")


	var students []Students
	err = json.Unmarshal([]byte(data), &students)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(students)
}