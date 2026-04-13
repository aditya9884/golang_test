/*
In Chefland, a tennis game involves 4 referees.
Each referee has to point out whether he considers the ball to be inside limits or outside limits.
The ball is considered to be IN if and only if all the referees agree that it was inside limits.
input:
The first line of input will contain a single integer T, denoting the number of test cases.
Each test case consists of a single line of input containing 4 integers 
R1,R2,R3,R4denoting the decision of the respective referees.
Here R can be either 0 or 1 where 0would denote that the referee considered
the ball to be inside limits whereas 1 denotes that they consider it to be outside limits.
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Please enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {

		var r1,r2,r3,r4 int 
		fmt.Println("Please enter the value eithe 0 or 1")
		fmt.Println("Enter r1:")
		fmt.Scan(&r1)

		fmt.Println("Enter r2:")
		fmt.Scan(&r2)

		fmt.Println("Enter r3:")
		fmt.Scan(&r3)

		fmt.Println("Enter r4:")
		fmt.Scan(&r4)

		if r1 == 0 && r2 == 0 && r3 ==0 && r4 == 0 {
			fmt.Println("The ball is within the limit so: IN")
		}else{
			fmt.Println("The ball is not within the limit so: OUT")
		}

		t--
	}
}