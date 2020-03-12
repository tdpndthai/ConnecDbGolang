package main

import (
	"fmt"
	"reflect"
	"structs/computer" //package import phải để lên đầu tiên @@
)

type image struct {  
    data map[int]int
}


type name struct {  
    firstName string
    lastName string
}

type person struct{
	//string
	int
}

func  main()  {
	var spec computer.Spec

	spec.Maker = "thái"
	//spec.model lỗi
	spec.Price = 29
	fmt.Println("spec",spec)
	//Cấu trúc bằng nhau (Structs Equality)
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	res := reflect.DeepEqual(name1,name2)
	if res==true {
		fmt.Println("name1 name2 trùng nhau")
	}else{
		fmt.Println("name1 name2 không trùng")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
	}
	

	image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 154,
	}}
	//chứa kiểu dữ liệu map nên ko so sánh được
    // if image1 == image2 {
    //     fmt.Println("image1 and image2 are equal")
	// }
	res1 :=  reflect.DeepEqual(image1,image2)
	if res1 == true {
		fmt.Println("img1 và img2 trùng nhau")
	}else{
		fmt.Println("không trùng nhau")
	}

	p1 := person{29}
	p2 := person{19}
	
	if p1 == p2 {
		fmt.Println("p1 trùng p2")
	}else{
		fmt.Println("p1 khác p2")
	}

}