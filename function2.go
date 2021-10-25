package main

import "fmt"

// Function return named value
// func getFullName()(firstName string, middleName string, lastName string){
func getFullName()(firstName, middleName, lastName string){
  firstName = "Mahrus"
  middleName = "Hola"
  lastName = "Khomaini"

  return
  // No need to return like below, just `return`.
  // return firstName, middleName, lastName
}

func spaceCode(){
  fmt.Println("-------------------------------------------------------------")
}


// Variadic function
// Can only use thiw way only for last argument
// In python we called it an args.
func sumAll(numbers ...int) int {
  fmt.Printf("Type = %T\n", numbers)
  // numbers is a slice, not an array !!!

  total := 0
  for _, number := range numbers {
    total += number
  }
  return total
}


func main(){
  first, middle, last := getFullName()
  fmt.Println(first, middle, last)

  spaceCode()

  // fmt.Println(sumAll())
  fmt.Println(sumAll(12, 4, 6 , 5 , 10 , 20))

  spaceCode()

  // Slice parameters
  // In python we user `*args` or like passing *list in function
  numbers := []int{10, 22, 12, 40, 12, 10}
  total := sumAll(numbers...)
  fmt.Println(total)

}
