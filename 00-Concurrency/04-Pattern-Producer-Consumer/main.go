package main

import (
	"fmt"
	"time"
)

// mô hình Producer Consumer, giúp tăng tốc độ xử lý chung của chương trình bằng cách cân bằng sức mạnh của các thread "sản xuất" (produce) và "tiêu thụ" (consume).
// producer: liên tục tạo ra một chuỗi số nguyên dựa trên bội số factor và đưa vào channel
func Producer(factor int) (out chan int) {
	out = make(chan int) // Tạo kênh out
	go func() {          // Bắt đầu một goroutine
		for i := 0; i < 5; i++ { // Lặp 10 lần
			out <- factor * i // Gửi giá trị factor * i vào kênh out
		}
		close(out) // Đóng kênh out khi kết thúc
	}()
	return out // Trả về kênh out
}

// consumer: liên tục lấy các số từ channel ra để print
func Consumer(in1, in2 chan int) chan int {
	result := make(chan int)
	go func() {
		for {
			select {
			case <-in1:
				result <- <-in1
			case <-in2:
				result <- <-in2
			}
		}

	}()
	return result
}
func main() {
	// tạo một chuỗi số với bội số 3
	from3 := Producer(3)

	// tạo một chuỗi số với bội số 5
	from5 := Producer(5)
	// tạo consumer
	result := Consumer(from3, from5)

	// for msg := range result {
	// 	fmt.Println("golang:", <-msg)
	// }

	for i := 0; i < 5; i++ { // Lặp 10 lần
		fmt.Println("golang result:", <-result)
	}

	// for v := range result {
	// 	fmt.Println("golang:", v)
	// }

	// thoát ra sau khi chạy trong một khoảng thời gian nhất định
	time.Sleep(1 * time.Second)
}
