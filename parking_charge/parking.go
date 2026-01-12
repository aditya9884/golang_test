/*Chef needs to park her car while she watches a movie. 
The parking charges at the theater are as follows:

Rs. X, X for the first 1 hour
Rs. Y, Y for every extra hour after the first hour
If Chef parks her car for 
H, H hours, what is the total parking charges that she should pay?
*/

package main 

import "fmt"

func main() {
	var x , y , h int 

	fmt.Println("Enter the value of x :")
	fmt.Scan(&x)

	fmt.Println("Enter the value of y :")
	fmt.Scan(&y)

	fmt.Println("Enter the value of h :")
	fmt.Scan(&h)

	z := x + y*(h-1)
	fmt.Println("The total amount should be paid:", z)

}