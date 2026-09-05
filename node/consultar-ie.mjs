// Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Recurso premium (plano com IE). Requer Node.js 18+ (fetch nativo). Modulo ES (.mjs).
// Uso: CNPJAPI_KEY=cnpj_sua_chave node consultar-ie.mjs 00776574000156 SP
//      (a UF e opcional; sem ela, faz a varredura nacional)

const cnpj = process.argv[2] ?? "00776574000156";
const uf = process.argv[3] ?? "";
const apiKey = process.env.CNPJAPI_KEY ?? "cnpj_sua_chave";

const url = new URL(`https://api.cnpjapi.com.br/consulta/ie/${cnpj}`);
if (uf) url.searchParams.set("uf", uf);

const resposta = await fetch(url, {
  headers: { Authorization: `Bearer ${apiKey}` },
});

if (!resposta.ok) {
  console.error(`Falha na consulta de IE: HTTP ${resposta.status}`);
  process.exit(1);
}

const { resultados } = await resposta.json();
if (resultados.length === 0) {
  console.log("Nenhuma inscricao estadual encontrada.");
} else {
  for (const ie of resultados) {
    console.log(`${ie.uf}: ${ie.ie || "(nao-contribuinte)"} - ${ie.situacao}`);
  }
}
