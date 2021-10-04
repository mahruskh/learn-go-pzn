package main

import "fmt"


func main(){
  var friends [10]string

  // Nothing ???
  fmt.Println(friends[0])

  friends[0] = "Mahrus"
  friends[1] = "Khomaini"

  fmt.Println(friends[0])
  fmt.Println(friends[1])

  // Test assign more than available index
  // It's error.
  // friends[10] = "Roby"
  // fmt.Println(friends[10])

  // Assign skip index, no assing to index 3, but 4.
  friends[4] = "Dodi"
  fmt.Println(friends[4])

  fmt.Println(friends)
  fmt.Println("-------------------------------------------------------")

  // Cannot redefine???
  // friends = [5]string{
  //   "Budi",
  // }
  // fmt.Println(friends)

  // Array must same data type

  values := [5]int32{
    32, -43, 46,
  }
  fmt.Println(values)

}
