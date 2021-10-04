package main

import "fmt"

func main(){
  const (
      minScore float32 = 75
      minAttend uint8 = 60
  )
  var (
    finalScore float32 = 87.5
    attendance uint8 = 80
  )

  isScorePassed := finalScore >= minScore
  fmt.Println("Is score passed ?: ", isScorePassed)

  isAttendPassed := attendance >= minAttend
  fmt.Println("Is attendance passed?: ", isAttendPassed)


  isPassed := isScorePassed && isAttendPassed
  fmt.Println("Final Passed: ", isPassed)


  // Clean code
  fmt.Println(finalScore >= minScore && attendance >= minAttend)

}
