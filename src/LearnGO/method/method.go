package main

//"math"
import (
	"fmt"
)

// type Employee struct {
// 	name     string
// 	salary   int
// 	currency string
// }

// type Employee struct {
//     name string
//     age  int
// }

//tạo phương thức displaySalary(truyền vao Employee)

// func (e Employee) displaySalary() {
// 	fmt.Printf("salary of %s is %d %s", e.name, e.salary, e.currency)
// }

/*
 phương thức displaySalary() được chuyển thành hàm displaySalary với Employee là tham số truyền vào
*/
// func displaySalary(e Employee) {
//     fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
// }

//hoàn toàn có thể tạo ra phương thức tên giống nhau với kiểu dữ liệu truyền vào khác nhau
// type Rectangle struct {
//     length int
//     width  int
// }

// type Circle struct {
//     radius float64
// }

// func (r Rectangle) Area() int {
//     return r.length * r.width
// }

// func (c Circle) Area() float64 {
//     return math.Pi * c.radius * c.radius
// }

// /*
// Phương thức với vật nhận là giá trị
// */
// func (e Employee) changeName(newName string) {
//     e.name = newName
// }

// /*
// Phương thức với vật nhận là con trỏ
// */
// func (e *Employee) changeAge(newAge int) {
//     e.age = newAge
// }

// type address struct {
//     city  string
//     state string
// }

// func (a address) fullAddress() {
//     fmt.Printf("Full address: %s, %s", a.city, a.state)
// }

// type person struct {
//     firstName string
//     lastName  string
//     address //person có trường address này là 1 struc khác nên có thể gọi được phương thức fullAddress
// }

// type rectangle struct {
//     length int
//     width  int
// }

// func area(r rectangle) {
//     fmt.Printf("Area Function result: %d\n", (r.length * r.width))
// }

// func (r rectangle) area() {
//     fmt.Printf("Area Method result: %d\n", (r.length * r.width))
// }

// type rectangle struct {
//     length int
//     width  int
// }

// func perimeter(r *rectangle) {
//     fmt.Println("perimeter function output:", 2*(r.length+r.width))

// }

// func (r *rectangle) perimeter() {
//     fmt.Println("perimeter method output:", 2*(r.length+r.width))
// }

// func (a int) add(b int) {  //lỗi vì phương thức add và kiểu int không được định nghĩa trên cùng 1 package,int là kiểu dựng sẵn
//Cách để giải quyết vấn đề này là tạo một kiểu dữ liệu là bí danh cho kiểu dựng sẵn int, sau đó tạo một phương thức mà kiểu dữ liệu bí danh này chính là vật nhận.
// }
//chúng ta đã tạo một kiểu dữ liệu bí danh là myInt cho kiểu int.
//Sau đó, ta định nghĩa phương thức add với myInt chính là vật nhận.
type myInt int //cách giải quyết
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {

	// emp1 := Employee{"nguyễn danh thái", 29, "VND"}
	// emp1.displaySalary()
	//displaySalary(emp1)

	// 	Go không phải là một ngôn ngữ lập trình hướng đối tượng thuần túy và nó không hỗ trợ các lớp (class). Do đó viết các phương thức trên các kiểu dữ liệu là một cách để có được hành vi tương tự như các lớp.
	// Các phương thức có thể trùng tên nếu được xác định trên các kiểu dữ liệu khác nhau trong khi các hàm thì không được phép trùng tên. Giả sử chúng ta có 2 struct là Square và Circle.
	// Chúng ta có thể định nghĩa 2 phương thức cùng có tên là Area trên cả Square và Circle. Hãy xem chương trình sau.

	// r := Rectangle{
	//     length: 10,
	//     width:  5,
	// }
	// fmt.Printf("Area of rectangle %d\n", r.Area())
	// c := Circle{
	//     radius: 12,
	// }
	// fmt.Printf("Area of circle %f", c.Area())

	// e := Employee{
	//     name: "nguyễn danh thái",
	//     age:  22,
	// }
	// fmt.Printf("Employee name before change: %s", e.name)
	// e.changeName("Michael Andrew") //chạy qua phương thức là name của e lại trở về như cũ
	// fmt.Printf("\nEmployee name after change: %s", e.name)

	// fmt.Printf("\n\nEmployee age before change: %d", e.age)
	// (&e).changeAge(51)
	// fmt.Printf("\nEmployee age after change: %d", e.age)

	//Phương thức của trường ẩn danh
	//Các phương thức thuộc về một trường ẩn danh nằm trong một struct nào đó có thể được gọi như thể chúng là của chính struct đó.

	// p := person{
	//     firstName: "Elon",
	//     lastName:  "Musk",
	//     address: address {
	//         city:  "Los Angeles",
	//         state: "California",
	//     },
	// }
	// //câu lệnh tưởng mình
	// //p.address.fullAddress(),
	// p.fullAddress()

	// r := rectangle{
	//     length: 10,
	//     width:  5,
	// }
	// area(r)
	// r.area()

	// p := &r
	// /*
	//    lỗi biên dịch, cannot use p (type *rectangle) as type rectangle
	//    in argument to area
	// */
	// //area(p)

	// p.area()//dùng con trỏ p gọi đến phương thức area() có vật nhận là giá trị
	//Ta thấy hàm func area(r rectangle) nhận đối số là giá trị và phương thức func (r rectangle) area() nhận vật nhận là giá trị.

	// Chúng ta gọi đến hàm area với đối số là giá trị: area(r) và nó hoạt động bình thường.
	// Tương tự, chúng ta gọi đến phương thức area sử dụng vật nhận là giá trị: r.area() và nó cũng hoạt động bình thường.
	//Sau đó, chúng ta khởi tạo 1 con trỏ p trỏ tới r: p := &r. Nếu chúng ta cố truyền con trỏ này vào hàm area (hàm area chỉ chấp nhận giá trị), thì trình biên dịch sẽ báo lỗi.
	//dòng code p.area() gọi đến phương thức area, phương thức area vốn được khai báo nhận vật nhận là giá trị giờ lại sử dụng vật nhận là con trỏ p.
	//Và điều này hoàn toàn hợp lệ. Lý do là dòng code p.area() sẽ được Go ngầm hiểu là (*p).area().

	//Tương tự như với đối số là giá trị, một hàm có đối số là con trỏ sẽ chỉ chấp nhận con trỏ (line 94)
	//trong khi phương thức có vật nhận là con trỏ sẽ chấp nhận cả vật nhận là giá trị lẫn vật nhận là con trỏ.
	// r := rectangle{
	//     length: 10,
	//     width:  5,
	// }
	// p := &r //pointer to r
	// perimeter(p)
	// p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	//r.perimeter()//dùng giá trị r gọi đến phương thức perimeter() có vật nhận là con trỏ

	//Phương thức đối với kiểu dữ liệu non-struct
	//Để định nghĩa một phương thức trên một kiểu dữ liệu,
	//thì phương thức và kiểu dữ liệu của vật nhận của phương thức đó phải được định nghĩa trên cùng một package.
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
