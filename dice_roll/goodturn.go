/*
Chef and Chefina are playing with dice. In one turn, both of them roll their dice at once.
They consider a turn to be good if the sum of the numbers on their dice is greater than 6.
Given that in a particular turn Chef and Chefina got X and Y on their respective dice,
find whether the turn was good.
x + y > 6 output is yes 
x + y <= 6 output is no
*/

package main 

import "fmt"


func main(){
	var t int 

	fmt.Println("Please enter the no time dice role:")
	fmt.Scan(&t)

	for i:=1; i <= t; i++ {
		var x , y int 

		
			fmt.Println("Enter the dice role 1st:")
			fmt.Scan(&x)
		
			fmt.Println("Enter the dice role 2nd:")
			fmt.Scan(&y)

			i := x + y 

			if i > 6 {
				fmt.Println("Yes")
			}else{
				fmt.Println("No")
			}
		

	}
}