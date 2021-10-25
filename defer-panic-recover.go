package main

import "fmt"

func endApp() {
	fmt.Println("Done called function.")

  // When panic then program will throwback exvaption to up
  // And recover must call in end off program, like in endApp defer func.
  errorMessage := recover()
  if errorMessage != nil {
    fmt.Println("Error with Messsage:", errorMessage)
  }

  fmt.Println("----------------------------------------------------------")
}

func runApp(value int) {
  defer endApp()
  fmt.Println("Start running app.")

  // Testing has error
  // result := 100 / value
  // fmt.Println(result)

  // Testing with `panic`
  // Just for example, need another case.
  isError := false
  if value == 0 {
    isError = true
  }

  // TESTING, still cannot catch panic
  // Must in end defer func, try another case!!!
  // errorMessage := recover()
  // fmt.Println("Error with Messsage:", errorMessage)

  if isError {
    panic("ERROR RUNNING APP, DIVISION BY ZERO")
    // Error raise but defer `endApp` still execute.
  }
  result := 100 / value
  fmt.Println(result)

  fmt.Println("App is running...")

  // Example basic call func
  // logging()


}

func main() {
  // ============== RECOVER
  // Error and logging func not called
	// runApp(0)

  // Error but logging func still called
  // runApp(0)

  // If above error then below not excute.
  runApp(10)

  // =============== PANIC
  // I think panic is Exception in python.
  // panic, will stop app.
  runApp(7)


  // =================== RECOVER
  // Catch the panic, and app still running
  runApp(0)

}
