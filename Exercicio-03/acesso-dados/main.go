package main


import (
	"fmt"
	"strconv"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main(){

	var erro error
	db, erro = sql.Open("mysql", "root:root@tcp(localhost:3306)/Curriculos")
	ChecagemDeErro(erro)
	fmt.Println("Conectado!")

	rows, erro := db.Query("select id, habilidade, descricao from Habilidades")
	ChecagemDeErro(erro)

	defer rows.Close()
	for rows.Next() {
		var h Habilidade
		erro := rows.Scan(&h.id, &h.habilidade, &h.descricao)
		fmt.Println("id = " + strconv.Itoa(h.id) + ", habilidade = " + h.habilidade + ", descricao = " + h.descricao)
		ChecagemDeErro(erro)
	}
}

func ChecagemDeErro(erro error) {
	if erro != nil {
		panic(erro.Error())
	}
}

type Habilidade struct {
	id        	int    `db:"id" json:"id"`
	habilidade	string `db:"habilidade" json:"habilidade"`
	descricao  	string `db:"descricao" json:"descricao"`
}