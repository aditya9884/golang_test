/*
A new TV streaming service was recently started in Chefland called the Chef-TV.
A group of N friends in Chefland want to buy Chef-TV subscriptions. We know that 
6 people can share one Chef-TV subscription. Also, the cost of one Chef-TV subscription is 
X rupees. Determine the minimum total cost that the group of 
N friends will incur so that everyone in the group is able to use Chef-TV.
Input Format
The first line contains a single integer 
T — the number of test cases. Then the test cases follow.
The first and only line of each test case contains two integers 
N and X — the size of the group of friends and the cost of one subscription.
Output Format
For each test case, output the minimum total cost that the
group will incur so that everyone in the group is able to use Chef-TV.
*/

package main 

import "fmt"

func main(){
	var t int 
	fmt.Println("Enter the no of test case:")
	fmt.Scan(&t)

	for t > 0 {
		var x, a int
		fmt.Println("Enter the no of person:")
		fmt.Scan(&x)

		fmt.Println("Enter the cost of subscription:")
		fmt.Scan(&a)

		/*
		this is hardcode not recommded used for the limited no of user 

		if 0 < x && x <= 6 {
			y := 1 * a 
			fmt.Println("The price of subscription:",y)
		} else if 6 < x && x <= 12 {
			b := 2 * a 
			fmt.Println("The price of subscription:",b)
		}else if 12 < x && x <= 18 {
			c := 3 * a 
			fmt.Println("The price of subscription:",c)
		}else if 18 < x && x <= 24 {
			d := 4 * a
			fmt.Println("The price of subscription:",d)
		}
		*/

		multiplier := (x + 5) / 6
		result := multiplier * a

		fmt.Println("The cost of subscription:",result)


		t--

	}
}