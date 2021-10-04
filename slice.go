package main

import "fmt"

func main(){
  var months = [12]string{
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  }
  fmt.Println(months, len(months))

  fmt.Println("-------------------------------------------------------")

  var friends = [...]string{
      "Budi", "Eko", "Andre", "Doni", "Hamzah", "Alex",
      "Pandji", "Babe", "Arif",
  }
  fmt.Println(friends, len(friends))

  var friendsTeam1 = friends[3:8]
  fmt.Println(friendsTeam1)
  fmt.Println(friendsTeam1[0])

  friends[4] = "Zulham"
  fmt.Println(friends, len(friends), cap(friends))
  fmt.Println(friendsTeam1, len(friendsTeam1), cap(friendsTeam1))

  fmt.Println("-------------------------------------------------------")

  friendsTeam1[3] = "Mamat"
  fmt.Println(friendsTeam1)
  fmt.Println(friends)

  fmt.Println("-------------------------------------------------------")

  var allFriendsTeam = friends[:]
  fmt.Println(allFriendsTeam)

  fmt.Println("-------------------------------------------------------")
}
