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

// Sanções e embargos (CGU e IBAMA): vêm em toda consulta nos planos pagos.
const { Sancoes: sancoes } = empresa;
if (sancoes === null || sancoes === undefined) {
  console.log(
    empresa.SancoesDisponivelEm
      ? "Sanções e embargos: disponível nos planos pagos"
      : "Sanções e embargos: não foi possível verificar agora",
  );
} else if (!sancoes.PossuiRegistros) {
  console.log("Sanções e embargos: nada consta");
} else {
  console.log(`Sanções e embargos: ${sancoes.Itens.length} registro(s)`);
  for (const item of sancoes.Itens) {
    console.log(`  ${item.Fonte}: ${item.Tipo} (CNPJ ${item.CNPJRegistrado})`);
  }
}
