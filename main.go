package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	a := "hello"
	b := 10
	fmt.Println(a, b)
	c := &a
	fmt.Println(a, c)
	*c = "world"
	fmt.Println(a, *c)
	d := &b
	*d = *d / 2
	fmt.Println(b)
}
