package main

import (
	"fmt"
	"log"
	"net/http"
)

type Persona struct {
    Nombre   string
    Apellido string
    Edad     uint8
}

type Perro struct {
	Edad 	uint8
	Color	string
	Raza 	string
}

type Credentials struct {
	username string
	password string
}


func (p Persona) saludar() {
	log.Printf("Hola me llamo %s %s y tengo %d años",p.Nombre,p.Apellido,p.Edad)
}

func (p Perro) describe_perro() {
	log.Printf("Mi perro tiene %d años, es de color %s y es de la raza %s",p.Edad,p.Color,p.Raza)
}

func (c Credentials) get_credentials() {
	fmt.Printf(("Enter your username: "))
	fmt.Scanln(&c.username)
	fmt.Printf("Enter your password: ")
	fmt.Scanln(&c.password)

	fmt.Printf("Your username is %s and your password is %s \n",c.username,c.password)
}


func main () {
	credentials := Credentials{}

	credentials.get_credentials()

	p := Persona{"Daniel","Cascales",23}
	p.saludar()
	perro := Perro{3,"negro","pitbull"}

	perro.describe_perro()

	//HTTP requests
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	fmt.Printf("Status code: %s",resp.Status)

	if err != nil {
	   	log.Fatalln(err)
	}
}