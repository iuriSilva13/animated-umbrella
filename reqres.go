package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resposta, _ := http.Get("https://reqres.in/api/users/2")
	body, _ := ioutil.ReadAll(resposta.Body)
	fmt.Printf("status da resposta: %s\n", body)
}
