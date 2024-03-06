package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) method() {
	fmt.Println(p.X, p.Y)
}
func (p Point) movePointMethod(x, y int) Point {
	p.X += x
	p.Y += y
	return p
}
func (p *Point) movePointMethodPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	pointers()
	structs()
	slicesAndArrays()
	maps()

	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCallback, "sum")
	fmt.Println(result)

	multiplyCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	result = doSomething(multiplyCallback, "multiply")
	fmt.Println(result)

	orderPrice := totalPrice(1) //sum = 1
	fmt.Println(orderPrice(1))  // sum = 1+1 closure
	fmt.Println(orderPrice(2))  // sum = 1+1+2 closure

	p := Point{X: 1, Y: 2}
	fmt.Println(p)
	fmt.Println(movePoint(p, 1, 1))
	fmt.Println(p)
	movePointPtr(&p, 1, 1)
	fmt.Println(p)
	p.movePointMethod(1, 5)
	fmt.Println(p)
	p.movePointMethodPtr(1, 5)
	fmt.Println(p)
	v := &p
	v.movePointMethodPtr(1, 1)
	fmt.Println(p)

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
func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Print(s, " ")
	return result
} // doSomething is callback function
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
} // totalPrice is a closure
func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}
func movePointPtr(p *Point, x, y int) {
	p.X += x
	p.Y += y
}
