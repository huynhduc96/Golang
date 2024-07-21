package main

import "fmt"

// hàm được đặt tên
func Add(a, b int) int {
	return a + b
}

// // hàm ẩn danh
// var Add = func(a, b int) int {
// 	return a + b
// }

// Cả tham số truyền vào và các giá trị trả về đều có thể được đặt tên:
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	// return về value và ok đã được đặt tên
	return
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	var a = []interface{}{123, "abc"}

	// tương đương với lời gọi trực tiếp `Print(123, "abc")`
	Print(a...) // 123 abc

	// tương đương với lời gọi `Print([]interface{}{123, "abc"})`
	Print(a) // [123 abc]

}
