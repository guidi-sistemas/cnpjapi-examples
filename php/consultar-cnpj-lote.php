<?php

// Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer PHP 8.0+ com a extensao cURL.
// Exige API key de um plano PAGO (o lote nao esta no plano gratuito).
// Uso: CNPJAPI_KEY=cnpj_sua_chave php consultar-cnpj-lote.php 00776574000156,00000000000000,abc

$cnpjs = isset($argv[1]) ? explode(',', $argv[1]) : ['00776574000156', '00000000000000', 'abc'];
$apiKey = getenv('CNPJAPI_KEY') ?: 'cnpj_sua_chave';

$ch = curl_init('https://api.cnpjapi.com.br/consulta/lote');
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_POSTFIELDS => json_encode(['cnpjs' => $cnpjs]),
    CURLOPT_HTTPHEADER => [
        "Authorization: Bearer {$apiKey}",
        'Content-Type: application/json',
    ],
    CURLOPT_TIMEOUT => 15,
]);

$body = curl_exec($ch);
$status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($status !== 200) {
    fwrite(STDERR, "Falha na consulta: HTTP {$status}\n");
    exit(1);
}

foreach (json_decode($body, true) as $item) {
    if ($item['status'] === 'encontrado') {
        echo $item['cnpj'], ' - ', $item['dados']['RazaoSocial'], PHP_EOL;
    } else {
        echo $item['cnpj'], ' - ', $item['status'], PHP_EOL;
    }
}
