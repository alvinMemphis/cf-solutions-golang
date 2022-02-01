package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func max(array []int) int{
	max:=0
	for index,value:=range array{

		if index==0||  value>max {
			max=value
		}

	}
	return max
	
	
}

func solve(n int,a []int,b []int) int{

			for i := 0; i < n; i++ {
			if b[i] > a[i] {
				tmp := b[i]
				b[i] = a[i]
				a[i] = tmp
			}
		}
sort.Ints(a)
sort.Ints(b)
	return a[n-1]*b[n-1]
}

 func looper(array []string,index int){ 

    n,_:=strconv.Atoi(array[index-2])
    a:=make([]int,n)
    b:=make([]int,n)
    a_string:=strings.Split(array[index-1]," ")
    b_string:=strings.Split(array[index]," ")


    for i,value:=range a_string{
	    int_val,_:= strconv.Atoi(value)
	   a[i]=int_val
    }
    for i,value:=range b_string{
	    int_val,_:= strconv.Atoi(value)
	    b[i]=int_val
    }


    
		fmt.Println(solve(n,a,b))
		
		
}

func main() {



    var string_array []string
	file:= bufio.NewScanner(os.Stdin)
	
	
	

	for file.Scan(){
		
		string_array=append(string_array,file.Text())
		
	}



    loops,_:=strconv.Atoi(string_array[0])
    for i:=0;i<loops;i++{

  looper(string_array,loops+(loops*i))

    }
}
