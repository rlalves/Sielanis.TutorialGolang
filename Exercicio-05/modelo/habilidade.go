package modelo

type Habilidade struct {
	Id        	int    `db:"id" json:"id"`
	Habilidade	string `db:"habilidade " json:"habilidade"`
	Descricao  	string `db:"descricao " json:"descricao"`
}