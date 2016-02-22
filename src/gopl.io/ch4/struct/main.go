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

var emp = make(map[int]*Employee)

func main() {
	var dilbert Employee
    var thirdvar *Employee = &dilbert
    thirdvar.Position = "china"
    thirdvar.ID = 999
    thirdvar.Name = "zzz"
    fmt.Printf("%v\n", thirdvar.Position)
    (*thirdvar).Position = "cccc"
    fmt.Printf("%v\n", (*thirdvar).Position)

    emp[1] = thirdvar
    fmt.Println(EmployeeByID(1))
}

func EmployeeByID(id int) *Employee {
    return emp[id]
}