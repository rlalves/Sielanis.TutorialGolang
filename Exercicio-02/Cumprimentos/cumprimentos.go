package Cumprimentos

import (
	"fmt"
	"errors"
	"math/rand"
)

func Ola(nome string) (string, error){
	if nome == "" {
		return "", errors.New("Nome vazio.")
	}

	var mensagem string
	mensagem = fmt.Sprintf(FormatosDeOla(), nome)
	return mensagem, nil
}

func Olas(nomes []string) (map[string]string, error) {
	mensagens := make(map[string]string)
	for _, nome := range nomes {
		mensagem, erro := Ola(nome)
		if erro != nil {
			return nil, erro
		}
		mensagens[nome] = mensagem
	}

	return mensagens, nil
}

func FormatosDeOla() string {
	formatos := []string{
		"Olá, %v. Bem vindo!",
		"Olá, %v. Bem vindo! Hoje é um dia lindo!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo e ensolarado!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo e ensolarado! Aproveite!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo e ensolarado! Aproveite muito!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo e ensolarado! Aproveite muito a vida!",
		"Olá, %v. Bem vindo! Hoje está um dia lindo e ensolarado! Aproveite muito a vida e a natureza!",
	}

	return formatos[rand.Intn(len(formatos))]
}