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

## Consultar em lote (plano pago)

Tem um plano pago? `POST /consulta/lote` consulta até **20 CNPJs numa só chamada** (conta como 1 requisição pro rate limit). Cada exemplo abaixo aceita uma lista de CNPJs separada por vírgula (ou usa 3 de teste por padrão, ilustrando os status `encontrado`/`nao_encontrado`/`invalido`). Veja a [referência completa](https://cnpjapi.com.br/docs/consulta-lote).

| Linguagem | Arquivo | Rodar |
|---|---|---|
| Python | [`python/consultar_cnpj_lote.py`](python/consultar_cnpj_lote.py) | `python python/consultar_cnpj_lote.py 00776574000156,00000000000000,abc` |
| Node.js | [`node/consultar-cnpj-lote.mjs`](node/consultar-cnpj-lote.mjs) | `node node/consultar-cnpj-lote.mjs 00776574000156,00000000000000,abc` |
| C# / .NET | [`csharp/lote/ConsultarCnpjLote.cs`](csharp/lote/ConsultarCnpjLote.cs) | `cd csharp/lote && dotnet run 00776574000156,00000000000000,abc` |
| PHP | [`php/consultar-cnpj-lote.php`](php/consultar-cnpj-lote.php) | `php php/consultar-cnpj-lote.php 00776574000156,00000000000000,abc` |
| Go | [`go/lote/main.go`](go/lote/main.go) | `cd go/lote && go run main.go 00776574000156,00000000000000,abc` |
| Java | [`java/ConsultarCnpjLote.java`](java/ConsultarCnpjLote.java) | `java java/ConsultarCnpjLote.java 00776574000156,00000000000000,abc` |

## Consultar a Inscrição Estadual (IE) (plano premium)

Tem um plano com **Inscrição Estadual**? `GET /consulta/ie/{cnpj}` retorna a IE de um CNPJ na **fonte oficial** da SEFAZ. Passe a **UF** como 2º argumento (1 crédito) ou omita para a varredura nacional (3 créditos). A resposta é `{ "cnpj": "...", "resultados": [ { "uf": "SP", "ie": "...", "situacao": "habilitado", ... } ] }`. Veja a [referência completa](https://cnpjapi.com.br/docs/consulta-ie).

| Linguagem | Arquivo | Rodar |
|---|---|---|
| Python | [`python/consultar_ie.py`](python/consultar_ie.py) | `python python/consultar_ie.py 00776574000156 SP` |
| Node.js | [`node/consultar-ie.mjs`](node/consultar-ie.mjs) | `node node/consultar-ie.mjs 00776574000156 SP` |
| C# / .NET | [`csharp/ie/ConsultarIe.cs`](csharp/ie/ConsultarIe.cs) | `cd csharp/ie && dotnet run 00776574000156 SP` |
| PHP | [`php/consultar-ie.php`](php/consultar-ie.php) | `php php/consultar-ie.php 00776574000156 SP` |
| Go | [`go/ie/main.go`](go/ie/main.go) | `cd go/ie && go run main.go 00776574000156 SP` |
| Java | [`java/ConsultarIe.java`](java/ConsultarIe.java) | `java java/ConsultarIe.java 00776574000156 SP` |

## Tratamento de erros e limites

- `403` - recurso premium (lote ou Inscrição Estadual) num plano que não o inclui.
- `404` - CNPJ não encontrado na base pública.
- `422` - corpo do lote inválido (`cnpjs` ausente/vazio, ou mais de 20 itens), ou, na IE, CNPJ inválido / UF fora da cobertura.
- `429` - limite por minuto ou cota mensal excedidos; o cabeçalho `Retry-After` diz quantos segundos aguardar. Veja [Limites e planos](https://cnpjapi.com.br/docs/rate-limit).
- Para processar muitos CNPJs, use o [lote](#consultar-em-lote-plano-pago) (plano pago, até 20 por chamada) ou, sem plano pago, faça uma chamada por CNPJ respeitando o rate limit.

## Contribuindo

Sugestões e correções são bem-vindas via issue ou pull request. Mantido por [Guidi Sistemas](https://cnpjapi.com.br/empresa), operadora da CNPJAPI.
