package main

import "fmt"

func main() {
  
 
 
  var input string
  var hours int
  var wage int

  fmt.Println("Welcome to the Paystub Calculator!")

  fmt.Println("Start by entering your hourly wage.")

  fmt.Scanln(&wage)

  var OvertimeMultiplier int = wage * 150 / 100
  fmt.Println(OvertimeMultiplier)

  fmt.Println()

  fmt.Println("Now, enter the amount of hours you worked this pay period.")

  fmt.Scanln(&hours)

  fmt.Println()

  fmt.Println("Did you work overtime this pay period?")
  fmt.Println("input 'yes' or 'no'")

  var overTimeworked int

  fmt.Scanln(&input)

  if input == "yes" {
    fmt.Println("How many hours of overtime did you work?")
    fmt.Scanln(&overTimeworked)
  } else {
    Taxes := wage * hours * 12 / 100
  fmt.Println(Taxes)
    fmt.Println()
  }

fmt.Println(overTimeworked)

 Overtime := overTimeworked * OvertimeMultiplier
 fmt.Println(Overtime)

Taxes := wage * hours + Overtime * 12 / 100
  fmt.Println(Taxes)

  fmt.Print(wage * hours + Overtime * 12 / 100 - Taxes)



}