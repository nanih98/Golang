package main 


import (
	"fmt"
)

func main(){
	for i:=1; i<100;i++ {
		if i % 3 == 0 {
			fmt.Printf("El número %d es divisible entre 3 \n",i)
		}
	}
}

