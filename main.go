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
	slicesAndArrays()
	maps()
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
func slicesAndArrays() {
	animalsArr := [4]string{"dog", "cat", "elephant", "horse"}
	b := animalsArr[:2]
	c := animalsArr[1:]
	fmt.Println(b, c)
	c[0] = "monkey"
	d := animalsArr[:]
	fmt.Println(animalsArr, b, c, d)
}
func maps() {
	pointsMap := map[string]Point{}
	otherPointsMap := make(map[int]Point)
	fmt.Println(pointsMap, otherPointsMap)
	pointsMap["a"] = Point{X: 1, Y: 2}
	fmt.Println(pointsMap, pointsMap["a"])

	value, ok := pointsMap["a"]
	if ok {
		fmt.Println(value.X, value.Y)
	} else {
		fmt.Println("value doesn't exist")
	}
}
