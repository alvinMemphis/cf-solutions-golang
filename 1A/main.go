package main

import "fmt"

func solve(input int64) string {

	if input%2 == 0 && input!=2{
		return "YES"
	} else {
		return "No"
	}
}

func main() {

	var input int64
	fmt.Scanf("%d",&input)
fmt.Println(solve(input))
}
