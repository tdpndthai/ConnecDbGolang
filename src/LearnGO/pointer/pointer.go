package main

import (  
    "fmt"
)

// func change(val *int) {   //truyền value của biến b,tham chiếu ngược tới địa chỉ và value của a
//     *val = 55 //thay đổi a bằng cách tham chiếu ngược 
// }

// func modify(arr *[3]int) {   //tham chiếu ngược địa chỉ biến lấy ra ra value
// 	(*arr)[0] = 90 //gán 90 cho thành phần đầu tiên của arr
// 	//có thể viết thành 
// 	//arr[0] = 90
// }

func modify(sls []int) {  
    sls[0] = 90
}

func main() {  
    // b := 255
    // var a *int = &b //gán địa chỉ của biến b cho a có kiểu int
    // fmt.Printf("Type of a is %T\n", a)
	// fmt.Println("address of b is", a)
	
	//biến b được khởi tạo là nil, sau đó được gán vào địa chỉ của biến a. Ví dụ trên sẽ xuất ra:
	// a := 25
    // var b *int
    // if b == nil {
    //     fmt.Println("b is", b) //nil
    //     b = &a
    //     fmt.Println("b after initialization is", b)// 0xc0000100b0
	// }
	

	//tham chiếu ngược : truy cập vào biến mà con trỏ trỏ tới *(biến) 
	// b := 255
    // a := &b //gán địa chỉ của biến b cho a
    // fmt.Println("address of b is", a)
	// fmt.Println("value of b is", *a) //tham chiếu lại giá trị của biến b 
	
	// b := 255
    // a := &b //địa chỉ của b là địa chỉ của a,giá trị của b là giá trị của a
    // fmt.Println("address of b is", a)
    // fmt.Println("value of b is", *a)
    // *a++ //thay đổi giá trị a dẫn đến b cũng thay đổi
	// fmt.Println("new value of b is", b)
	
	//Sử dụng con trỏ trong hàm
    // a := 58
    // fmt.Println("value of a before function call is",a)
    // b := &a
    // change(b)
	// fmt.Println("value of a after function call is", a)

	//Hãy sử dụng con trỏ cho slice thay vì array trong hàm
	// a := [3]int{89, 90, 91}
    // modify(&a) //lấy địa chỉ biến a
	// fmt.Println(a)
	
	a := [3]int{89, 90, 91}
    modify(a[:])
	fmt.Println(a)
	
	//Go không hỗ trợ phép toán số học trên con trỏ
	// b := [...]int{109, 110, 111}
    // p := &b
    // p++
}