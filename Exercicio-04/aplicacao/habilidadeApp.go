package aplicacao

import (
	"exercicio04/modelo"
	"exercicio04/repositorio"
)

func ObterHabilidade() []modelo.Habilidade {
	habilidades := repositorio.ObterHabilidades()
	return habilidades
}

func InserirHabilidade(habilidade string, descricao string) {
	repositorio.InserirHabilidade(habilidade, descricao)
}