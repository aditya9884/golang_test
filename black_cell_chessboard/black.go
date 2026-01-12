//Given (n is even), determine the number of black cells in an 
//n×n chessboard.

package main 

import "fmt"

func main() {
	var n int 

	fmt.Println("Please enter number:")
	fmt.Scan(&n)

	if n % 2 == 0 {
		
		z := (n * n)/2
		fmt.Println("The number of black cell in chessboard is:", z)

	}else {
		fmt.Println("The chessboard can not be odd")
	}
}