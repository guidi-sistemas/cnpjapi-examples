# CNPJAPI - exemplos de integração

Exemplos oficiais e rodáveis de como consultar **CNPJ** pela [**CNPJAPI**](https://cnpjapi.com.br) em várias linguagens. A CNPJAPI é uma **API REST brasileira** para consulta de dados cadastrais de empresas por CNPJ, com base nos dados públicos da Receita Federal.

- **Site:** <https://cnpjapi.com.br>
- **Documentação:** <https://cnpjapi.com.br/docs>
- **Guias por linguagem:** <https://cnpjapi.com.br/guias>

## Como funciona

Uma consulta é uma requisição HTTP `GET`, autenticada por uma **API key** no cabeçalho `Authorization`:

```
GET https://api.cnpjapi.com.br/{cnpj}
Authorization: Bearer cnpj_sua_chave
```

O `{cnpj}` são **apenas os 14 dígitos**, sem pontuação. A resposta é um JSON com os campos em **PascalCase** (`RazaoSocial`, `SituacaoCadastral`, `AtividadePrincipal`, `Endereco`, `QSA`, ...). Veja o [dicionário de campos](https://cnpjapi.com.br/docs/consulta-cnpj).

## Já usa a ReceitaWS?

A CNPJAPI também responde no **formato da ReceitaWS**, campo a campo - basta trocar o host: `GET https://api.cnpjapi.com.br/v1/cnpj/{cnpj}` (ou `?formato=receitaws`). Veja [Compatível com a ReceitaWS](https://cnpjapi.com.br/docs/receitaws-compat).

## API key

Crie sua conta gratuita em <https://app.cnpjapi.com.br>, gere a sua API key e exporte-a numa variável de ambiente antes de rodar os exemplos:

```bash
export CNPJAPI_KEY="cnpj_sua_chave"
```

## Exemplos por linguagem

Cada exemplo aceita um CNPJ como argumento (ou usa um de teste por padrão) e imprime a razão social e a situação cadastral.

| Linguagem | Arquivo | Rodar |
|---|---|---|
| Python | [`python/consultar_cnpj.py`](python/consultar_cnpj.py) | `pip install requests && python python/consultar_cnpj.py 00776574000156` |
| Node.js | [`node/consultar-cnpj.mjs`](node/consultar-cnpj.mjs) | `node node/consultar-cnpj.mjs 00776574000156` |
| C# / .NET | [`csharp/ConsultarCnpj.cs`](csharp/ConsultarCnpj.cs) | `cd csharp && dotnet run 00776574000156` |
| PHP | [`php/consultar-cnpj.php`](php/consultar-cnpj.php) | `php php/consultar-cnpj.php 00776574000156` |
| Go | [`go/main.go`](go/main.go) | `cd go && go run main.go 00776574000156` |
| Java | [`java/ConsultarCnpj.java`](java/ConsultarCnpj.java) | `java java/ConsultarCnpj.java 00776574000156` |

## Tratamento de erros e limites

- `404` - CNPJ não encontrado na base pública.
- `429` - limite por minuto ou cota mensal excedidos; o cabeçalho `Retry-After` diz quantos segundos aguardar. Veja [Limites e planos](https://cnpjapi.com.br/docs/rate-limit).
- Para processar muitos CNPJs, faça **uma chamada por CNPJ** respeitando o rate limit (veja [consultar em lote](https://cnpjapi.com.br/guias/consultar-cnpj-em-lote)).

## Contribuindo

Sugestões e correções são bem-vindas via issue ou pull request. Mantido por [Guidi Sistemas](https://cnpjapi.com.br/empresa), operadora da CNPJAPI.
