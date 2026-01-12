/*
Your task is very simple: given two integers A and 
B, write a program to add these two numbers and output the sum.
input:-
The first line contains an integer T, the total number of test cases.
Then follow T lines, each line contains two integers, 
A and B.
output:-
For each test case, add A and B and display the sum in a new line.
Testcase 1: 
1+2=3. Hence the first output is 3
Testcase 2: 
100+200=300. Hence the second output is 300
*/

package main 

import "fmt"

func main(){
	var t int 

	fmt.Println("Please enter total no of input:")
	fmt.Scan(&t)

	for i:= 1; i <= t; i++{
		var a , b int 

		fmt.Println("Enter the first input no:",)
		fmt.Scan(&a)

		fmt.Println("Enter the 2nd input no:")
		fmt.Scan(&b)

		i := a + b
		fmt.Println("The sum of the input:",i)
	}
}