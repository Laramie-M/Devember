package main


import (
    "fmt"
    //"io"
    "os"
)

func errorCheck(e error){
	if e != nil{
		panic(e)
	}
}

func main() {
	head := make([]byte,4)
	f, err := os.Open("../tmp/dat")
	defer f.Close()
	defer fmt.Println("Closing File")
	errorCheck(err)

	fmt.Println(head)



}

