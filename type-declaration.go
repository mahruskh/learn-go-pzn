package main

import "fmt"

func main(){
  type (
      NoIdentity string
      hasAccess bool
  )

  var noKTP NoIdentity = "2141732186328168"
  var readAccess hasAccess = true

  fmt.Println("No. KTP: ", noKTP)
  fmt.Println("Read access:", readAccess)
}
