package main

import "fmt"

func main(){
  for i := 0; i < 10; i++ {
    if i == 9 {
      fmt.Println("Stop using `break` at loop", i)
      break
    }

    if i == 3 || i == 7 {
      fmt.Println("Skip using `continue` at loop", i)
    }

    fmt.Println("Loop ", i)
  }

}
