/*

Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
 */

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		counts[w] = counts[w] + 1
	}
	return counts
}

func main() {
	// wc.Test(WordCount)
	fmt.Println(WordCount("Help! I need some body. Help! Not just any body. Help! You know I need some one. Help!"))
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println(WordCount("A man a plan a canal panama."))
}
