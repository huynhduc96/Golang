package main

import (
	"fmt"
	"math"
)

// Các interface trong Go cung cấp một cách để xác định hành vi của một đối tượng: nếu đối tượng đó có thể làm những việc như thế này, thì nó có thể được sử dụng ở đây.
// Trong Go, bạn không thể "bắt buộc" một struct phải implement một interface theo nghĩa truyền thống như trong các ngôn ngữ lập trình hướng đối tượng khác. Go dựa vào cơ chế kiểm tra kiểu thời điểm biên dịch (compile-time type checking) và không có khái niệm kế thừa (inheritance) theo kiểu class.

// Ngôn ngữ Go hiện thực mô hình hướng đối tượng thông qua cơ chế Interface.

// Ví dụ có một interface con vịt, xác định khả năng Quacks:
type Duck interface {
	Quacks()
}

// một struct động vật bất kì
type Animal struct {
}

// con này có khả năng `Quacks` như vịt
func (a Animal) Quacks() {
	fmt.Println("The animal quacks")
}

// hàm dành cho vịt
func Scream(duck Duck) {
	duck.Quacks()
}

func main() {
	// a là một một vật thuộc struct Animal
	a := Animal{}

	// vì a có khẳng năng `Quacks` như vịt nên
	// ta có thể sử dụng nó như một con vịt trong hàm này
	Scream(a)

	exInterface2()
}

// Một ví dụ khác
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func exInterface2() {

	// 1 - Chúng ta định nghĩa interface Shape với hai phương thức Area() và Perimeter().
	// 2 - Rectangle và Circle thực hiện interface Shape bằng cách triển khai các phương thức tương ứng.
	// 3- Hàm PrintShapeInfo() nhận một đối tượng có kiểu Shape và sử dụng các phương thức Area() và Perimeter() của nó mà không cần biết kiểu cụ thể của đối tượng.
	rect := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 5}

	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
}
