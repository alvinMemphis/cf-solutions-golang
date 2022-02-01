
package main

import (
	"fmt"
)


func main() {


	var size int 
	
	fmt.Scanf("%d\n",&size)
	outputs:=make([]int ,size)

	inputs:=make([]int ,size)

	for i:=0;i<size;i++{

		fmt.Scanf("%d\n",&inputs[i])

	}

		fmt.Scanf("\n")
	for i:=0;i<size;i++{
		var reminder int
		if inputs[i]%7==0{
			outputs[i]=inputs[i]
		}else{
			
		reminder=inputs[i]%7
		if inputs[i]%10<reminder {
			outputs[i]=inputs[i]-reminder+7
		} else{

		outputs[i]=inputs[i]-reminder
		}
	
			
		}

}
for _,value :=range outputs{
	fmt.Printf("%d\n",value)
}
}
