package main

import "fmt"

func main(){
  var friends = [...]string{
      "Budi", "Eko", "Andre", "Arif",
  }
  fmt.Println(friends)

  var slice1 = friends[2:3]
  fmt.Println(slice1)

  fmt.Println("-------------------------------------------------------")

  var slice2 = append(slice1, "MAHRUS")
  fmt.Println(slice1)
  fmt.Println(slice2)
  fmt.Println(friends)


  fmt.Println("-------------------------------------------------------")

  newSlice := make([]string, 2, 5)
  newSlice[0] = "Mahrus"
  newSlice[1] = "Khomaini"

  fmt.Println(newSlice, len(newSlice), cap(newSlice))

  fmt.Println("-------------------------------------------------------")

  copiedSlice := make([]string, len(newSlice), cap(newSlice))

  copy(copiedSlice, newSlice)
  fmt.Println(copiedSlice, len(copiedSlice), cap(copiedSlice))

  copiedSlice[1] = "HORE"
  fmt.Println(copiedSlice, len(copiedSlice), cap(copiedSlice))

  fmt.Println("-------------------------------------------------------")

  anArray := [5]int{1, 2, 3, 4 ,5}
  anArray2 := [...]int{1, 2, 3}
  aSlice := []int{1, 2, 3, 4}

  fmt.Println(anArray)
  fmt.Println(anArray2)
  fmt.Println(aSlice)

}
