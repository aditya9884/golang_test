/*
Greater Average
You are given 3 numbers A,B, and C.
Determine whether the average of A and B is strictly greater than 
C or not?
The first line of input will contain a single integer T, denoting the number of test cases.
Each test case consists of 3 integers A,B, and C.
*/

package main 

import "fmt"

func main(){
	var t int 

	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {

		var a , b , c float64
		fmt.Println("Enter the first number:")
		fmt.Scan(&a)

		fmt.Println("Enter the second number:")
		fmt.Scan(&b)

		fmt.Println("Enter the third number:")
		fmt.Scan(&c)

		z := (a+b)/2

		if z > c {
			fmt.Println("The average of two number is greater than third numb:",z)
		}else{
			fmt.Println("The average of two number is not greater than two numb",z)
		}

		t--
	}
	
}