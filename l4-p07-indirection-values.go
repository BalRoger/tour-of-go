/*

Methods and pointer indirection (2)
The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:
    var v VertexD
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:
	var v VertexD
	fmt.Println(v.Abs()) // OK
	p := &v
	fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
 */

package main

import (
	"fmt"
	"math"
)

type VertexD struct {
	X, Y float64
}

func (v VertexD) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v VertexD) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := VertexD{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertexD{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
