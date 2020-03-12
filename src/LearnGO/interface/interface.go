package main

import (
	"fmt"
)

// //định nghĩa 1 interface
// type VowelsFinder interface {  
//     FindVowels() []rune
// }

// type MyString string

// //MyString implements VowelsFinder

// func (ms MyString) FindVowels() []rune {  
//     var vowels []rune
//     for _, rune := range ms {
//         if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
//             vowels = append(vowels, rune)
//         }
//     }
//     return vowels
// }


// type SalaryCalculator interface {  //tên interface
//     CalculateSalary() int //tên thuộc tính
// }

// type Permanent struct {  
//     empId    int
//     basicpay int
//     pf       int
// }

// type Contract struct {  
//     empId  int
//     basicpay int
// }
// // cả 2 struct Permanent và Contract đều đang implement interface SalaryCalculator.

// //tiền lương của nhân viên permanent bằng tổng của basic pay và pf
// func (p Permanent) CalculateSalary() int {  
//     return p.basicpay + p.pf
// }

// //tiền lương của nhân viên contract chỉ là basic pay
// func (c Contract) CalculateSalary() int {  
//     return c.basicpay
// }

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
// func totalExpense(s []SalaryCalculator) { //s chứa mảng salarycalculator
//     expense := 0
//     for _, v := range s { //duyệt từng value trong mảng 
//         expense = expense + v.CalculateSalary()
//     }
//     fmt.Printf("Total Expense Per Month $%d", expense)
// }


// type Tester interface {  //interface Tester
//     Test() //hàm
// }

// type MyFloat float64

// func (m MyFloat) Test() {  //triển khai hàm Test() dựa trên kiểu dữ liệu tự định nghĩa
//     fmt.Println(m)
// }

// func describe(t Tester) {  
//     fmt.Printf("Interface type %T value %v\n", t, t)
// }

// Kiểu interface rỗng
// Một interface không có phương thức nào thì được gọi là interface rỗng. 
// Nó được biểu diễn dưới dạng interface{}. Vì interface rỗng không có thương thức, 
// nên tất cả các kiểu dữ liệu đều có thể implement interface rỗng.
// func describe(i interface{}) {  
//     fmt.Printf("Type = %T, value = %v\n", i, i)
// }

// func assert(i interface{}) {  
// 	//s := i.(int) //ép giá trị int cho interface
// 	//nếu giá trị không phải kiểu int
// 	v,ok := i.(int) //ép kiểu giá trị cơ bản của interface i về kiểu cụ thể là int
//     fmt.Println(v,ok)
// }


// func findType(i interface{}) {  
//     switch i.(type) { //khai báo sử dụng switch với kiểu dữ liệu
//     case string: //so sánh kiểu dữ liệu của i với kiểu dữ liệu được chỉ định,dòng này là string
//         fmt.Printf("I am a string and my value is %s\n", i.(string))
//     case int:
// 		fmt.Printf("I am an int and my value is %d\n", i.(int))
// 	case float64: 
// 		fmt.Printf("I am an float and my value is %0f\n", i.(float64))
//     default:
//         fmt.Printf("Unknown type\n")
//     }
// }

type Describer interface {  //interface describer
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  //method triển khai interface,truyền vào struct person
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  //hàm kiểm tra truyền vào 1 interface rỗng 
    switch v := i.(type) { //lấy ra type của interface truyền vào
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    // name := MyString("Sam Anderson")
    // var v VowelsFinder
    // v = name // Điều này được chấp nhận vì MyString implement interface VowelsFinder
	// fmt.Printf("Vowels are %c", v.FindVowels())
	
	// pemp1 := Permanent{1, 5000, 20}
    // pemp2 := Permanent{2, 6000, 30}
    // cemp1 := Contract{3, 3000}
    // employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	// totalExpense(employees)
	
	// var t Tester
    // f := MyFloat(89.7)
    // t = f //data của t là 89.7
    // describe(t)
	// t.Test()
	
	//hàm describe(i interface{}) lấy một interface rỗng làm đối số do đó nó có thể truyền vào bất kỳ kiểu dữ liệu nào.
	// s := "Hello World"
    // describe(s)
    // i := 55
    // describe(i)
    // strt := struct {
    //     name string
    // }{
    //     name: "Naveen R",
    // }
	// describe(strt)
	
	//Type assertion được sử dụng để ép kiểu giá trị cơ bản của interface.
	// var s interface{} = 56 //biến s đang là interface có data kiểu int
	// assert(s) //56 true
	// var i interface{} = "Steven Paul"
	// assert(i) //0 false
	
	// Trường hợp switch với biểu thức tùy chọn là kiểu dữ liệu được dùng để so sánh kiểu dữ liệu cụ thể 
	// của một interface với nhiều kiểu dữ liệu được chỉ định trong các  trường hơp khác nhau. 
	// Nó tương tự như trường hợp switch case.

	// Cú pháp cho switch với kiểu dữ liệu cũng tương tự như cú pháp i.(T) trong Type assertion. 
	// Kiểu T ở đây chính là các từ khóa của các kiểu dữ liệu mà chúng ta muốn so sánh. Hãy theo dõi chương trình dưới đây.

	// findType("Naveen")
    // findType(77)
	// findType(89.98)
	
	findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}

