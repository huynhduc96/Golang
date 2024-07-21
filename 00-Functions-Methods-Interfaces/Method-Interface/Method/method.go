package main

import (
	"fmt"
	"image/color"
	"math"
)

// Go không có class, tuy nhiên chúng ta có thể định nghĩa các phương thức (Method) cho type (kiểu).
// Method là một hàm với đối số (argument) đặc biệt gọi là receiver.

type Vertex struct {
	X, Y float64
}

// method Abs() với receiver 'v'
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// kết quả:
	// 5

	// ex1
	ex1()
}

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	// thuộc tính ẩn danh
	Point

	// thuộc tính bình thường
	Color color.RGBA
}

// Chúng ta có thể định nghĩa ColoredPoint như một struct có 3 trường, nhưng ở đây chúng ta sẽ dùng struct Point chứa X và Y để thay thế.

func ex1() {
	// khai báo một đối tượng thuộc struct
	var cp ColoredPoint

	// có thể gán thẳng vào thuộc tính X
	// không cần phải thông qua Point
	cp.X = 1

	// có thể truy cập X bằng cách này
	fmt.Println(cp.Point.X)
	// "1"

	// hoặc gán vào Y thông qua Point
	cp.Point.Y = 2

	// và truy cập Y bằng cách này
	fmt.Println(cp.Y)
	// "2"
}
