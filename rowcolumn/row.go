// to calculate the total no of row and column 

package main 

import "fmt"

func main() {

	var r , c , e int 

	fmt.Println("Enter the no of row:")
	fmt.Scan(&r)

	fmt.Println("Enter the no of:")
	fmt.Scan(&c)

	fmt.Println("Enter the extra no of row:")
	fmt.Scan(&e)

	a := (r + e) * c 
	fmt.Println("The total no of cell:", a)
}
