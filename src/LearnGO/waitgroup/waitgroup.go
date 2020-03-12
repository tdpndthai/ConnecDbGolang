package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //Bộ đếm được giảm xuống theo lệnh gọi wg.Done trong processGoroutine
}

func main() {
	//Waitgroup được sử dụng để chờ bộ sưu tập Goroutines hoàn tất.
	//Kiểm soát bị chặn cho đến khi tất cả các Goroutines thực hiện xong
	no := 3
	var wg sync.WaitGroup //khởi tạo 1 waitgroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg) //sinh ra i số process
	}
	wg.Wait() //đợi cho đến khi bộ đếm trở thành không
	fmt.Println("All go routines finished executing")
}
