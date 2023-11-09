package repositorio

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"exercicio05/modelo"
)

var db *sql.DB

func AbrirConexao() {
	var erro error
	db, erro = sql.Open("mysql", "root:root@tcp(localhost:3306)/Curriculos")
	ChecagemDeErro(erro)
}

func FechaConexao() {
	db.Close()
}

func ObterHabilidades() []modelo.Habilidade {
	AbrirConexao()

	rows, erro := db.Query("SELECT * FROM Habilidades")
	ChecagemDeErro(erro)

	var habilidades []modelo.Habilidade

	defer rows.Close()
	for rows.Next() {
		var h modelo.Habilidade

		erro = rows.Scan(&h.Id, &h.Habilidade, &h.Descricao)
		ChecagemDeErro(erro)

		habilidades = append(habilidades, h)
	}
	
	FechaConexao()

	return habilidades
}

func InserirHabilidade(habilidade string, descricao string) {
	AbrirConexao()

	stmt, erro := db.Prepare("INSERT INTO Habilidades (Habilidade, Descricao) VALUES (?, ?)")
	ChecagemDeErro(erro)

	stmt.Exec(habilidade, descricao)

	FechaConexao()
}

func ExcluirHabilidade(id int) {
	AbrirConexao()

	stmt, erro := db.Prepare("DELETE FROM Habilidades WHERE id = ?")
	ChecagemDeErro(erro)

	stmt.Exec(id)

	FechaConexao()
}

func ObterQuantidadeDeHabilidades() int {
	AbrirConexao()

	rows, erro := db.Query("SELECT COUNT(*) FROM Habilidades")
	ChecagemDeErro(erro)

	var quantidade int

	defer rows.Close()
	for rows.Next() {
		erro = rows.Scan(&quantidade)
		ChecagemDeErro(erro)
	}

	FechaConexao()

	return quantidade
}

func ObterHabilidadePorId(codigo int) modelo.Habilidade {
	AbrirConexao()

	rows, erro := db.Query("SELECT * FROM Habilidades WHERE Id = ?", codigo)
	ChecagemDeErro(erro)

	var h modelo.Habilidade

	defer rows.Close()
	for rows.Next() {
		erro = rows.Scan(&h.Id, &h.Habilidade, &h.Descricao)
		ChecagemDeErro(erro)
	}

	FechaConexao()

	return h
}

func ChecagemDeErro(erro error) {
	if erro != nil {
		panic(erro.Error())
	}
}