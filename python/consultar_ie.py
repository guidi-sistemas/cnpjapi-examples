"""Consulta a Inscrição Estadual (IE) de um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).

Recurso premium (plano com IE). Requer: pip install requests
Uso:    CNPJAPI_KEY=cnpj_sua_chave python consultar_ie.py 00776574000156 SP
        (a UF e opcional; sem ela, faz a varredura nacional)
"""

import os
import sys

import requests


def consultar_ie(cnpj: str, api_key: str, uf: str = "") -> dict:
    params = {"uf": uf} if uf else {}
    resposta = requests.get(
        f"https://api.cnpjapi.com.br/consulta/ie/{cnpj}",
        params=params,
        headers={"Authorization": f"Bearer {api_key}"},
        timeout=15,
    )
    resposta.raise_for_status()
    return resposta.json()


def main() -> None:
    cnpj = sys.argv[1] if len(sys.argv) > 1 else "00776574000156"
    uf = sys.argv[2] if len(sys.argv) > 2 else ""
    api_key = os.environ.get("CNPJAPI_KEY", "cnpj_sua_chave")

    dados = consultar_ie(cnpj, api_key, uf)
    if not dados["resultados"]:
        print("Nenhuma inscricao estadual encontrada.")
        return
    for ie in dados["resultados"]:
        numero = ie["ie"] or "(nao-contribuinte)"
        print(f"{ie['uf']}: {numero} - {ie['situacao']}")


if __name__ == "__main__":
    main()
