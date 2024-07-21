package main

import (
	"fmt"
	"sync"
)

// Trong hầu hết các ngôn ngữ hiện đại, vấn đề chia sẻ tài nguyên được giải quyết bằng cơ chế đồng bộ hóa như khóa (lock) nhưng Golang có cách tiếp cận riêng là chia sẻ giá trị thông qua channel.

//  Golang có một triết lý được thể hiện bằng slogan:
// Do not communicate by sharing memory; instead, share memory by communicating.
// Do not communicate through shared memory, but share memory through communication.

// Mặc dù các vấn đề tương tranh đơn giản như tham chiếu đến biến đếm có thể được hiện thực bằng atomic operations hoặc mutex lock, nhưng việc kiểm soát truy cập thông qua Channel giúp cho code của chúng ta clean và "Golang" hơn.

/**
 * ? Cách 1 : Sử dụng mutex
 */

func mutex() {
	var mu sync.Mutex

	// Unlock mutex ngay từ đầu
	mu.Unlock()

	go func() {
		mu.Lock()
		fmt.Println("Hello World")
		mu.Unlock()
	}()

	mu.Lock()
	fmt.Println("Will not run this code 1")
	mu.Unlock()
	fmt.Println("Will not run this code")
}

/**
 * ? Cách 2 : Sử dụng Channel
 */

// Đồng bộ hóa với mutex là một cách tiếp cận ở mức độ tương đối đơn giản. Bây giờ ta sẽ sử dụng một unbuffered channel để hiện thực đồng bộ hóa:
func channel_simple() {
	done := make(chan int, 1)

	go func() {
		fmt.Println("Hello World")

		// gửi 1 giá trị vào channel thông báo kết thúc goroutine này
		done <- 1
	}()

	// main thread nhận giá trị từ channel và thoát khỏi
	// trạng thái block
	<-done

}

// Dựa trên buffered channel, chúng ta có thể dễ dàng mở rộng thread print đến N.
// Ví dụ sau là mở 10 goroutine để in riêng biệt: (tuy nhiên, ko đảm bảo thứ tự in)
func channel_simple_2() {
	done := make(chan int, 10)

	// mở ra N goroutine
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("Hello World", i)
			done <- i
		}()
	}

	// đợi cả 10 goroutine hoàn thành
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

/**
 * ? Cách 3 : Sử dụng sync.WaitGroup thay cho Channel
 */

func channel_wait_group() {
	var wg sync.WaitGroup

	// mở N goroutine
	for i := 0; i < 10; i++ {
		// tăng số lượng sự kiện chờ, hàm này phải được
		// đảm bảo thực thi trước khi bắt đầu 1 goroutine chạy nền
		wg.Add(1)

		go func() {
			fmt.Println("Hello World")

			// cho biết hoàn thành một sự kiện
			wg.Done()
		}()
	}

	// đợi N goroutine hoàn thành
	wg.Wait()
}
