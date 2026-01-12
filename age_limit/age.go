/*
First line will contain T, number of test cases. Then the test cases follow.
Each test case consists of a single line of input, containing three integers 
X,Y, and A as mentioned in the statement.
output: For each test case, output YES if Chef is eligible to give the exam, NO otherwise.
1≤T≤1000
20≤X<Y≤40
10≤A≤50
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Please enter no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x, y, a int 
		
		
			fmt.Println("Enter the  lower limit:")
			fmt.Scan(&x)
			fmt.Println("Enter the upper limit:")
			fmt.Scan(&y)
		
			fmt.Println("Enter the current age:")
			fmt.Scan(&a)

			if 20 <= x && x <= 40 && 20 <=y && y <= 40 && a >= x && a < y {
				fmt.Println("Yes")
			}else {
				fmt.Println("No")
			}

		t--

	}
}