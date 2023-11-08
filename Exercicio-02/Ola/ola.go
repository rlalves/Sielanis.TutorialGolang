package main

import (
	"fmt"
	"log"
	"sielanis/cumprimentos"
)

func main() {
	log.SetPrefix("Log: ")
	log.SetFlags(0)

	nomes := []string{"Rodrigo", "Pedro", "Davi", "Fulano", "Ciclano", "Beltrano"}
	mensagens, erro := Cumprimentos.Olas(nomes)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(mensagens)
}