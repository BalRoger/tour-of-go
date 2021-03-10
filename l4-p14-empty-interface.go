/*

The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of
arguments of type interface{}.
 */

package main

import "fmt"

func main() {
	var i interface{}
	describee(i)

	i = 42
	describee(i)

	i = "hello"
	describee(i)
}

func describee(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
