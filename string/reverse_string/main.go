// reverse the string  given

package main

import (
	"fmt"
)

func StringReverse(InputString string) (ResultString string) {

	for _, c := range InputString {
		ResultString = string(c) + ResultString
	}
	return
}

func main() {

	var String1 string
	fmt.Println("Enter the string : ")
	fmt.Scan(&String1)
	fmt.Println("The result for", String1, "is :", StringReverse(String1))

}
