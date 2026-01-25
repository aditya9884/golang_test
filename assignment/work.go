/*
Janmansh has to submit 3 assignments for Chingari before 
10 pm and he starts to do the assignments at X pm. Each assignment takes him 
1 hour to complete. Can you tell whether he'll be able to complete 
all assignments on time or not
*/

package main 


import "fmt"

func main(){
	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {

		var x int 
		fmt.Println("Please enter the time:")
		fmt.Scan(&x)

		if x <= 7 {
			fmt.Println("Work will complete at 10 p.m")
		}else{
			fmt.Println("Work will not complete at 10 p.m")
		}

		t--
	}
}