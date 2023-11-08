package main

import (
	"fmt"
	"log"
	"sielanis/cumprimentos"
)

func main() {
	log.SetPrefix("Log: ")
	log.SetFlags(0)

	mensagem, erro := Cumprimentos.Ola("Rodrigo")
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(mensagem)
}