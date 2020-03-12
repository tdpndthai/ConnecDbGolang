package main

import (
	"fmt"
)

// func find(num int, nums ...int) { //num=item, nums = list ,kiểm tra item có xuất hiện trong list hay không
//     //chuyển nums -> int[]{} : như 1 slice mới và thao tác với nó
// 	fmt.Printf("Loại của số là: %T\n", nums)
// 	found := false
// 	for index, value := range nums {
// 		if value == num {
// 			fmt.Println(num, "tìm thấy tại ví trí", index, "trong", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "không tìm thấy trong", nums)
// 	}
// 	fmt.Printf("\n")

// }

// func addItem(item int,list ...int){
//     list = append(list,item)
//     fmt.Println(list)
// }

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

// func main() {
// 	find(89, 89, 90, 95)
// 	find(45, 56, 67, 45, 90, 109)
// 	find(78, 38, 56, 98)
// 	find(87)
// }

func main(){
    //vì variadic param nó convert tham số truyền vào thành 1 slices thế nên nếu muốn truyền vào 1 slice thì đặt dấu '...' sau slices
    //nums := []int{89, 90, 95}
    //find(89, nums...)
    //addItem(45,nums...)

    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)
}




//Variadic function (Hàm bất định) là hàm chấp nhận một đối số có biến kiểu số (variable number of arguments)
//Nếu tham số cuối cùng của một hàm được ký hiệu bằng ...T, hàm đó sẽ chấp nhận mọi đối số chứa số (number) thuộc kiểu T cho tham số cuối.
