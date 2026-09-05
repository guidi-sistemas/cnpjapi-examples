<?php

// Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Recurso premium (plano com IE). Requer PHP 8.0+ com a extensao cURL.
// Uso: CNPJAPI_KEY=cnpj_sua_chave php consultar-ie.php 00776574000156 SP
//      (a UF e opcional; sem ela, faz a varredura nacional)

$cnpj = $argv[1] ?? '00776574000156';
$uf = $argv[2] ?? '';
$apiKey = getenv('CNPJAPI_KEY') ?: 'cnpj_sua_chave';

$url = "https://api.cnpjapi.com.br/consulta/ie/{$cnpj}";
if ($uf !== '') {
    $url .= '?uf=' . urlencode($uf);
}

$ch = curl_init($url);
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => ["Authorization: Bearer {$apiKey}"],
    CURLOPT_TIMEOUT => 15,
]);

$body = curl_exec($ch);
$status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($status !== 200) {
    fwrite(STDERR, "Falha na consulta de IE: HTTP {$status}\n");
    exit(1);
}

$dados = json_decode($body, true);
if (empty($dados['resultados'])) {
    echo 'Nenhuma inscricao estadual encontrada.', PHP_EOL;
} else {
    foreach ($dados['resultados'] as $ie) {
        $numero = $ie['ie'] !== '' ? $ie['ie'] : '(nao-contribuinte)';
        echo "{$ie['uf']}: {$numero} - {$ie['situacao']}", PHP_EOL;
    }
}
