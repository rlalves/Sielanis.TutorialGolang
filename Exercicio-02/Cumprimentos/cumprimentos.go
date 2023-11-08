package Cumprimentos

import (
	"fmt"
	"errors"
)

func Ola(nome string) (string, error){
	if nome == "" {
		return "", errors.New("Nome vazio.")
	}

	var mensagem string
	mensagem = fmt.Sprintf("Olá, %v. Bem vindo!", nome)
	return mensagem, nil
}