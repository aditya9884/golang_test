/* Alice has scored X marks in her test and Bob has scored 
Y marks in the same test. Alice is happy if she scored at least twice the marks of Bob’s score. 
Determine whether she is happy or not.
input : Alice has scored X=2 marks whereas Bob has scored 
Y=1 mark. As Alice has scored twice as much as Bob (i.e.X≥2Y), the answer is Yes.
*/

package main 

import "fmt"

func main() {
	var x, y int 

	fmt.Println("Enter the mark of alice:")
	fmt.Scan(&x)

	fmt.Println("Enter the mark of the Bob's:")
	fmt.Scan(&y)

	if x >= 2*y {
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}