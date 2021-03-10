/*

Interfaces
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. VertexF (the value type) doesn't implement Abser because the Abs
method is defined only on *VertexF (the pointer type).
 */

package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloaty(-math.Sqrt2)
    v := VertexF{3, 4}

    a = f  // a MyFloaty implements Abser
    a = &v // a *VertexF implements Abser

    // In the following line, v is a VertexF (not *VertexF)
    // and does NOT implement Abser.
    // a = v

    fmt.Println(a.Abs())
}

type MyFloaty float64

func (f MyFloaty) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type VertexF struct {
    X, Y float64
}

func (v *VertexF) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}