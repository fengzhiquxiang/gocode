package main 

import (
	"fmt"
	"sort"
	"os"
	"bufio"
)

func main() {
	ages := make(map[string]int)
	ages["me"] = 40
	ages["yu"] = 18
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
	ages["me"]++
	var names []string
	for name := range ages {
	    names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
	    fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println(ages == nil)   
	if age, ok := ages["bob"]; !ok {
		fmt.Printf("%d is not exist\n", age)
	} 

	seen := make(map[string]bool) // a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }

	// fmt.Println(ages["me"])
	// ages["me"] = 40111
	// fmt.Println(ages["me"])
	// delete(ages,"me")
	// fmt.Println(ages["me"])
	// _ = &ages["me"] // compile error: cannot take address of map element
}