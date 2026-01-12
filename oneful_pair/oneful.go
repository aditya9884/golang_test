// Chef defines a pair of positive integers (a,b) to be a 
// Oneful Pair, if a+b+(a⋅b)=111

package main 

import"fmt"

func main(){
	var a, b int 

	fmt.Println("Enter the first number:")
	fmt.Scan(&a)

	fmt.Println("Enter the second number:")
	fmt.Scan(&b)

	if (a+b+(a*b)) == 111 {
		fmt.Println("Yes the number is oneful")
	}else{
		fmt.Println("No the number is not oneful")
	}
}
