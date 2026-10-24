// Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Uso: CNPJAPI_KEY=cnpj_sua_chave go run main.go 00776574000156
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Empresa struct {
	RazaoSocial       string `json:"RazaoSocial"`
	SituacaoCadastral struct {
		Descricao string `json:"Descricao"`
	} `json:"SituacaoCadastral"`
	// Sanções e embargos (CGU e IBAMA): vêm em toda consulta nos planos pagos. nil no
	// gratuito (sancoes_disponivel_em) ou quando não deu para verificar (sancoes_erro).
	Sancoes *struct {
		PossuiRegistros bool `json:"possui_registros"`
		Itens           []struct {
			Fonte          string `json:"fonte"`
			Tipo           string `json:"tipo"`
			CnpjRegistrado string `json:"cnpj_registrado"`
		} `json:"itens"`
	} `json:"sancoes"`
	SancoesDisponivelEm string `json:"sancoes_disponivel_em"`
}

func main() {
	cnpj := "00776574000156"
	if len(os.Args) > 1 {
		cnpj = os.Args[1]
	}
	apiKey := os.Getenv("CNPJAPI_KEY")
	if apiKey == "" {
		apiKey = "cnpj_sua_chave"
	}

	req, _ := http.NewRequest(http.MethodGet, "https://api.cnpjapi.com.br/"+cnpj, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("falha na consulta: HTTP %d", resp.StatusCode)
	}

	var empresa Empresa
	if err := json.NewDecoder(resp.Body).Decode(&empresa); err != nil {
		log.Fatal(err)
	}

	fmt.Println(empresa.RazaoSocial, "-", empresa.SituacaoCadastral.Descricao)

	switch s := empresa.Sancoes; {
	case s == nil && empresa.SancoesDisponivelEm != "":
		fmt.Println("Sanções e embargos: disponível nos planos pagos")
	case s == nil:
		fmt.Println("Sanções e embargos: não foi possível verificar agora")
	case !s.PossuiRegistros:
		fmt.Println("Sanções e embargos: nada consta")
	default:
		fmt.Printf("Sanções e embargos: %d registro(s)\n", len(s.Itens))
		for _, it := range s.Itens {
			fmt.Printf("  %s: %s (CNPJ %s)\n", it.Fonte, it.Tipo, it.CnpjRegistrado)
		}
	}
}
