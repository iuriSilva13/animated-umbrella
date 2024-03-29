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
		fmt.Printf("o nome completo do usuario é:%s\n", nomeCompleto)
	}
}

func consultarNomesNoServidorRemoto(endPoint string) []string {
	var listaNomes []string

	req, _ := http.NewRequest("GET", endPoint, nil)
	resposta, err := httpClient.Do(req)

	if err != nil {
		return listaNomes
	}

	if resposta.StatusCode >= 400 {
		return listaNomes
	}

	body, _ := ioutil.ReadAll(resposta.Body)
	var dadosRecebidos reqres_response_users
	_ = json.Unmarshal(body, &dadosRecebidos)

	for _, usuario := range dadosRecebidos.Data {
		nome := obterNomeCompleto("https://reqres.in/api/users/" + fmt.Sprintf("%+v", usuario.Id))
		if len(nome) == 0 {
			return listaNomes
		}
		listaNomes = append(listaNomes, nome)
	}

	return listaNomes
}

func obterNomeCompleto(endPoint string) string {
	var respostaErro string
	req, _ := http.NewRequest("GET", endPoint, nil)
	Resposta, err := httpClient.Do(req)

	if err != nil {
		return respostaErro
	}

	if Resposta.StatusCode >= 400 {
		return respostaErro
	}

	body, _ := ioutil.ReadAll(Resposta.Body)
	var resposta reqres_response_user
	_ = json.Unmarshal(body, &resposta)

	return resposta.DadosUsuario.First_name + " " + resposta.DadosUsuario.Last_name
}
