package main


import "fmt"

func main(){
	for i := 1; i<10; i++ {
		if i % 2 == 0 {
			fmt.Printf("El número %d es divisible \n", i)
		}
	}
}