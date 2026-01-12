/*
Recently, Chef visited his doctor. The doctor advised Chef to drink at least 
2000 ml of water each day.
Chef drank X ml of water today. Determine if Chef followed the doctor's advice or not.
*/

package main 

import "fmt"


func main(){

	var t int 

	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var a int 

		fmt.Println("Enter the ml of water drank today:")
		fmt.Scan(&a)

		if a >= 2000 {
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
		t--
	}
}