package main
import "fmt"
func main() {
	i := 1
	for i <= 10 {
		if isEven(i) && i%3 ==0 {
			fmt.Println("FizzBuzz")
			}else if isEven(i) {
				fmt.Println("Fizz")
			}else if i%3==0{
				fmt.Println("Buzz")
					}else{
						fmt.Println(i)
						}
		i++
	}
}
func isEven(a int)bool{
	return a%2==0
}
