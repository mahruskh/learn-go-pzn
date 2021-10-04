package main

import "fmt"

func main(){
  person := map[string]string{
      "firstName": "Mahrus",
      "lastName": "Khomaini",
      "address": "Malang",
  }

  fmt.Println(person)
  fmt.Println(person["firstName"])
  fmt.Println(person["lastName"])
  fmt.Println(person["address"])

  person["address"] = "Mlg"
  fmt.Println(person["address"])

  person["gender"] = "Male"
  fmt.Println(person)
  fmt.Println(person["gender"])

  // Test error unknown key
  // Not error, return none???
  fmt.Println(person["status"])

  fmt.Println("--------------------------------------------------------")

  // Testing remove
  fmt.Println(person, len(person))
  delete(person, "gender")
  fmt.Println(person, len(person))

  // Remove unknown key
  // Not error, None???
  delete(person, "gender")

  fmt.Println("--------------------------------------------------------")

  // Test new map
  who := make(map[string]string)
  fmt.Println(who)

  // Test error type
  var sayNumber = make(map[int]string)
  sayNumber[1] = "One"
  sayNumber[2] = "Two"
  sayNumber[3] = "Three"
  fmt.Println(sayNumber, len(sayNumber))
  fmt.Println(sayNumber[1])
  fmt.Println(sayNumber[2])
  fmt.Println(sayNumber[3])

}
