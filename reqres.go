package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Usuario struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
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

	httpClient := &http.Client{Timeout: time.Second * 2}
	req, _ := http.NewRequest("GET", endPoint, nil)
	Resposta, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("%+v\n", err)
		return ListaNomes
	}

	if Resposta.StatusCode >= 400 {
		fmt.Println("Resposta nao chegou")
		return ListaNomes
	}

	body, _ := ioutil.ReadAll(Resposta.Body)
	var dadosRecebidos respostaUsuarios
	_ = json.Unmarshal(body, &dadosRecebidos)

	for _, usuario := range dadosRecebidos.Data {
		Nome := obterNomeCompleto("https://reqres.in/api/users/" + fmt.Sprintf("%+v", usuario.Id))
		ListaNomes = append(ListaNomes, Nome)
	}

	return ListaNomes
}

func obterNomeCompleto(ListaId string) string {
	httpClient := &http.Client{Timeout: time.Second * 2}
	req, _ := http.NewRequest("GET", ListaId, nil)
	Resposta, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("%+v\n", err)
		return ListaId
	}

	if Resposta.StatusCode >= 400 {
		fmt.Println("Resposta nao chegou")
		return ListaId
	}

	body2, _ := ioutil.ReadAll(Resposta.Body)
	var dadosUsuario respostaUsuario
	_ = json.Unmarshal(body2, &dadosUsuario)

	return dadosUsuario.ValoresData.First_name + " " + dadosUsuario.ValoresData.Last_name
}
