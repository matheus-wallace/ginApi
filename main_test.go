package main

import (
	"ginApi/controllers"
	"ginApi/database"
	"ginApi/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Matheus Pereira", CPF: "11111111111", RG: "12345678"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/matheus", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria ser iguais")
	mockDaResposta := `{"API diz:":"E ai matheus, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConectaComOBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
