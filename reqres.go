package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dadosData struct {
	ValoresData dadosUsuario `json:"data"`
}
type dadosUsuario struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}
type respostaUsuarios struct {
	Page        int            `json:"page"`
	Per_page    int            `json:"per_page"`
	Total       int            `json:"total"`
	Total_pages int            `json:"total_pages"`
	Data        []dadosUsuario `json:"data"`
}

func main() {
	ReceberNomes := obterNomes("https://reqres.in/api/users?page=4")
	for _, NomeCompleto := range ReceberNomes {
		fmt.Printf("o nome completo do usuario é:%+v\n", NomeCompleto)
	}
}
func obterNomes(endPoint string) []string {
	var resultado []string
	resposta, _ := http.Get(endPoint)
	body, _ := ioutil.ReadAll(resposta.Body)
	var resp respostaUsuarios
	_ = json.Unmarshal(body, &resp)

	for _, usuario := range resp.Data {
		Idusuarios := "https://reqres.in/api/users/" + fmt.Sprintf("%+v", usuario.Id)
		resultado = obterEndpoint(Idusuarios, resp)
	}

	return resultado

}
func obterEndpoint(listaNomes string, dadosId respostaUsuarios) []string {
	var resultado []string
	respostaEndpoint, _ := http.Get(listaNomes)
	body2, _ := ioutil.ReadAll(respostaEndpoint.Body)
	var dadosRecebios dadosData
	_ = json.Unmarshal(body2, &dadosRecebios)

	for _, usuario2 := range dadosId.Data {
		resultado = append(resultado, usuario2.First_name+" "+usuario2.Last_name)
	}

	return resultado

}
