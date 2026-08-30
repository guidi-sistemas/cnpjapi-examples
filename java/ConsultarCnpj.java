// Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).
// Java 17+ (o HttpClient e nativo desde o 11). Roda como arquivo unico:
// Uso: CNPJAPI_KEY=cnpj_sua_chave java ConsultarCnpj.java 00776574000156
//
// A resposta e JSON (campos em PascalCase); aqui imprimimos o corpo cru.
// Para mapear em objetos, use uma lib JSON (Jackson, Gson).

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class ConsultarCnpj {
    public static void main(String[] args) throws Exception {
        String cnpj = args.length > 0 ? args[0] : "00776574000156";
        String apiKey = System.getenv().getOrDefault("CNPJAPI_KEY", "cnpj_sua_chave");

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("https://api.cnpjapi.com.br/" + cnpj))
                .header("Authorization", "Bearer " + apiKey)
                .GET()
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            System.err.println("Falha na consulta: HTTP " + response.statusCode());
            System.exit(1);
        }

        System.out.println(response.body());
    }
}
