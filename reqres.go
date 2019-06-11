package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Usuario struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}
type respostaUsuario struct {
	ValoresData Usuario `json:"data"`
}
type respostaUsuarios struct {
	Page        int       `json:"page"`
	Per_page    int       `json:"per_page"`
	Total       int       `json:"total"`
	Total_pages int       `json:"total_pages"`
	Data        []Usuario `json:"data"`
}

func main() {
	dadosRespostaUsuarios := obterId("https://reqres.in/api/users?page=4")
	for _, RespostaNomeCompleto := range dadosRespostaUsuarios {
		fmt.Printf("o nome completo do usuario é:%+v\n", RespostaNomeCompleto)
	}
}
func obterId(endPoint string) []string {
	var ListaNomes []string

	resposta, _ := http.Get(endPoint)
	body, _ := ioutil.ReadAll(resposta.Body)
	var dadosRecebidos respostaUsuarios
	_ = json.Unmarshal(body, &dadosRecebidos)

	for _, usuario := range dadosRecebidos.Data {
		Nome := obterNomeCompleto("https://reqres.in/api/users/" + fmt.Sprintf("%+v", usuario.Id))
		ListaNomes = append(ListaNomes, Nome)
	}

	return ListaNomes
}

func obterNomeCompleto(ListaId string) string {
	respostaEndpoint, _ := http.Get(ListaId)
	body2, _ := ioutil.ReadAll(respostaEndpoint.Body)
	var dadosUsuario respostaUsuario
	_ = json.Unmarshal(body2, &dadosUsuario)

	return dadosUsuario.ValoresData.First_name + " " + dadosUsuario.ValoresData.Last_name
}