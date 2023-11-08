package repositorio

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"exercicio04/modelo"
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

func ChecagemDeErro(erro error) {
	if erro != nil {
		panic(erro.Error())
	}
}