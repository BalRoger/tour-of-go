/*
Nil interface values
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate
which concrete method to call.
 */

package main

import "fmt"

type IIII interface {
	M()
}

func main() {
	var i IIII
	describeMe(i)
	i.M()
}

func describeMe(i IIII) {
	fmt.Printf("(%v, %T)\n", i, i)
}
