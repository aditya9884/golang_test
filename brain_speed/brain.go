/* In ChefLand, human brain speed is measured in bits per second (bps).
 Chef has a threshold limit of X
X bits per second above which his calculations are prone to errors.
 If Chef is currently working at Y
Y bits per second, is he prone to errors?
If Chef is prone to errors print YES, otherwise print NO.
input 7 9 , output is yes 
Chef's current brain speed of 9 bps is greater than the threshold of 
7 bps, hence Chef is prone to errors.
*/

package main 

import "fmt"

func main(){
	var x , y int 

	fmt.Println("Enter the speed of the thresshold:")
	fmt.Scan(&x)
	fmt.Println("Enter the speed of the brain:")
	fmt.Scan(&y)

	if y > x {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}