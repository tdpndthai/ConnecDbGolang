package main

import (
	"fmt"
)

func main() {
	// for i := 1; i <= 10; i++ {
	//     if i%2 == 0 {
	//         continue //chia hết cho 2 thì chuyển sang lần lặp tiếp theo của vòng lặp
	//     }
	//     fmt.Printf("%d ", i)
	// }

	//biến thể
	// i := 0
	// for ;i <= 10; { // biến khởi tạo và biến chạy có thể để trống
	//     fmt.Printf("%d ", i)
	//     i += 2
	// }

	//có thể bỏ dấu chấm phẩy -> giống vòng lặp while
	// i := 0
	// for i <= 10 {
	//     fmt.Printf("%d ", i)
	//     i += 2
	// }

	// 	for i := 1; i <= 10; i++ {
	// 		if i > 5 {
	// 			break //vòng lặp dừng nếu i > 5
	// 		}
	// 		fmt.Printf("%d ", i)
	// 	}
	// 	fmt.Printf("\n Câu lệnh ngay sau vòng lặp")

	for no, i := 10, 1; i <= 10 && no <= 19; no,i = no+1,i+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

}
