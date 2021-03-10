/*

Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
    var v VertexC
    ScaleFunc(v, 5)  // Compile error!
    ScaleFunc(&v, 5) // OK
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
    var v VertexC
    v.Scale(5)  // OK
    p := &v
    p.Scale(10) // OK

For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is
called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the
Scale method has a pointer receiver.
 */

package main

import "fmt"

type VertexC struct {
	X, Y float64
}

func (v *VertexC) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *VertexC, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexC{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &VertexC{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
