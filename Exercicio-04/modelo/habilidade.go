package modelo

type Habilidade struct {
	Id        	int    `db:"id"`
	Habilidade	string `db:"habilidade"`
	Descricao  	string `db:"descricao"`
}