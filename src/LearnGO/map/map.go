package main

import (
	"fmt"
	"reflect"
)


func main()  {
	//cú pháp khởi tạp var abc = make(map[type of key]type of value)
	
	// var personSalary map[string]int //khai báo thông thường map sẽ là nil (=null)
	// if personSalary == nil {
	// 	fmt.Println("map chưa khởi tạo")
	// 	personSalary = make(map[string]int)
	// }

	//cách 1
	// personSalary := make(map[string]int)  //khởi tạo 1 map với loại key là string,loại của giá trị là int
    // personSalary["steve"] = 12000
    // personSalary["jamie"] = 15000
	// personSalary["mike"] = 9000
	//cách 2 vừa khởi tạo map vừa khai báo
	// personSalary := map[string]int {
    //     "steve": 12000,
    //     "jamie": 15000,
    // }
    // personSalary["thái"] = 9000
	// fmt.Println("personSalary map contents:", personSalary)
	

	//Truy cập các phần tử của map
	// personSalary := map[string]int{
	// 	"thái" : 1500,
	// 	"sơn" : 1200,
	// }
	// personSalary["and"] = 1300
	//emm := "thái"
	//kiểm tra xem key có tồn tại trong map personSalary
	// value,ok := personSalary[emm]//nếu có key trong map thì ok trả về giá trị true,value trả về value của key
	// if ok ==true {
	// 	fmt.Println("lương của",emm,"là",value)
	// } else{
	// 	fmt.Println(emm,"not found")
	// }
	//fmt.Println("lương của",emm,"là",personSalary[emm])// nếu không có emm = "thái" thì giá trị trả về sẽ là bằng 0

	//Dạng range của vòng lặp for
	// fmt.Println("all item of map")
	// for key,value := range personSalary {
	// 	fmt.Printf("personSalary[%s] = %d\n",key,value)
	// }

	//XÓA PHẦN TỬ: CÚ PHÁP delete(map, key) Hàm delete này không trả về bất kỳ giá trị nào.

	// personSalary["mike"] = 9000
    // fmt.Println("map before deletion", personSalary)
    // delete(personSalary, "mike")
	// fmt.Println("map after deletion", personSalary)
	// fmt.Println("chiều dài của map",len(personSalary)) //chiều dài của map
	
	//Map là kiểu tham chiếu
	//Khi một map được gán cho một biến mới, cả hai đều trỏ đến cùng một cấu trúc dữ liệu nội bộ
	//Do đó những thay đổi được thực hiện ở một biến sẽ ánh xạ đến biến kia.
	// fmt.Println("original person salary",personSalary)
	// newpersonSalary := personSalary
	// newpersonSalary["and"] = 1400 
	// fmt.Println("new",personSalary)
	// fmt.Println("new",newpersonSalary)

	//Map không thể so sánh với nhau bằng toán tử ==. Ta chỉ có thể sử dụng == để kiểm tra xem một map có phải là nil hay không.

	map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

	//map2 := map1
	map2 := map[string]int{
		"ba": 2,
		"badg": 3,
	}

    // if map1 == map2 {
	// }
	//muốn so sánh 2 map với nhau dùng reflect.DeepEqual(map1,map2)

	res := reflect.DeepEqual(map1,map2)
	fmt.Println("2 map có trùng nhau không?",res)
	
}