package main

import (
	"fmt"
)

// Structures (cấu trúc) được người dùng định nghĩa cho một tập hợp các trường.
// Nó có thể được sử dụng để nhóm dữ liệu vào một cấu trúc thay vì viết từng loại riêng biệt dưới các dạng riêng biệt.
//Employee là tên cấu trúc 
// type Employee struct {  
// 	lastName, firstName string
//     age,salary int
// }

//structures có trường ẩn danh
// type Person struct {  //2 trương vô danh có kiểu là string và int
//     string
//     int
// }

//Cấu trúc lồng nhau (Nested structs​)
// type Address struct {  
//     city, state string
// }
// type Person struct {  
//     name string
//     age int
//     address Address //chứa 1 trường là 1 cấu trúc
// }


//Các trường được khuyến khích (Promoted fields​)
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age  int
    Address
}


//Có thể khai báo cấu trúc mà không khai báo một kiểu mới và kiểu cấu trúc này được gọi là cấu trúc ẩn danh.
/*var employee struct {  
	firstName, lastName string
	age int
}*/
func main() {

    //creating structure using field names
    // emp1 := Employee{//không cần duy trì thức tự tên trường
    //     firstName: "Sam",
    //     age:       25,
    //     salary:    500,
    //     lastName:  "Anderson",
    // }

    //creating structure without using field names
    // emp2 := Employee{"Thomas", "Paul", 29, 800} //duy trì thứ tự tên trường
	// emp3 := Employee{"ádfa","adf",12,35}
    // fmt.Println("Employee 1", emp1)
    // fmt.Println("Employee 2", emp2)
    // fmt.Println("Employee 3 firstname", emp3.firstName)

    //Cấu trúc ẩn danh
    // var emp4 Employee 
    // fmt.Println("emm 4",emp4) //emm 4 {  0 0} firstname và lastname null nên sẽ không hiện,kiểu int thì in ra mặc định là 0 

    //Cấu trúc 0 (Zero value of a structure)

    // emp5 := Employee{
    //     firstName: "John",
    //     //lastName:  "Paul",
    // }
    // fmt.Println("Employee 5", emp5)

    //Trỏ đến một cấu trúc (Pointers to a struct)
        //cách 1
    // emp8 := &Employee{"Sam", "Anderson", 55, 6000} //gán giá trị và địa chỉ của employee này cho emp8 
    // fmt.Println("First Name:", (*emp8).firstName)
    // fmt.Println("Age:", (*emp8).age)
    //cách 2
    // emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    // fmt.Println("First Name:", emp8.firstName)
    // fmt.Println("Age:", emp8.age)

    // p := Person{"thái",22}
    // fmt.Println(p)

    //theo mặc định, tên của trường ẩn danh là tên loại của nó.
    //-> structs Person có 2 trường có tên string và int.
    // var p1 Person 
    // p1.string = "sơn"
    // p1.int = 19
    // fmt.Println(p1)

    //Cấu trúc lồng nhau (Nested structs​)
    //Có thể là một cấu trúc chứa một trường mà lần lượt là một cấu trúc. Những loại cấu trúc này được gọi là cấu trúc lồng nhau.


    // var p Person
    // p.name = "thái"
    // p.age = 22
    // p.address=Address{//cách 1
    //     //city: "hà nội",
    //     state: "test",
    // }
    // p.address.city = "tphcm" //cách 2
    // fmt.Println("Name:", p.name)
    // fmt.Println("Age:",p.age)
    // fmt.Println("City:",p.address.city)
    // fmt.Println("State:",p.address.state)

    
    //Các trường được khuyến khích (Promoted fields​)
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    p.city = "hà nội"
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field

    //Xuất cấu trúc (Exported Structs and Fields​)
}