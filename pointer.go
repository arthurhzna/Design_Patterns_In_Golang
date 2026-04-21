package main

import "fmt"

func main() {
	var test int
	test = 10

	p := &test

	fmt.Println(test) // 10
	fmt.Println(p)    // address memory
	fmt.Println(*p)   // 10
}

// test --> value
// &test --> address memory
// p --> address memory, pointer to test
// *p --> value, value of test
