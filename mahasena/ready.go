/*
Kattapa was known to be a very superstitious person. He believed that a soldier
is "lucky" if the soldier is holding an even number of weapons, and "unlucky" otherwise.
He considered the army as "READY FOR BATTLE" if the count of "lucky" soldiers is strictly
greater than the count of "unlucky" soldiers, and "NOT READY" otherwise.
input : 
The first line of input consists of a single integer N denoting the number of soldiers.
The second line of input consists of N space separated integers A1, A2, ..., AN, where Ai 
denotes the number of weapons that the ith soldier is holding.
output:
Generate one line output saying "READY FOR BATTLE", if the army satisfies the conditions 
that Kattapa requires or "NOT READY" otherwise (quotes for clarity).
*/


package main 

import "fmt"

func main() {

	var n int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&n)

	lucky := 0
	unlucky := 0

	for i:=0; i<n; i++{
		var weapon int 
		fmt.Println("Enter the no of weapon:")
		fmt.Scan(&weapon)

		if weapon%2 == 0{
			lucky++
		}else{
			unlucky++
		}
	}

	if lucky > unlucky{
		fmt.Println("REady for battle")
	}else{
		fmt.Println("not ready")
	}
}