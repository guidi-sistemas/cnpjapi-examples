// Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Recurso premium (plano com IE). Java 17+ (o HttpClient e nativo desde o 11). Arquivo unico:
// Uso: CNPJAPI_KEY=cnpj_sua_chave java ConsultarIe.java 00776574000156 SP
//      (a UF e opcional; sem ela, faz a varredura nacional)
//
// A resposta e JSON com { cnpj, fonte, cache_hit, as_of, resultados:[...] };
// aqui imprimimos o corpo cru. Para mapear em objetos, use uma lib JSON (Jackson, Gson).

import java.net.URI;
import java.net.URLEncoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;

public class ConsultarIe {
    public static void main(String[] args) throws Exception {
        String cnpj = args.length > 0 ? args[0] : "00776574000156";
        String uf = args.length > 1 ? args[1] : "";
        String apiKey = System.getenv().getOrDefault("CNPJAPI_KEY", "cnpj_sua_chave");

        String url = "https://api.cnpjapi.com.br/consulta/ie/" + cnpj;
        if (!uf.isEmpty()) {
            url += "?uf=" + URLEncoder.encode(uf, StandardCharsets.UTF_8);
        }

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(url))
                .header("Authorization", "Bearer " + apiKey)
                .GET()
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            System.err.println("Falha na consulta de IE: HTTP " + response.statusCode());
            System.exit(1);
        }

        System.out.println(response.body());
    }
}
