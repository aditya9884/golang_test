/*
Chef is taking his baby steps into the world of programming.
The very first program he's tasked to write is as follows:
"Given two integers A and B, print A+B."
Unfortunately, Chef makes a typo: his program outputs 
A×B instead of A+B
The correct answer is 4+7=11, but Chef's program prints 4×7=28.
The difference between them is ∣28−11∣=17.
1≤A,B≤10
*/

package main 


import "fmt"


func main(){

	var a , b int 

	fmt.Println("Choose the value between 1 to 10")

	fmt.Println("Enter the first no:")
	fmt.Scan(&a)

	fmt.Println("Enter the second no:")
	fmt.Scan(&b)

	if a >=1 && a <=10 && b >=1 && b <= 10 {
		z := (a * b) - (a + b)

		// to make the negative to positive 
		if z < 0 {
			z = -z
		}
		fmt.Println("The resultant value is:", z)
	}else{
		fmt.Println("The value is out of range")
	}

}