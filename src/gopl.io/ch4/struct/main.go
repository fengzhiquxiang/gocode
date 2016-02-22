// Charcount computes counts of Unicode characters.
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {

	var dilbert Employee
    var thirdvar *Employee = &dilbert
    thirdvar.Position = "china"
    fmt.Printf("%v\n", thirdvar.Position)
    (*thirdvar).Position = "cccc"
    fmt.Printf("%v\n", (*thirdvar).Position)
}
