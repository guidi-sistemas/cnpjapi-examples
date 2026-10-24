// Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer .NET 6+. Uso: CNPJAPI_KEY=cnpj_sua_chave dotnet run 00776574000156

using System.Net.Http.Headers;
using System.Text.Json;

string cnpj = args.Length > 0 ? args[0] : "00776574000156";
string apiKey = Environment.GetEnvironmentVariable("CNPJAPI_KEY") ?? "cnpj_sua_chave";

using var http = new HttpClient();
http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);

using var resposta = await http.GetAsync($"https://api.cnpjapi.com.br/{cnpj}");

if (!resposta.IsSuccessStatusCode)
{
    Console.Error.WriteLine($"Falha na consulta: HTTP {(int)resposta.StatusCode}");
    Environment.Exit(1);
}

using var doc = JsonDocument.Parse(await resposta.Content.ReadAsStringAsync());
var raiz = doc.RootElement;
var razao = raiz.GetProperty("RazaoSocial").GetString();
var situacao = raiz.GetProperty("SituacaoCadastral").GetProperty("Descricao").GetString();
Console.WriteLine($"{razao} - {situacao}");

// Sanções e embargos (CGU e IBAMA): vêm em toda consulta nos planos pagos.
if (!raiz.TryGetProperty("sancoes", out var sancoes) || sancoes.ValueKind == JsonValueKind.Null)
{
    Console.WriteLine(raiz.TryGetProperty("sancoes_disponivel_em", out _)
        ? "Sanções e embargos: disponível nos planos pagos"
        : "Sanções e embargos: não foi possível verificar agora");
}
else if (!sancoes.GetProperty("possui_registros").GetBoolean())
{
    Console.WriteLine("Sanções e embargos: nada consta");
}
else
{
    var itens = sancoes.GetProperty("itens");
    Console.WriteLine($"Sanções e embargos: {itens.GetArrayLength()} registro(s)");
    foreach (var item in itens.EnumerateArray())
    {
        Console.WriteLine($"  {item.GetProperty("fonte").GetString()}: {item.GetProperty("tipo").GetString()} (CNPJ {item.GetProperty("cnpj_registrado").GetString()})");
    }
}
