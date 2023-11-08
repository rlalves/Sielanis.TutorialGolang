package Cumprimentos

import "fmt"

func Ola(nome string) string{
	var mensagem string
	mensagem = fmt.Sprintf("Olá, %v. Bem vindo!", nome)
	return mensagem
}