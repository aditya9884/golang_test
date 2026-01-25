/*
in Chefland, there are X schools, and each school has Y students.
The year end results are in and a total of Z students passed the exams.
Assuming that all students appeared for the exams, find whether the number 
of students who passed in Chefland was strictly greater than 50%.
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x , y , z int 
		fmt.Println("Enter the no of school:")
		fmt.Scan(&x)

		fmt.Println("Enter the no of student taking the exam:")
		fmt.Scan(&y)

		fmt.Println("Enter the no student passed the exam:")
		fmt.Scan(&z)

		if z > (x * y)/2 {
			fmt.Println("The % of passing student is more than 50%")
		}else{
			fmt.Println("The % of passing student is less than 50%")
		}

		t--
	}
}