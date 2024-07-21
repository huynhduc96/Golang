package main

// Tác vụ atomic trên một vùng nhớ chia sẻ thì đảm bảo rằng vùng nhớ đó chỉ có thể được truy cập bởi một Goroutine tại một thời điểm. Để đạt được điều này ta có thể dùng sync.Mutex.
import (
	// package cần dùng
	"fmt"
	"sync"
	"sync/atomic"
)

// total là một atomic struct
var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	// thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		// chặn các Goroutines khác vào
		total.Lock()
		// bây giờ, lệnh total.value += i được đảm bảo là atomic (đơn nguyên)
		total.value += i
		// bỏ chặn
		total.Unlock()
	}
}

func main_ex1() {
	// khai báo wg để main Goroutine dừng chờ các Goroutines khác trước khi kết thúc chương trình
	var wg sync.WaitGroup
	// wg cần chờ 2 Goroutines khác
	wg.Add(2)
	// thực thi Goroutines thứ nhất
	go worker(&wg)
	// thực thi Goroutines thứ hai
	go worker(&wg)
	// wg bắt đầu đợi để 2 Goroutines kia xong
	wg.Wait()
	// in ra kết quả thực thi
	fmt.Println(total.value)
}

/**
* ? Thay vì dùng mutex, chúng ta cũng có thể dùng package sync/atomic,
* ? đây là giải pháp hiệu quả hơn đối với một biến số học.
 */

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	// wg thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		// lệnh cộng atomic.AddUint64 total được đảm bảo là atomic (đơn nguyên)
		atomic.AddUint64(&total2, i)
	}
}

func main() {
	// wg được dùng để dừng hàm main đợi các Goroutines khác
	var wg sync.WaitGroup
	// wg cần đợi hai Goroutines gọi lệnh Done() mới thực thi tiếp
	wg.Add(2)
	// tạo Goroutines thứ nhất
	go worker2(&wg)
	// tạo Goroutines thứ hai
	go worker2(&wg)
	// bắt đầu việc đợi
	wg.Wait()
	// in ra kết quả
	fmt.Println(total2)
}
