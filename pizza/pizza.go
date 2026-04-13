/*
Each pizza consists of 4 slices. There are N friends and each friend needs exactly 
X slices.
Find the minimum number of pizzas they should order to satisfy their appetite.
Input Format
Each test case consists of two integers N and X, the number of friends and the number of slices each friend wants respectively.
Output Format
For each test case, output the minimum number of pizzas required.
*/

package main 

import(
	"fmt"
	
)


func main() {
	var t int 
	fmt.Println("Please enter the no test case:")
	fmt.Scan(&t)

	for t > 0 {

		var n, x float64 
		fmt.Println("Please enter the no of friend:")
		fmt.Scan(&n)

		fmt.Println("Please enter the no of slices:")
		fmt.Scan(&x)

		a := n * x 
		// this is main trick in the code
		b := (a+3)/4 

		fmt.Println("The no of pizza required :",b)

		t--
	}

}