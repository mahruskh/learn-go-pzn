package main

import (
    "fmt"
    "strings"
)


func sayHi(){
  fmt.Println("Hello...")
}

func spaceCode(){
  fmt.Println("-------------------------------------------------------------")
}

/*
  - Function param must define data type of parameter
  - Passing value data type must same as data type of parameter
  - Parameter must be filled.
  - Go has default value for param???
*/
func sayFriend(firstName string, age int32) {
  fmt.Println("Hello", firstName, age, "years age.")
}

func getLastName(name string) string {
  name = strings.TrimSpace(name)
  if name == "" {
    return "Not Found!!!"
  }

  words := strings.Split(name, " ")
  // fmt.Println(words)
  // fmt.Printf("Type = %T\n", words)
  // fmt.Println(len(words))

  // Cannot like in python ????
  // return words[-1]

  return words[len(words) - 1]
}

func getFirstLastName(name string) (string, string, int)  {
  name = strings.TrimSpace(name)
  if name == "" {
    return "None", "None", 0
  }

  words := strings.Split(name, " ")
  length := len(words)
  return words[0], words[length - 1], length
}

func main(){
  // Error never used
  // sayHi

  sayHi()
  for i := 0; i < 10; i++ {
    sayHi()
  }

  spaceCode()

  sayFriend("Mahrus", 25)
  sayFriend("Budi", 20)

  spaceCode()

  lastName := getLastName("Mahrus Khomaini")
  fmt.Println(lastName)

  fmt.Println(getLastName("Budi Gunawan Raharjo"))
  fmt.Println(getLastName("  "))
  fmt.Println(getLastName("Eko Budi    "))

  spaceCode()

  // fmt.Println(getFirstLastName("Mahrus Khomaini"))
  // firstName, lastName, length := getFirstLastName("Mahrus Khomaini")
  firstName, lastName, length := getFirstLastName("Budi Gunawan Raharjo")
  fmt.Println(firstName, lastName, length)






}
