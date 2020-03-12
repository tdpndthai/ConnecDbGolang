package main

import (  
	"fmt"
	"unicode/utf8"
)

// func printBytes(s string) {  
//     for i:= 0; i < len(s); i++ {
//         fmt.Printf("%x ", s[i])
//     }
// }

// func printChars(s string) {  
//     // for i:= 0; i < len(s); i++ {
//     //     fmt.Printf("%c ",s[i])
// 	// }
// 	runes := []rune(s) //in 1 vài kí tự đặc biệt 
// 	for i:= 0; i < len(runes); i++ {
// 		fmt.Printf("%c ",runes[i])
// 	}
// }

func length(s string) {  
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
 }

func printCharsAndBytes(s string) {  
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func mutate(s []rune) string {  
	s[0] = 'a' //s[0] được thay đổi bằng byte của 'a'
	return string(s) //chuyển đổi lại rune thành string
 }

//  func mutate(s string)string {  
// 	s[0] = 'a' // bất kỳ ký tự unicode hợp lệ trong ngoặc nháy đơn đều là một rune 
// 	return s
//  }

func main() {  
    // name := "Hello World"
	// printBytes(name)
	// fmt.Printf("\n")
	// printChars(name)
	// name = "Señor"
	// printBytes(name)
	// fmt.Printf("\n")
	// printChars(name) //in ra S e Ã ± o r bị lỗi cần -> rune

	// name := "Señor"
	// printCharsAndBytes(name)

	//Thiết lập string từ slice chứa các byte
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str) //Café

	//thay mã dạng hex bằng dạng decimal tương ứng
	// byteSlice := []byte{67, 97, 102, 195, 169} // tương đương phần thập phân của {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	// str := string(byteSlice)
	// fmt.Println(str) //Café

	//độ dài của chuỗi string
	word1 := "Señor" 
	length(word1)
	word2 := "Pets"
	length(word2)

	//Kiểu dữ liệu string là bất biến
	//Kiểu dữ liệu string là bất biến, có nghĩa là, một khi string được tạo ra, nó không thể bị thay đổi. Hãy cùng xem ví dụ dưới đây:
	// h := "hello"
	// fmt.Println(mutate(h))

	//Để làm việc được với string, chúng ta chuyển string thành slice chứa các rune, 
	//chúng ta thay đổi slice và sau đó chuyển lại thành string mới.
	h := "hello"
	fmt.Println(mutate([]rune(h))) //truyền vào function slice chứa các rune trả về aello
}