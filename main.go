package main

import (  
    "fmt"
    "io/ioutil"
)
func main(){
  data, err := ioutil.ReadFile("/home/quan/go/src/qlsv/student2.json")
  
  if err != nil {
    fmt.Println(err)
  }
  fmt.Print(string(data))
}



