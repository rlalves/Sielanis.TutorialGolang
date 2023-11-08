package Cumprimentos

import (
	"testing"
	"regexp"
)

func TestOlaNome(t *testing.T) {
	nome := "Rodrigo"
	esperado := regexp.MustCompile(`\b`+nome+`\b`)
	mensagem, erro := Ola(nome)

	if !esperado.MatchString(mensagem) || erro != nil {
		t.Fatalf(`Ola(nome) = %q, %v, esperado math para %#q, nil`, mensagem, erro, esperado)
	}
}

func TestOlaVazio(t *testing.T) {
	mensagem, erro := Ola("")

	if mensagem != "" || erro == nil {
		t.Fatalf(`Ola("") = %q, %v, esperado "", erro `, mensagem, erro)
	}
}