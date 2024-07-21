package main

import (
	"fmt"
	"os"
)

// Lệnh defer trì hoãn việc thực thi hàm cho tới khi hàm bao ngoài nó return. Các đối số trong lời gọi defer được đánh giá ngay lặp tức nhưng lời gọi không được thực thi cho tới khi hàm bao ngoài nó return.
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	ex3()
}

// Mỗi lời gọi defer được push vào stack và thực thi theo thứ tự ngược lại khi hàm bao ngoài nó kết thúc.
//Ta thường sử dụng defer cho việc đóng hoặc giải phóng tài nguyên:

// Ex:
// 01- Đóng file giống như try-finally:
func CloseFileEx() {
	f, err := os.Create("file")
	if err != nil {
		panic("cannot create file")
	}

	// chắc chắn file sẽ được close dù hàm có bị panic hay return
	defer f.Close()
	fmt.Fprintf(f, "hello")
}

// 02- Đóng file và xử lý panic giống như try-catch-finally:
func CloseFileEx2() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	// . là folder hiện tại
	f, err := os.Create(".")
	if err != nil {
		panic("cannot create file")
	}
	defer f.Close()

	// không quan trọng chuyện gì xảy ra thì file cũng sẽ được close
	// để đơn giản nên ở đây bỏ qua bước kiểm ra close result
	fmt.Fprintf(f, "hello")
}

// 03 Cũng giống như block finally thì lời gọi defer cũng có thể làm cho kết quả trả về thay đổi:

func yes() (text string) {
	defer func() {
		text = "no"
	}()
	return "yes"
}

// Result : no
func ex3() {
	fmt.Println(yes())
}
