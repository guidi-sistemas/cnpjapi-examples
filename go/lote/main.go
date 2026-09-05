// Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).
// Exige API key de um plano PAGO (o lote nao esta no plano gratuito).
// Uso: CNPJAPI_KEY=cnpj_sua_chave go run main.go 00776574000156,00000000000000,abc
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type corpoLote struct {
	Cnpjs []string `json:"cnpjs"`
}

type itemLote struct {
	Cnpj   string `json:"cnpj"`
	Status string `json:"status"`
	Dados  *struct {
		RazaoSocial string `json:"RazaoSocial"`
	} `json:"dados"`
}

func main() {
	cnpjs := []string{"00776574000156", "00000000000000", "abc"}
	if len(os.Args) > 1 {
		cnpjs = strings.Split(os.Args[1], ",")
	}
	apiKey := os.Getenv("CNPJAPI_KEY")
	if apiKey == "" {
		apiKey = "cnpj_sua_chave"
	}

	corpo, _ := json.Marshal(corpoLote{Cnpjs: cnpjs})
	req, _ := http.NewRequest(http.MethodPost, "https://api.cnpjapi.com.br/consulta/lote", bytes.NewReader(corpo))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("falha na consulta: HTTP %d", resp.StatusCode)
	}

	var itens []itemLote
	if err := json.NewDecoder(resp.Body).Decode(&itens); err != nil {
		log.Fatal(err)
	}

	for _, item := range itens {
		if item.Status == "encontrado" && item.Dados != nil {
			fmt.Println(item.Cnpj, "-", item.Dados.RazaoSocial)
		} else {
			fmt.Println(item.Cnpj, "-", item.Status)
		}
	}
}
