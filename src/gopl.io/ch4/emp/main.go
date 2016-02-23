// Charcount computes counts of Unicode characters.
package main

import (
	"fmt"
	// "time"

    "gopl.io/ch4/emp/haha"
)

var emp = make(map[int]*haha.Employee)

func main() {
	var dilbert haha.Employee
    var thirdvar *haha.Employee = &dilbert
    thirdvar.Position = "china"
    thirdvar.ID = 999
    thirdvar.Name = "zzz"
    fmt.Printf("%v\n", thirdvar.Position)
    (*thirdvar).Position = "cccc"
    fmt.Printf("%v\n", (*thirdvar).Position)

    emp[1] = thirdvar
    fmt.Println(EmployeeByID(1))
}

func EmployeeByID(id int) *haha.Employee {
    return emp[id]
}