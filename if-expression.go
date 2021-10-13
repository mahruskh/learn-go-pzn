package main

import "fmt"


func main(){
  firstName := "Mahrus"
  // lastName := "Khomaini"
  // age := 25

  if firstName == "Mahrus" {
    fmt.Println("Your first name is", firstName)
  }

  if firstName == "mahrus" {
    fmt.Println("Not execute this block.")
  }


  if firstName == "Budi" {
    fmt.Println("I am", firstName)
  } else {
    fmt.Println("Who the hell are you?")
  }

  firstName = "Andre"
  if firstName == "Mahrus" || firstName == "Andre" {
    fmt.Println("Got you", firstName)
  }

  // Maye be code below will error number data type in several case.
  var score = 90
  if score >= 95 {
    fmt.Println("Perfect!")
  } else if score >= 90 && score < 95 {
    fmt.Println("Great!")
  } else if score >= 85 && score < 90 {
    fmt.Println("Good!")
  } else if score >= 75 && score < 85 {
    fmt.Println("Just Passed!")
  } else {
    fmt.Println("Are you kidding me?")
  }


  // Cannot like this ??? like in python or PHP or others.
  // error: incompatible types in binary expression
  /*
  var value = 10
  if 5 < value < 20 {
    fmt.Println("Whoaaaa!!!")
  }
  */

  // WHY ??? Mayme i'm wrong in practice.
  // error: incompatible types in binary expression
  /*
  var total int32 = 50
  if int32(20) < value < int32(total) {
    fmt.Println("Horay!!!")
  }
  */

  fmt.Println("-----------------------------------------------------------")

  // Short statement
  fullName := "Kevin De Bruyne"
  if length := len(fullName); length > 4 {
      fmt.Println("To length!!!")
  } else {
    fmt.Println("No bad.")
  }

  // Why just use this ??????, need another example!!!
  if len(fullName) > 4 {
      fmt.Println("Hola!!!")
  }

  fmt.Println("-----------------------------------------------------------")

  // Testing not comprison
  var total = 10

  // error: expected boolean expression
  // Go cannot ternary operation????
  // if total {
  //   fmt.Println(total)
  // }

  anyTotal := total != 0
  if anyTotal {
    fmt.Println(total)
  }

  fmt.Println("-----------------------------------------------------------")

  firstName = "Mahrus"

  // error: expected boolean expression
  // if firstName {
  //   fmt.Println(firstName)
  // }

  if len(firstName) > 0{
      fmt.Println(firstName)
  }

  fmt.Println("-----------------------------------------------------------")

  // Testing with Array
  var cities [10]string
  fmt.Println(cities)

  // error: expected boolean expression
  // if cities {
  //   fmt.Println(cities)
  // }

  // Wrong logic if using Array len
  if len(cities) > 0 {
    fmt.Println(cities)
  }

  cities[0] = "Malang"
  // still error: expected boolean expression
  // if cities {
  //   fmt.Println(cities)
  // }

  fmt.Println("-----------------------------------------------------------")

  // Testing Slice
  fruits := []string{
      "Apple", "Banana", "Watermelon",
  }
  animals := make([]string, 0, 10)

  // error: expected boolean expression
  // if fruits {
  //   fmt.Println(cities)
  // }

  if len(fruits) > 0 {
    fmt.Println(fruits)
  }

  fmt.Println(animals, len(animals), cap(animals))
  if len(animals) > 0 {
    fmt.Println(animals)
  }


  /* Conclution
  - Condition must only be in comparison values.
  - Cannot ternary expression like in python ???
  */

}
