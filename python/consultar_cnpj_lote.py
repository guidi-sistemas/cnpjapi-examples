"""Consulta CNPJs em lote pela API da CNPJAPI (https://cnpjapi.com.br).

Requer: pip install requests
Exige API key de um plano PAGO (o lote não está no plano gratuito).
Uso:    CNPJAPI_KEY=cnpj_sua_chave python consultar_cnpj_lote.py 00776574000156,00000000000000,abc
"""

import os
import sys

import requests


def consultar_lote(cnpjs: list[str], api_key: str) -> list[dict]:
    resposta = requests.post(
        "https://api.cnpjapi.com.br/consulta/lote",
        headers={"Authorization": f"Bearer {api_key}"},
        json={"cnpjs": cnpjs},
        timeout=15,
    )
    resposta.raise_for_status()
    return resposta.json()


def main() -> None:
    padrao = ["00776574000156", "00000000000000", "abc"]
    cnpjs = sys.argv[1].split(",") if len(sys.argv) > 1 else padrao
    api_key = os.environ.get("CNPJAPI_KEY", "cnpj_sua_chave")

    for item in consultar_lote(cnpjs, api_key):
        if item["status"] == "encontrado":
            print(item["cnpj"], "-", item["dados"]["RazaoSocial"])
        else:
            print(item["cnpj"], "-", item["status"])


if __name__ == "__main__":
    main()
