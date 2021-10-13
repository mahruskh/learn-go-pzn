package main

import "fmt"

func main(){
  category := "politic"

  switch category{
  case "economy":
    fmt.Println("ECONOMY", category)
  case "sport":
    fmt.Println("SPORT:", category)
  case "politic":
    fmt.Println("POLITIC", category)
  default:
    fmt.Println("DEFAULT News")
  }

  fmt.Println("--------------------------------------------------------")

  name := "Mahrus"

  switch length := len(name); length > 5{
  case true:
    fmt.Println("ONE")
  case false:
    fmt.Println("TWO")
  }

  fmt.Println("--------------------------------------------------------")

  total := len(name)
  switch{
  case total > 10:
    fmt.Println("Too long!")
  case total > 5:
    fmt.Println("Long!")
  default:
    fmt.Println("Good.")
  }


}
