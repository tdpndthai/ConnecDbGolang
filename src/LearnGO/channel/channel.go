package main

import(
	"fmt"
	//"time"
) 

// func hello(done chan bool) {  
//     fmt.Println("hello go routine is going to sleep")
//     time.Sleep(4 * time.Second)
//     fmt.Println("hello go routine awake and going to write to done")
//     done <- true //ghi dữ liệu vào done
// }

// Hàm sendData chuyển đổi channel này thành kênh chỉ gửi trong tham số sendch chan <- int. 
//Vì vậy, bây giờ channel chỉ được gửi bên trong sendData Goroutine nhưng nó là hai chiều trong Goroutine chính. 

// func sendData(sendch chan<- int) {  
//     sendch <- 10
// }

// func calcSquares(number int, squareop chan int) {  
//     sum := 0
//     for number != 0 {
//         digit := number % 10
//         sum += digit * digit
//         number /= 10
//     }
//     squareop <- sum
// }

// func calcCubes(number int, cubeop chan int) {  
//     sum := 0 
//     for number != 0 {
//         digit := number % 10
//         sum += digit * digit * digit
//         number /= 10
//     }
//     cubeop <- sum
// } 


func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}

func main() {
	//var a chan int //khởi tạo channel kiểu int,chưa có giá trị -> channel nil không được sử dụng và phải được xác định bằng cách
	// sử dụng make tương tự như map và slice
	//a := make(chan int)
	// if a == nil { //mới khởi tạo chưa có giá trị
	// 	fmt.Println("channel a is nil, going to define it")
	// 	a = make(chan int)
	// 	fmt.Printf("Type of a is %T", a)
	// }
	//Gửi và nhận từ channel
	//data := <-a // đọc từ channel a và lưu giá trị vào data
	//a <- data   // ghi dữ liệu từ data vào a
	// 	Hướng của mũi tên đối với channel chỉ định liệu dữ liệu được gửi hay nhận.
	// Trong dòng đầu tiên, mũi tên chỉ ra từ a và do đó chúng ta đang đọc từ channel a và lưu trữ giá trị vào biến data.
	// Trong dòng thứ hai, mũi tên chỉ về phía a và do đó chúng ta đang ghi vào channel a.

	//Khi dữ liệu được gửi đến một channel, điều khiển sẽ bị chặn trong câu lệnh gửi cho đến khi 
	//một số Goroutine khác đọc từ channel đó. 
	//Tương tự khi dữ liệu được đọc từ một channel, việc đọc bị chặn cho đến khi một số Goroutine ghi dữ liệu vào channel đó.
    // done := make(chan bool) //tạo mới 1 channel
    // fmt.Println("Main going to call hello go goroutine")
    // go hello(done)
    // <-done //đọc dữ liệu từ channel
	// fmt.Println("Main received data")
	
	// number := 589
    // sqrch := make(chan int)
    // cubech := make(chan int)
    // go calcSquares(number, sqrch)
    // go calcCubes(number, cubech)
    // squares, cubes := <-sqrch, <-cubech
    // fmt.Println("Final output", squares + cubes)

    // chnl := make(chan int) //1 channel chnl 2 chiều được tạo
    // go sendData(chnl) //truyền dưới dạng tham số cho sendData Goroutine
    // fmt.Println(<-chnl) //Chương trình này sẽ in 10 là đầu ra.

    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
