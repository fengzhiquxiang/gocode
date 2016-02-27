// 
package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	// w.X = 8            // equivalent to w.Circle.Point.X = 8
	// w.Y = 8            // equivalent to w.Circle.Point.Y = 8
	// w.Radius = 5       // equivalent to w.Circle.Radius = 5
	// w.Spokes = 20

	// w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
	    Circle: Circle{
	        Point:  Point{X: 81, Y: 18},
	        Radius: 35,
	    },
	    Spokes: 220, // NOTE: trailing comma necessary here (and at Radius)
	}

	// fmt.Println(w.X)
	// fmt.Println(w.Radius)

	fmt.Printf("%#v\n", w)//Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。
	//
}