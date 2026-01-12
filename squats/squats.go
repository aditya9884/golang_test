//Somu went to the gym today. He decided to do X sets of squats. Each set consists of 
//15 squats. Determine the total number of squats that he did today.

package main 

import "fmt"

func main(){

	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var a int 

		fmt.Println("Enter the set of squats:")
		fmt.Scan(&a)

		z := a * 15
		fmt.Println("The total no of squats:", z)

		t--
	}
}