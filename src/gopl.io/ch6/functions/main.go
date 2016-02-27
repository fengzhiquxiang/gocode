package main 

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point)print() {
	fmt.Printf("%d, %d\n", p.X, p.Y)
}

func main() {
	
	p := Point{1, 2}
	fn := p.print
	fn()
}