package modelo

type HabilidadeRequest struct {
    Id         int    `json:"id"`
    Habilidade string `json:"habilidade"`
    Descricao  string `json:"descricao"`
}