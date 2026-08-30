// Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer Node.js 18+ (fetch nativo). Modulo ES (.mjs).
// Uso: CNPJAPI_KEY=cnpj_sua_chave node consultar-cnpj.mjs 00776574000156

const cnpj = process.argv[2] ?? "00776574000156";
const apiKey = process.env.CNPJAPI_KEY ?? "cnpj_sua_chave";

const resposta = await fetch(`https://api.cnpjapi.com.br/${cnpj}`, {
  headers: { Authorization: `Bearer ${apiKey}` },
});

if (!resposta.ok) {
  console.error(`Falha na consulta: HTTP ${resposta.status}`);
  process.exit(1);
}

const empresa = await resposta.json();
console.log(empresa.RazaoSocial, "-", empresa.SituacaoCadastral.Descricao);
