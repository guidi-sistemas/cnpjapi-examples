// Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).
// Java 17+ (o HttpClient e nativo desde o 11). Roda como arquivo unico:
// Exige API key de um plano PAGO (o lote nao esta no plano gratuito).
// Uso: CNPJAPI_KEY=cnpj_sua_chave java ConsultarCnpjLote.java 00776574000156,00000000000000,abc
//
// A resposta e um array JSON (um item por CNPJ); aqui imprimimos o corpo cru.
// Para mapear em objetos, use uma lib JSON (Jackson, Gson).

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class ConsultarCnpjLote {
    public static void main(String[] args) throws Exception {
        String[] cnpjs = args.length > 0
                ? args[0].split(",")
                : new String[] { "00776574000156", "00000000000000", "abc" };
        String apiKey = System.getenv().getOrDefault("CNPJAPI_KEY", "cnpj_sua_chave");

        StringBuilder corpo = new StringBuilder("{\"cnpjs\":[");
        for (int i = 0; i < cnpjs.length; i++) {
            if (i > 0) {
                corpo.append(",");
            }
            corpo.append("\"").append(cnpjs[i]).append("\"");
        }
        corpo.append("]}");

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("https://api.cnpjapi.com.br/consulta/lote"))
                .header("Authorization", "Bearer " + apiKey)
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(corpo.toString()))
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            System.err.println("Falha na consulta: HTTP " + response.statusCode());
            System.exit(1);
        }

        System.out.println(response.body());
    }
}
