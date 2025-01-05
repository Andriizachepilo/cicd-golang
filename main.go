package main

import "fmt"

func main() {
var nam string = "hello" 
name(nam)
}

func name(name string) (string) {
return surname(name,"mysurname")
}

func surname(surname,nam string) (string) {
cons := surname + nam
fmt.Print(cons)
return cons
}

// func usernameLength(string []string) {

// }