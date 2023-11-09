package servico

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exercicio05/aplicacao"
	"exercicio05/modelo"
)

func Iniciar(){
	router := gin.Default()
	router.GET("/ObterHabilidades", ObterHabilidade)
	router.POST("/ObterHabilidadePorId", ObterHabilidadePorId)
	router.POST("/InserirHabilidadeHtml", InserirHabilidadeHtml)
	router.POST("/InserirHabilidadeJson", InserirHabilidadeJson)
	router.DELETE("/ExcluirHabilidade", ExcluirHabilidade)

	router.Run("localhost:8080")
}

func ObterHabilidade(c *gin.Context) {
	// Chame a função ObterHabilidade do pacote aplicacao
	habilidades := aplicacao.ObterHabilidade()

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, habilidades)
}

func ObterHabilidadePorId(c *gin.Context) {
	
	var request modelo.HabilidadeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao decodificar o JSON"})
        return
    }

	// Obtenha os dados enviados pelo cliente via json
    id := request.Id
    
	// Chame a função ObterHabilidadePorId do pacote aplicacao
	habilidade := aplicacao.ObterHabilidadePorId(id)

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, habilidade)
}

func InserirHabilidadeHtml(c *gin.Context) {
	// Obtenha os dados enviados pelo cliente via formulario
	habilidade := c.PostForm("habilidade")
	descricao := c.PostForm("descricao")

	// Chame a função InserirHabilidade do pacote aplicacao
	aplicacao.InserirHabilidade(habilidade, descricao)

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, gin.H{"mensagem": "Habilidade " + habilidade + " inserida com sucesso!"})
}

func InserirHabilidadeJson(c *gin.Context) {

	var request modelo.HabilidadeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao decodificar o JSON"})
        return
    }

	// Obtenha os dados enviados pelo cliente via json
    habilidade := request.Habilidade
    descricao := request.Descricao

	// Chame a função InserirHabilidade do pacote aplicacao
	aplicacao.InserirHabilidade(habilidade, descricao)

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, gin.H{"mensagem": "Habilidade " + habilidade + " inserida com sucesso!"})
}

func ExcluirHabilidade(c *gin.Context) {
	var request modelo.HabilidadeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao decodificar o JSON"})
        return
    }

	// Obtenha os dados enviados pelo cliente via json
    id := request.Id
    
	// Chame a função ObterHabilidadePorId do pacote aplicacao
	aplicacao.ExcluirHabilidade(id)

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, gin.H{"mensagem": "Habilidade excluida com sucesso!"})
}