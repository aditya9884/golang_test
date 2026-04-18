/*
The juicer sells each glass of sugarcane juice for 50 coins.
He spends 20% of his total income on buying sugarcane.
He spends 20% of his total income on buying salt and mint leaves.
He spends 30% of his total income on shop rent.
output
For each test case, output ona new line the juicer's profit when he sells N glasses
of juice.
*/

package main

import "fmt"

func main(){
	var t int 
	fmt.Println("Please enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x , b, c, d, a float64
		fmt.Println("Enter the no of glasses consumed:")
		fmt.Scan(&x)

		a = x * 50 
		b = .20 * a
		c = .20 * a
		d = .30 * a

		fmt.Println("Monery spent of buying sugarcan:",b)
		fmt.Println("Money spent on buying salt and mint:",c)
		fmt.Println("Money spent on shop rent:",d)

		e := a - (b + c + d)
		fmt.Printf("The total no of profit:%.0f\n",e)


		t--
	}
}