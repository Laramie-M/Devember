package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text,"\n")
	isPalindrome(text)
}
func reverse(a string) (b string){
	for _, i := range a{
		b = string(i) + b
	}
	fmt.Println(a)
	return b
}
func isPalindrome (a string){
	if a == reverse(a) {
		fmt.Println("Palindrome")
	}else{fmt.Println("Not a Palindrome")}
}