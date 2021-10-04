package main

import "fmt"


func main(){
  var (
      value32 int32 = 32767
      value64 int64 = int64(value32)
  )

  // Invalid value.
  value8 := int8(value64)

  fmt.Println("value8:", value8)
  fmt.Println("value64: ", value64)

  fmt.Println(float32(value64))


  fmt.Println("---------------------------------------------------")


  var (
      firstName = "Mahrus"
      lastName = "Khomaini"
      age = 25
  )

  firstNameChar := firstName[0]

  fmt.Println("First name:", firstName)
  fmt.Println("LastName name:", lastName)
  fmt.Println("Fist name char:", string(firstNameChar))
  fmt.Println("Age:", age)

}
