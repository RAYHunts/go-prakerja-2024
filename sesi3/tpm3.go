package main

import (
	"fmt"
	"os"
)

type Person struct {
    Name    string
    Address string
    Jobs    string
    Reason  string
}

func askForDetails() Person {
    if len(os.Args) != 5 {
        fmt.Println("Usage: go run main.go [Name] [Address] [Job] [Reason]")
        os.Exit(1)
    }

    person := Person{
        Name:    os.Args[1],
        Address: os.Args[2],
        Jobs:    os.Args[3],
        Reason:  os.Args[4],
    }

    return person
}

func displayDetails(person Person) {
    fmt.Println("Nama: ", person.Name)
    fmt.Println("Alamat: ", person.Address)
    fmt.Println("Pekerjaan: ", person.Jobs)
    fmt.Println("Alasan: ", person.Reason)
}

func main() {
    person := askForDetails()
    displayDetails(person)
}