package main

import "fmt"

func main(){
  const (
    phi = 3.24
    simplePhi = 22 / 7
    simplePhiStr = "22 / 7"
    pctValue = 100
  )

  // Error, cannot change.
  // phi = 22 / 7

  fmt.Println("Phi:", phi)
  fmt.Println("Simple Phi:", simplePhi)
  fmt.Println("Simple Phi String:", simplePhiStr)

}
