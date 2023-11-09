package aplicacao

import (
	"exercicio05/modelo"
	"exercicio05/repositorio"
)

func ObterHabilidade() []modelo.Habilidade {
	habilidades := repositorio.ObterHabilidades()
	return habilidades
}

func InserirHabilidade(habilidade string, descricao string) {
	repositorio.InserirHabilidade(habilidade, descricao)
}

func ExcluirHabilidade(id int) {
	repositorio.ExcluirHabilidade(id)
}

func ObterQuantidadeDeHabilidades() int {
	quantidade := repositorio.ObterQuantidadeDeHabilidades()
	return quantidade
}

func ObterHabilidadePorId(id int) modelo.Habilidade {
	habilidade := repositorio.ObterHabilidadePorId(id)
	return habilidade
}