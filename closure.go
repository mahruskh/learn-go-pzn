package main

import "fmt"

func spaceCode() {
	fmt.Println("-------------------------------------------------------------")
}

func main() {
	name := "Mahrus"
	age := 25

	fmt.Println(name, age)

	spaceCode()


  // Is function cannot have parameter default value????
	nextYear := func() {
		// This will change `age` value above.
		// age = 19

		// This will create a new variable `age` and it different with `age` above outide this func.
		// age := 20

		fmt.Println("Current age:", age)
		age++
		fmt.Println("Next age:", age)

		// Testing new variable
		isActivate := false
		fmt.Println(isActivate)
	}
	nextYear()

	spaceCode()

	fmt.Println(name, age)

	// Testign variable from `nextYear` func, and error.
	// Variable not found.
	// fmt.Println(isActivate)

}
