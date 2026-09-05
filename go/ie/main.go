// Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Recurso premium (plano com IE).
// Uso: CNPJAPI_KEY=cnpj_sua_chave go run main.go 00776574000156 SP
//      (a UF e opcional; sem ela, faz a varredura nacional)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type RespostaIE struct {
	Resultados []struct {
		UF       string `json:"uf"`
		IE       string `json:"ie"`
		Situacao string `json:"situacao"`
	} `json:"resultados"`
}

func main() {
	cnpj := "00776574000156"
	if len(os.Args) > 1 {
		cnpj = os.Args[1]
	}
	uf := ""
	if len(os.Args) > 2 {
		uf = os.Args[2]
	}
	apiKey := os.Getenv("CNPJAPI_KEY")
	if apiKey == "" {
		apiKey = "cnpj_sua_chave"
	}

	endpoint := "https://api.cnpjapi.com.br/consulta/ie/" + cnpj
	if uf != "" {
		endpoint += "?uf=" + url.QueryEscape(uf)
	}

	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("falha na consulta de IE: HTTP %d", resp.StatusCode)
	}

	var dados RespostaIE
	if err := json.NewDecoder(resp.Body).Decode(&dados); err != nil {
		log.Fatal(err)
	}

	if len(dados.Resultados) == 0 {
		fmt.Println("Nenhuma inscricao estadual encontrada.")
		return
	}
	for _, ie := range dados.Resultados {
		numero := ie.IE
		if numero == "" {
			numero = "(nao-contribuinte)"
		}
		fmt.Printf("%s: %s - %s\n", ie.UF, numero, ie.Situacao)
	}
}
