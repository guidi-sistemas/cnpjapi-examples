// Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).
// Requer .NET 6+. Exige API key de um plano PAGO (o lote nao esta no plano gratuito).
// Uso: CNPJAPI_KEY=cnpj_sua_chave dotnet run 00776574000156,00000000000000,abc

using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text.Json;

string[] cnpjs = args.Length > 0
    ? args[0].Split(',')
    : new[] { "00776574000156", "00000000000000", "abc" };
string apiKey = Environment.GetEnvironmentVariable("CNPJAPI_KEY") ?? "cnpj_sua_chave";

using var http = new HttpClient();
http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);

using var resposta = await http.PostAsJsonAsync("https://api.cnpjapi.com.br/consulta/lote", new { cnpjs });

if (!resposta.IsSuccessStatusCode)
{
    Console.Error.WriteLine($"Falha na consulta: HTTP {(int)resposta.StatusCode}");
    Environment.Exit(1);
}

using var doc = JsonDocument.Parse(await resposta.Content.ReadAsStringAsync());
foreach (var item in doc.RootElement.EnumerateArray())
{
    var cnpj = item.GetProperty("cnpj").GetString();
    var status = item.GetProperty("status").GetString();
    if (status == "encontrado")
    {
        var razao = item.GetProperty("dados").GetProperty("RazaoSocial").GetString();
        Console.WriteLine($"{cnpj} - {razao}");
    }
    else
    {
        Console.WriteLine($"{cnpj} - {status}");
    }
}
