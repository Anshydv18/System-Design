// package main

// import (
// 	"fmt"
// )

// type Reactangle struct {
// 	Width  int
// 	Height int
// }

// type Circle struct {
// 	Redius int
// }

// func CalculateArea(shape interface{}) interface{} {
// 	totalArea := 0.0

// 	switch s := shape.(type) {
// 	case Circle:
// 		return s
// 		// totalArea += 3.14 * float64(s.Redius) * float64(s.Redius)
// 	case Reactangle:
// 		totalArea += (float64(s.Height) * float64(s.Width))
// 	default:
// 		panic("unknown shape!")
// 	}

// 	return totalArea
// }

// func main() {

// 	var shape interface{} = Circle{Redius: 20}
// 	//circle := Circle{Redius: 10}

// 	fmt.Println(CalculateArea(shape))
// 	//fmt.Println(CalculateArea(shapes))
// }

package main

import "fmt"

// Shape interface
type Shape interface {
    Area() float64
}

// Rectangle struct
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Circle struct
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Create a slice of Shape interface
    shapes := []Shape{
        Rectangle{Width: 10, Height: 20},
        Circle{Radius: 10},
    }

    // Iterate and calculate areas
    for _, shape := range shapes {
        fmt.Printf("Area: %.2f\n", shape.Area())
    }
}
