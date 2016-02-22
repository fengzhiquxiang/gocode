package main 

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(a[:3])
	reverse(a[4:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}