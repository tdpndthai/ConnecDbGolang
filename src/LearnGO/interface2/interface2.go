package main

import "fmt"

// //Triển khai Interfaces sử dụng con trỏ và giá trị

type Describer interface {  //tạo mới interface
    Describe()
}
// type Person struct {  
//     name string
//     age  int
// }

// func (p Person) Describe() { //implemented using value receiver  
//     fmt.Printf("%s is %d years old\n", p.name, p.age)
// }

// type Address struct {  
//     state   string
//     country string
// }

// func (a *Address) Describe() { //implemented using pointer receiver  
//     fmt.Printf("State %s Country %s", a.state, a.country)
// }

//triển khai nhiều interface
type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

//nhúng interface 
//Bất kỳ loại nào được cho là để triển khai Interfaces EmployeeOperations nếu nó cung cấp các định nghĩa phương thức cho các phương thức 
//có trong cả hai Interfaces SalaryCalculator và LeaveCalculator.
type EmployeeOperations interface {  
    SalaryCalculator
    LeaveCalculator
}
//Employee struct triển khai thực hiện giao diện EmployeeOperations
//vì nó cung cấp định nghĩa cho cả hai phương thức DisplaySalary và CalculateLeavesLeft.
type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

//Cấu trúc Employee được định nghĩa cung cấp các triển khai cho phương thức DisplaySalary của Interfaces SalaryCalculator 
//và phương thức CalculateLeavesLeft của Interfaces LeaveCalculator.

func (e Employee) DisplaySalary() {  //implement SalaryCalculator
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  //implement LeaveCalculator
    return e.totalLeaves - e.leavesTaken
}



func main() {  
    // var d1 Describer //biến d1 với kiểu dữ liệu Descrriber
    // p1 := Person{"Sam", 25} //p1 :kiểu giá trị của Person nó được gán cho d1
	// d1 = p1	
    // d1.Describe() //Sam is 25 years old 
    // p2 := Person{"James", 32}
    // d1 = &p2 //gán địa chỉ,data của p2 cho d1
    // d1.Describe() //James is 32 years old 

    // var d2 Describer
    // a := Address{"Washington", "USA"}
    // /* compilation error if the following line is
    //    uncommented
    //    cannot use a (type Address) as type Describer
    //    in assignment: Address does not implement
    //    Describer (Describe method has pointer
    //    receiver)
    // */
    // //d2 = a

    // d2 = &a //This works since Describer interface
    // //is implemented by Address pointer in line 22
	// d2.Describe()
	
	//triển khai nhiều interface
	e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
	}
	// Biến e của kiểu Employee được gán cho empOp của kiểu EmployeeOperations. 
	// Tiếp theo, các phương thức DisplaySalary () và CalculateLeavesLeft () được gọi trên empOp.

	var empOp EmployeeOperations = e
    var s SalaryCalculator = e
    s.DisplaySalary()
    // var l LeaveCalculator = e
	// fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
	//Giá trị 0 của Interfaces
	//Giá trị bằng không của một Interfaces là  nil. Một Interfaces nil có cả giá trị cơ bản và cụ thể là nil.
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 là nil và có kiểu %T giá trị %v7\n",d1,d1)
	}
}