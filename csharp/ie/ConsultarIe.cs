// Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Recurso premium (plano com IE). Requer .NET 6+.
// Uso: CNPJAPI_KEY=cnpj_sua_chave dotnet run 00776574000156 SP
//      (a UF e opcional; sem ela, faz a varredura nacional)

using System.Net.Http.Headers;
using System.Text.Json;

string cnpj = args.Length > 0 ? args[0] : "00776574000156";
string uf = args.Length > 1 ? args[1] : "";
string apiKey = Environment.GetEnvironmentVariable("CNPJAPI_KEY") ?? "cnpj_sua_chave";

string url = $"https://api.cnpjapi.com.br/consulta/ie/{cnpj}";
if (uf.Length > 0)
{
    url += "?uf=" + Uri.EscapeDataString(uf);
}

using var http = new HttpClient();
http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);

using var resposta = await http.GetAsync(url);

if (!resposta.IsSuccessStatusCode)
{
    Console.Error.WriteLine($"Falha na consulta de IE: HTTP {(int)resposta.StatusCode}");
    Environment.Exit(1);
}

using var doc = JsonDocument.Parse(await resposta.Content.ReadAsStringAsync());
var resultados = doc.RootElement.GetProperty("resultados");

if (resultados.GetArrayLength() == 0)
{
    Console.WriteLine("Nenhuma inscricao estadual encontrada.");
    return;
}

foreach (var ie in resultados.EnumerateArray())
{
    var numero = ie.GetProperty("ie").GetString();
    if (string.IsNullOrEmpty(numero)) numero = "(nao-contribuinte)";
    Console.WriteLine($"{ie.GetProperty("uf").GetString()}: {numero} - {ie.GetProperty("situacao").GetString()}");
}
