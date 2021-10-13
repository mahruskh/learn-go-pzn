package main

import "fmt"

func main(){
  counter := 1
  for counter <= 10 {
    fmt.Println("Loop:", counter)
    counter++
  }

  fmt.Println("----------------------------------------------")

  for index := 1; index <= 15; index++ {
    fmt.Println("Index", index)
  }

  fmt.Println("-------------------basic loop --------------------")

  friends := []string{"Budi", "Doni", "Bella", "Eko"}

  // Basic ways
  for i := 0; i < len(friends); i++ {
    fmt.Println(i, friends[i])
  }

  fmt.Println("------------------ range ---------------------")

  // Using range
  for index, value := range friends {
    fmt.Println(index, value)
  }

  fmt.Println("------------------ unused variable ---------------------")
  for _, value := range friends {
    fmt.Println(value)
  }

  fmt.Println("------------------ Loop map ---------------------")

  person := make(map[string]string)
  person["name"] = "Mahrus Khomaini"
  person["city"] = "Malang"
  person["learn"] = "Go lang"

  for key, value := range person {
    fmt.Println(key, value)
  }

}
