package main

import "fmt"

func main(){
  result := 10 + 5
  fmt.Println("result:", result)

  fmt.Println("---------------------------------------------------")


  value1 := 8
  value2 := 10.5
  result2 := float32(value1) + float32(value2)
  fmt.Println("result2:", result2)


  fmt.Println("---------------------------------------------------")
  total := 10
  fmt.Println("total", total)

  total += 2
  fmt.Println("total +2:", total)

  total -= 1
  fmt.Println("total -1:", total)

  total /= 3
  fmt.Println("total / 3: ", total)

  total *= 5
  fmt.Println("total * 3:", total)

  total %= 4
  fmt.Println("total % 4:", total)

  total++
  fmt.Println("total ++:", total)

  total--
  fmt.Println("total --:", total)


}
