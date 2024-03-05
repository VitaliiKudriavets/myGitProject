package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) method() {
	fmt.Println(p.X, p.Y)
}

func main() {
	pointers()
	structs()
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
func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(p1)
	fmt.Println(p1.X, p1.Y)
	g := &p1
	fmt.Println(g.Y)
	g.method()
}
