/*
Remove Outermost Parentheses
You are given a valid parentheses string s.
A valid parentheses string is composed of '(' and ')' that are properly balanced.
A valid parentheses string is called primitive if it cannot be split into two smaller valid parentheses strings.
Your task is to remove the outermost parentheses from every primitive part of 
s and print the final result.
input : The first line contains an integer T, the number of test cases.
Each test case consists of a single line containing the string s.
output : For each test case, print the modified string after removing the outermost parentheses.
4
((()))
(()(()))
()()
((())())(()(()))
output:
(())
()(())
(())()()(())
*/



package main 

import (
	"fmt"
	"strings"
)

/*
func removeParentheses(s string) string{

	//Check if string has at least 2 characters and starts/ends with parens
	if len(s) >= 2 && s[0] == '(' && s[len(s)-1] == ')'{
		return s[1:len(s)-1]
	}
	return s // return unchanged if no outer parentheses
}
*/


func removeOuterParentheses(s string) string {
	result := strings.Builder{}
	balance := 0

	for _, char := range s {
		if char == '(' {
			if balance > 0 {
				result.WriteRune(char)
			}
			balance++
		} else {
			balance--
			if balance > 0 {
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}

func main() {
	var t int 

	fmt.Println("Enter the no of test case:")
	fmt.Scanln(&t)

	for t > 0 {

		var a string
		fmt.Println("Enter the no of parentheses:")
		fmt.Scanln(&a)

		fmt.Println("The result :",removeOuterParentheses(a))

		t--
	}

	
}