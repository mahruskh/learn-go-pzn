package main

import "fmt"

func main(){
  value1 := 10.5
  value2 := 7

  fmt.Printf("Type = %T\n", value1)
  fmt.Printf("Type = %T\n", value2)

  fmt.Println(value1 > float64(value2))
  fmt.Println("halo" > 2)


  // Question?
  // Can only compare value with the same specific type?
  // - string cannot compare with number. (v)
  // - number with difference type cannot compare (?)
}
