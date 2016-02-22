package main 

import (
	"fmt"
	"crypto/sha256"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB  
)

func main() {
	var a [3]int = [3]int{2, 3, 4}
	var b [5]int = [...]int{2, 3, 4, 5, 6}
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	// fmt.Println(a[0])
	zero(&a)
	for i,v := range a{
		fmt.Printf("a[%d] => %d\n", i, v)
	}

	fmt.Printf("------------------------\n")

	zero1(&b)
	for _, v := range b {
		fmt.Printf("%d (%T)\n", v, v)
	}

	fmt.Println(RMB, symbol[RMB])

	fmt.Printf("------------------------\n")

	c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%X\n%X\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func zero(ptr *[3]int) {
    for i := range ptr {
        ptr[i] = 0
    }
}

func zero1(ptr *[5]int) {
    *ptr = [5]int{}
}