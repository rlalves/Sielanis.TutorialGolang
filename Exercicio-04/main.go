package main

import (
	"fmt"
	"exercicio04/aplicacao"
)

func main() {
	aplicacao.InserirHabilidade("Go", "Iniciante")
	aplicacao.InserirHabilidade("Java", "Conhecimento mediano")
	fmt.Println(aplicacao.ObterHabilidade())
}