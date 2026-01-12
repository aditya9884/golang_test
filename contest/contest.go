/* Each input file contains of a single line, with three integers N,A and 
B - the number of new users, the number of users who just saw the problem
and didn't make any submission, and the number of users who made a submission 
but could not solve any problem correctly.
*/

package main 

import "fmt"


func main(){
	var n, a, b int

	fmt.Println("Enter the new user :")
	fmt.Scanln(&n)

	fmt.Println("Number of user who just saw the problem and did not submit:")
	fmt.Scanln(&a)

	fmt.Println("Number of user made the submission and not solved:")
	fmt.Scanln(&b)

	z := (n - a)
	fmt.Println("The no user who submited problem:", z)
	c := (z - b) 
	fmt.Println("who able to solve at least one problem:", c)
}