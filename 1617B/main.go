package main

import (
	"bufio"
	"fmt"
	"os"
)

//func gcd(n int64) string{
//	return strconv.FormatInt(5,10)
//}
//func solve(input int64) string {

//return gcd(input)	
//}

func main() {

	var inputs []string
	input:= bufio.NewScanner(os.Stdin)
	for input.Scan(){
		inputs=append(inputs,input.Text())
	}
	for _,value := range inputs {
		fmt.Println(value)
}
}
