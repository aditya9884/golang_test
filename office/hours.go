/*
Recently Chef joined a new company. In this company, the employees have to work for 
X hours each day from Monday to Thursday. Also, in this company, Friday is
called Chill Day — employees only have to work for Y hours (Y<X) on Friday.
Saturdays and Sundays are holidays.
Determine the total number of working hours in one week.
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x , y int 
		fmt.Println("Enter the no of working hours:")
		fmt.Scan(&x)

		fmt.Println("Enter the no of working hours on friday:")
		fmt.Scan(&y)

		a := (4 * x) + y 
		if y < x {
			fmt.Println("The total no of working hours:",a)
		}else {
			fmt.Println("Friday can not have long working hours.")
		}

		t--
	}
}