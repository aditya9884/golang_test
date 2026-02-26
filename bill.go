/*
For each bill you pay using CRED, you earn X CRED coins.
At CodeChef store, each bag is worth 100 CRED coins.
Chef pays Y number of bills using CRED. Find the maximum number of bags
he can get from the CodeChef store.
Input Format
First line will contain T, number of test cases. Then the test cases follow.
Each test case contains of a single line of input, two integers X and Y.
Output Format
For each test case, output in a single line - the maximum number of bags
Chef can get from the CodeChef store.
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x, y int 
		fmt.Println("Enter the cred coins:")
		fmt.Scan(&x)

		fmt.Println("The bill pay using CRED:")
		fmt.Scan(&y)

		result := x * y 

		c := result/100
		fmt.Println(c)
		
		t--
	}
}