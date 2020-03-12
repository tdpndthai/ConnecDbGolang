package main

import (  
    "fmt"
)

// func changeLocal(num [5]int) {  
//     num[0] = 55
//     fmt.Println("inside function ", num)

// }

func printarray(a [3][2]string) {  
    for _, v1 := range a { //lặp qua dòng
        for _, v2 := range v1 { //lặp qua cột
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}


func main() {  
    // var a [3]int //khai báo mảng nguyên có độ dài bằng 3,chưa khởi tạp giá trị nên các phần tử mảng = 0
	// fmt.Println(a)
	
	// var a [3]int //int array with length 3
    // a[0] = 12 // array index starts at 0
    // a[1] = 78
    // a[2] = 50
	// fmt.Println(a)
	
	//khai báo mảng ngắn
	//a := [3]int{12, 78, 50} // short hand declaration to create array
	//a := [3]int{12} //gán giá trị vào phẩn tử đầu trong mảng,không nhất thiết phải khai báo tất cả giá trị cho các phần tử mảng [12 0 0]
	//fmt.Println(a)

	//chú ý: mảng không thể thay đổi kích cỡ
	// a := [3]int{5, 78, 8}
    // var b [5]int
	// b = a //not possible since [3]int and [5]int are distinct types
	
	// Mảng trong Go là các loại giá trị và không phải là loại tham chiếu. 
	// Điều này có nghĩa là khi chúng được gán cho một biến mới, một bản copy của bản gốc sẽ được gán cho biến mới này. 
	// Nếu thay đổi được thực hiện trên biến mới, nó sẽ không xét lại trong mảng ban đầu.
	// a := [...]string{"USA", "China", "India", "Germany", "France"}
    // b := a // a copy of a is assigned to b
    // b[0] = "Singapore"
    // fmt.Println("a is ", a)
	// fmt.Println("b is ", b) 
	

	//Tương tự, khi mảng được truyền cho các hàm như các tham số, chúng được truyền theo giá trị và mảng ban đầu không thay đổi.
	// num := [...]int{5, 6, 7, 8, 8} //bỏ qua chiều dài của mảng thay bằng [...]
    // fmt.Println("before passing to function ", num)
    // changeLocal(num) //num is passed by value
	// fmt.Println("after passing to function ", num)
	
	//độ dài mảng
	// a := [...]float64{67.7, 89.8, 21, 78}
	// fmt.Println("length of a is",len(a))
	
	//Vòng lặp mảng sử dụng range (phạm vi)
	// a := [...]float64{67.7, 89.8, 21, 78}
    // for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
    //     fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    // }


	a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}