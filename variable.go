package main

import "fmt"


func main(){
  var name string
  var age int8

  name = "Mahrus Khomaini"
  age = 25

  fmt.Println("Name:", name)
  fmt.Println("Age:", age)

  fmt.Println("---------------------------------------------------")


  var newName = "New Mahrus Khomaini"
  var newAge = 26

  fmt.Println("New Name:", newName)
  fmt.Println("New Age:", newAge)

  fmt.Println("---------------------------------------------------")


  firstName := "Mahrus"
  lastName := "Khomaini"

  fmt.Println("First name:", firstName)
  fmt.Println("Last name:", lastName)


  fmt.Println("---------------------------------------------------")

  var (
      fullName string = "Mahrus Khomaini"
      currentAge int8 = 25
  )

  fmt.Println("Fullname: ", fullName)
  fmt.Println("Current Age:", currentAge)

}
