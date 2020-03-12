package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	// finger := 4
	// switch finger {
	// case 1:
	// 	fmt.Println("ngón cái")
	// default:
	// 	fmt.Println("không hợp lệ")
	// }

	//nhiều điều kiện trong 1 case
	// letter := "i"

	// switch letter {
	// case "a", "e", "i", "o", "u":
	// 	fmt.Println("nguyên âm")
	// default:
	// 	fmt.Println("phụ âm")
	// }

	switch num := number(); { //num không phải là một hằng số
	case num < 50:
		fmt.Printf("%d thì nhỏ hơn 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d thì nhỏ hơn 100\n", num)
		fallthrough //chỉ nhảy xuống kiểm tra 1 case phía sau nó không xuống cái tiếp theo nữa,phải là lệnh cuối trong case
	case num < 200:
		fmt.Printf("%d thì nhỏ hơn 200", num)
	case num < 300:
		fmt.Printf("%d thì nhỏ hơn 200", num)
	}
}
