<?php

// Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer PHP 8.0+ com a extensao cURL.
// Uso: CNPJAPI_KEY=cnpj_sua_chave php consultar-cnpj.php 00776574000156

$cnpj = $argv[1] ?? '00776574000156';
$apiKey = getenv('CNPJAPI_KEY') ?: 'cnpj_sua_chave';

$ch = curl_init("https://api.cnpjapi.com.br/{$cnpj}");
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => ["Authorization: Bearer {$apiKey}"],
    CURLOPT_TIMEOUT => 10,
]);

$body = curl_exec($ch);
$status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($status !== 200) {
    fwrite(STDERR, "Falha na consulta: HTTP {$status}\n");
    exit(1);
}

$empresa = json_decode($body, true);
echo $empresa['RazaoSocial'], ' - ', $empresa['SituacaoCadastral']['Descricao'], PHP_EOL;

// Sanções e embargos (CGU e IBAMA): vêm em toda consulta nos planos pagos.
$sancoes = $empresa['sancoes'] ?? null;
if ($sancoes === null) {
    echo isset($empresa['sancoes_disponivel_em'])
        ? 'Sanções e embargos: disponível nos planos pagos'
        : 'Sanções e embargos: não foi possível verificar agora', PHP_EOL;
} elseif (!$sancoes['possui_registros']) {
    echo 'Sanções e embargos: nada consta', PHP_EOL;
} else {
    echo 'Sanções e embargos: ', count($sancoes['itens']), ' registro(s)', PHP_EOL;
    foreach ($sancoes['itens'] as $item) {
        echo "  {$item['fonte']}: {$item['tipo']} (CNPJ {$item['cnpj_registrado']})", PHP_EOL;
    }
}
