package mypackage 
import "secondpackage"
import "fmt"
func MyFunction(n int) int { 
n = n+1
fmt.Println("square of give number plus 1 is", secondpackage.Simplesquare(n))                    
	if n == 0 {
		return 1
	} else {
		return n + MyFunction(n-1)
	}
}


