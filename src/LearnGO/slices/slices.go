package main

import (
	"fmt"
)

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}

// func subtactOne(numbers []int) {  
//     for i := range numbers {
//         numbers[i] -= 2
//     }

// }

//slices tham chiếu đến các mảng hiện có
func main() {
	// a := [5]int{76, 77, 78, 79, 80}
	// var b []int = a[1:4] //creates a slice from a[1] to a[3]
	// fmt.Println(b)

	//1 cách khác để tạo slices
	// c := []int{6, 7, 8} //creates and array and returns a slice reference
	// fmt.Println(c)

	// Một slice không sở hữu bất kì một dữ liệu nào của riêng nó. Nó chỉ là đại diện cho một mảng cơ bản.
	// Bất kì một sửa đổi được thực hiện với slice sẽ được xét lại trong mảng cơ bản.
	// darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// dslice := darr[2:5] //ds slices từ darr[2] -> darr[4]
	// fmt.Println("array before", darr)
	// for i := range dslice { //lặp từng phần tử trong dslice mỗi phần tử tăng 1 đơn vị
	// 	dslice[i]++
	// }
	// fmt.Println("array after", darr) //mảng dslice thay đổi sẽ được xét lại trong mảng [57 89 91 83 101 78 67 69 59]

	// Khi một số slices nằm trong một mảng cơ bản, mỗi thay đổi tạo rã sẽ được xét lại trong mảng.
	// numa := [3]int{78, 79 ,80}
	// nums1 := numa[:] //tạo slice chứa tất cả các phần tử trong mảng
	// nums2 := numa[:]
	// fmt.Println("array before change 1",numa)
	// nums1[0] = 100
	// fmt.Println("array after modification to slice nums1", numa) //[100 79 80]
	// nums2[1] = 101
	// fmt.Println("array after modification to slice nums2", numa)// [100 101 80]

	// Độ dài của một slice chính là số lượng của các phần tử có trong slice đó.
	// Sức chứa của slice là chỉ số lượng phần tử cơ bản bắt đầu từ mục mà slice đó được tạo ra.
	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

	// 	Ở chương trình trên, fruitslice đã được tạo ra từ mục 1 và 2 của mảng fruitarray. Do đó độ dài của fruitslice là 2.
	// Độ dàicủa fruitarray là 7. fruitslice được tạo ra từ mục 1 của fruitarray.
	// Do đó sức chứa của fruitslice không có phần tử nào trong fruitarray bắt đầu từ mục 1, ví dụ từ orange có 6 kí tự do đó sức chứa của fruitslice là 6.
	// Chương trình này có đầu ra độ dài là 2 và sức chứa là 6.

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	// fruitslice = fruitslice[:cap(fruitslice)] //re-slice sẽ tạo slice với tất cả các phần tử trong capacity
	// //(lúc này phần tử trong fruitslice từ orange -> chikoo) đúng bằng capacity của fruitslice trước đó
	// fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

	// Tạo slice sử dụng make
	// func make([]T, len, cap) []T có thể được sử dụng để tạo slice bằng cách truyền vào type (loại), độ dài và sức chứa.
	// Tham số là tùy chọn và độ dài mặc định. Function make tạo ra một mảng và trả về một slice tham chiếu.
	//Các giá trị mặc định là 0 khi một slice được tạo sử dụng make. Chương trình trên sẽ có đầu ra [0 0 0 0 0].

	// i := make([]int, 5, 5)
	// fmt.Println(i)
	// nếu cap của slide nil thì golang tự gán cap vào len luôn

	// 	Thêm một slice
	// Như chúng ta đã biết là mảng bị giới hạn độ dài cố định và đồ dài của chúng không thể tăng lên được. Slices có sự linh động và các phần tử mới có thể thêm vào bằng cách sử dụng function append.
	// Hàm này có dạng func append(s []T, x ...T) []T


	//Khi các phần tử mới được nối vào thì một mảng mới cũng sẽ được tạo ra. 
	//Phần tử của mảng hiện có được sao chép vào mảng mới này và tham chiếu slice mới cho mảng này được trả về
	//Sức chứa của slice mới hiện tại đã nhiều gấp đôi slice cũ.,có append thêm phần tử nữa thì sức chứa của nó vẫn chỉ là giá trị gấp đôi slice cũ
	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3,len = 3
	// cars = append(cars, "Toyota")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6,len = 4,
	// cars = append(cars,"thaind")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	//Giá trị zero (không) của một loại slice là nil. nil slice có độ dài và sức chứa bằng 0. 
	//Có thể thêm một giá trị vào một slice nil bằng cách sử dụng function append
	// var names []string //zero value of a slice is nil , như kiểu khởi tạo slices null 
    // if names == nil {
    //     fmt.Println("slice is nil going to append")
    //     names = append(names, "John", "Sebastian", "Vinay","thai")
    //     fmt.Println("names contents:",names)
	// }
	
	//Cũng có thể thêm một slice này đến một slice khác bằng cách sử dụng toán tử ... 
    // veggies := []string{"potatoes","tomatoes","brinjal"}
	// fruits := []string{"oranges","applessss"}
	// fullname := []string{"thai","nguyen"}
	// food := append(veggies, fruits...)
	// test := append(fullname,fruits...)
	// fmt.Println("food:",food) //[potatoes tomatoes brinjal oranges apples]
	// fmt.Println("test: ",test,"cap test",cap(test))
	
	// arr := [...] int {1,2,3,4,5}
	// slices := arr[2:] //lấu từ giá trị index = 2 trở về trước vd 2: -> từ index 2 trở về sau
	// fmt.Println("slices",slices)

	//Các slice có thể coi như được biêủ thị bằng một cấu trúc. Trông nó như thế này
	// type slice struct {  
	// 	Length        int
	// 	Capacity      int
	// 	ZerothElement *byte
	// }

	//slice chuyển đến 1 function, mặc dù nó được truyền theo giá trị, con trỏ sẽ tham chiếu đến cùng một mảng cơ bản. 
	// Do đó khi slice chuyển một function như một tham số, các thay đổi bên trong function cũng sẽ hiển thị bên ngoài function.
	// nos := []int{8, 7, 6}
    // fmt.Println("slice before function call", nos) //[8 7 6]
    // subtactOne(nos)                               //function modifies the slice
	// fmt.Println("slice after function call", nos) //modifications are visible outside,[6 5 4]

	countriesNeeded := countries()
    fmt.Println(countriesNeeded)	
}
