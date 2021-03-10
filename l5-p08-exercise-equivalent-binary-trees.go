/*
Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary
trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's
concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:
	type Tree struct {
	    Left  *Tree
	    Value int
	    Right *Tree
	}
Continue description on next page.

1. Implement the Walk function.

2. Test the Walk function.

	The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values
		k, 2k, 3k, ..., 10k.
	Create a new channel ch and kick off the walker:
		go Walk(tree.New(1), ch)
	Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here.
 */

package main

import (
	//"golang.org/x/tour/tree"
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func addNode(t *Tree, ss int) {
	if t.Value == 0 {
		t.Value = ss

	} else if (ss < t.Value) {
		if (t.Left == nil) { t.Left = &Tree{} }
		addNode(t.Left, ss)

	} else {
		if (t.Right == nil) { t.Right = &Tree{} }
		addNode(t.Right, ss)
	}
}

func (t *Tree) New(i int) {
	s := rand.Perm(10)
	for j := range s {
		addNode(t, i * (s[j] + 1))
	}
}

func (t *Tree) String() string {
	if t == nil { return "<nil>"}
	return fmt.Sprintf("(L: %v V: %v R: %v)", t.Left.String(), t.Value, t.Right.String())
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	if t.Value == 10 { close(ch) }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		j, ok := <- c2
		if (!ok || i != j) {
			return false
		}
	}
	_, notDone := <- c2
	return !notDone
}


func main() {
	rand.Seed(int64(time.Millisecond))
	t := &Tree{}
	t.New(1)
	fmt.Println(t)
	c := make(chan int)
	go Walk(t, c)
	for i := range c {
		fmt.Println(i)
	}

	s := &Tree{}
	s.New(1)
	fmt.Println(s)
	d := make(chan int)
	go Walk(s, d)
	for i := range d {
		fmt.Println(i)
	}

	fmt.Print("t & s: ")
	if (!Same(t,s)) { fmt.Print("Not ")}
	fmt.Println("The Same\n")

	tt := &Tree{nil, -1, t}
	fmt.Println(tt)
	cc := make(chan int)
	go Walk(tt, cc)
	for i := range cc {
		fmt.Println(i)
	}

	fmt.Print("tt & t: ")
	if (!Same(tt, t)) { fmt.Print("Not ")}
	fmt.Println("The Same\n")

	fmt.Println(s.Right)
	fmt.Print("s.Right & s: ")
	if (!Same(s.Right, s)) { fmt.Print("Not ")}
	fmt.Println("The Same\n")
}
