// Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer Node.js 18+ (fetch nativo). Modulo ES (.mjs).
// Exige API key de um plano PAGO (o lote nao esta no plano gratuito).
// Uso: CNPJAPI_KEY=cnpj_sua_chave node consultar-cnpj-lote.mjs 00776574000156,00000000000000,abc

const cnpjs = (process.argv[2] ?? "00776574000156,00000000000000,abc").split(",");
const apiKey = process.env.CNPJAPI_KEY ?? "cnpj_sua_chave";

const resposta = await fetch("https://api.cnpjapi.com.br/consulta/lote", {
  method: "POST",
  headers: {
    Authorization: `Bearer ${apiKey}`,
    "Content-Type": "application/json",
  },
  body: JSON.stringify({ cnpjs }),
});

if (!resposta.ok) {
  console.error(`Falha na consulta: HTTP ${resposta.status}`);
  process.exit(1);
}

const itens = await resposta.json();
for (const item of itens) {
  if (item.status === "encontrado") {
    console.log(item.cnpj, "-", item.dados.RazaoSocial);
  } else {
    console.log(item.cnpj, "-", item.status);
  }
}
