/*

Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
 */

package main

import (
	//"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = byte('A')
	return 1, nil
}

func main() {
	//reader.Validate(MyReader{})
	reader := MyReader{}
	ba := []byte{0}

	n, err := reader.Read(ba)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, ba)
	fmt.Printf("b[:n] = %q\n\n", ba[:n])

	n, err = reader.Read(ba)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, ba)
	fmt.Printf("b[:n] = %q\n\n", ba[:n])

	n, err = reader.Read(ba)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, ba)
	fmt.Printf("b[:n] = %q\n\n", ba[:n])

	n, err = reader.Read(ba)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, ba)
	fmt.Printf("b[:n] = %q\n\n", ba[:n])
}
