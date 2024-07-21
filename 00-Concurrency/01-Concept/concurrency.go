package main

import (
	"fmt"
	"time"
)

// Concurrency is about dealing with lots of things at once-Rob Pike
// Parallelism is about doing lots of things at once-Rob Pike

//  Số lượng thread chạy song song trong cùng một thời điểm sẽ bằng với số lượng nhân CPU mà máy tính chúng ta có. Vì vậy khi chúng ta lập trình mà tạo quá nhiều thread thì cũng không có giúp cho chương trình chúng ta chạy nhanh hơn, mà còn gây ra lỗi và làm chậm chương trình. Theo kinh nghiệm lập trình chúng ta chỉ nên tạo số thread bằng số nhân CPU * 2.

//  Goroutines là một đơn vị concurrency của ngôn ngữ Go. Hay nói cách khác Golang sử dụng goroutine để xử lý đồng thời nhiều tác vụ.
//  Việc khởi tạo goroutines sẽ ít tốn chi phí hơn khởi tạo thread so với các ngôn ngữ khác.
//  Cách khởi tạo goroutine chỉ đơn giản thông qua từ khóa go.
//  Về góc nhìn hiện thực, goroutines và thread không giống nhau.

// PA1: Sử dụng time sleep để ước lượng thời gian , từ đó có thể chạy 2 goroutines một lúc
func main() {
	// sử dụng từ khoá go để tạo goroutine
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// chờ 1 giây để có thể chạy được goroutine
	//của hàm fmt.Println trước khi hàm main kết thúc
	time.Sleep(time.Second)
}
