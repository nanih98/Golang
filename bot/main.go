package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
	fmt.Println("Hello I'm a Bot")
	//fmt.Println("Ahora voy a printar un switch sencillo:")
	option := os.Args[1]
	i, err := strconv.Atoi(option)
	switch i {
	case 64
	: fmt.Println("Has introducido el número 64")
	default: fmt.Println("No entiendo la instrucción")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//fmt.Println(option, i)
}



// package main

// import (
//     "flag"
//     "fmt"
//     "os"
//     "strconv"
// )

// func main() {
//     flag.Parse()
//     s := flag.Arg(0)
//     // string to int
//     i, err := strconv.Atoi(s)
//     if err != nil {
//         // handle error
//         fmt.Println(err)
//         os.Exit(2)
//     }
//     fmt.Println(s, i)
// }