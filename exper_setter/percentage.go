/*
A problem setter is called an expert if at least 50% of their problems are approved by Chef.
Munchy submitted X problems for approval. If Y problems out of those were approved,
find whether Munchy is an expert or not.
*/

package main 

import"fmt"

func main(){
	var t int 
	fmt.Println("Please enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {

		var x , y float64
		fmt.Println("Enter the total no of problem:")
		fmt.Scan(&x)

		fmt.Println("Enter the no of solution done:")
		fmt.Scan(&y)

		c := (y/x)*100

		if c > 50 || c == 50 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}

		t--
	}
}