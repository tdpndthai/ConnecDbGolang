package main

import (
	"fmt"
	"time"
)
func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {

	// 	Khi một Goroutine mới khởi chạy, tính năng goroutine sẽ gọi trở lại ngay lập tức.
	//Không giống hàm, hệ thống không chờ Goroutine chạy xong, mà trả lại ngay lập tức dòng code tiếp theo ngay sau khi Goroutine gọi và bất kỳ giá trị trả lại nào từ Goroutine sẽ bị bỏ qua.
	// Goroutine chính cần chạy để các Goroutine khác chạy được. 
	//Nếu Goroutine chính dừng lại thì chương trình cũng dừng và không Goroutine nào chạy nữa.
	//go hello() //bắt đầu một Goroutine mới
	//Phương pháp sử dụng sleep trong Goroutine chính để đợi những Goroutine khác kết thúc việc thực thi là một mẹo nhỏ (hack) chúng ta sử dụng để hiểu cách làm việc của các Goroutines
	// time.Sleep(1 * time.Second) //ru ngủ goroutine khi được khởi chạy -> ru ngủ trong 1s bây giờ hàm hello có đủ thời gian để thực thi trước khi goroutine chính dừng lại
	// fmt.Println("main function")
	// Sau khi gọi go hello() tại dòng 11, hệ thống trả về ngay lập tức dòng tiếp theo của code mà không đợi hello goroutine kết thúc và in ra main function. 
	// Sau đó Goroutine chính dừng khi không còn đoạn code nào để thực thi và do đó, hello Goroutine không còn cơ hội để chạy.

	go numbers()
    go alphabets()
    time.Sleep(1100 * time.Millisecond)
    fmt.Println("main terminated")
}
