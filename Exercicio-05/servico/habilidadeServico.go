package servico

import (
	"github.com/gin-gonic/gin"
	"net/http"
	
	"exercicio05/aplicacao"
)

func Iniciar(){
	router := gin.Default()
	router.GET("/ObterHabilidades", ObterHabilidade)

	router.Run("localhost:8080")
}

func ObterHabilidade(c *gin.Context) {
	// Chame a função ObterHabilidade do pacote aplicacao
	habilidades := aplicacao.ObterHabilidade()

	// Retorne a resposta como JSON
	c.JSON(http.StatusOK, habilidades)
}