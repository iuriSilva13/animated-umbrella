package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type reqres_data_user struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
}
type reqres_response_user struct {
	DadosUsuario reqres_data_user `json:"data"`
}
type reqres_response_users struct {
	Page        int                `json:"page"`
	Per_page    int                `json:"per_page"`
	Total       int                `json:"total"`
	Total_pages int                `json:"total_pages"`
	Data        []reqres_data_user `json:"data"`
}

var httpClient = &http.Client{Timeout: time.Second * 2}

func main() {
	listaNomes := consultarNomesNoServidorRemoto("https://reqres.in/api/users?page=4")
	if len(listaNomes) == 0 {
		fmt.Println("Não recebi nenhuma resposta do servidor remoto")
		return
	}
	for _, nomeCompleto := range listaNomes {
		fmt.Printf("o nome completo do usuario é:%+v\n", nomeCompleto)
	}
}

func consultarNomesNoServidorRemoto(endPoint string) []string {
	var ListaNomes []string
	req, _ := http.NewRequest("GET", endPoint, nil)
	Resposta, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("%+v\n", err)
		return ListaNomes
	}

	if Resposta.StatusCode >= 400 {
		fmt.Println("Resposta não chegou")
		return ListaNomes
	}

	body, _ := ioutil.ReadAll(Resposta.Body)
	var dadosRecebidos reqres_response_users
	_ = json.Unmarshal(body, &dadosRecebidos)

	var BuscaNomes chan string = make(chan string)

	for _, usuario := range dadosRecebidos.Data {
		go obterNomeCompleto(BuscaNomes, "https://reqres.in/api/users/"+fmt.Sprintf("%+v", usuario.Id))
	}

	for i := 1; i <= len(dadosRecebidos.Data); i++ {
		Nome := <-BuscaNomes
		if len(Nome) > 0 {
			ListaNomes = append(ListaNomes, Nome)
		}
	}

	return ListaNomes
}

func obterNomeCompleto(channelDeResultado chan string, endPoint string) {
	var Nome string = ""
	req, _ := http.NewRequest("GET", endPoint, nil)
	Resposta, err := httpClient.Do(req)

	if err != nil {
		channelDeResultado <- Nome
		return
	}

	if Resposta.StatusCode >= 400 {
		channelDeResultado <- Nome
		return
	}

	body2, _ := ioutil.ReadAll(Resposta.Body)
	var resposta reqres_response_user
	_ = json.Unmarshal(body2, &resposta)

	Nome = resposta.DadosUsuario.First_name + " " + resposta.DadosUsuario.Last_name
	channelDeResultado <- Nome
}
