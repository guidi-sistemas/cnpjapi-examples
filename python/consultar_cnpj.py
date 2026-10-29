"""Consulta um CNPJ pela API da CNPJAPI (https://cnpjapi.com.br).

Requer: pip install requests
Uso:    CNPJAPI_KEY=cnpj_sua_chave python consultar_cnpj.py 00776574000156
"""

import os
import sys

import requests


def consultar(cnpj: str, api_key: str) -> dict:
    resposta = requests.get(
        f"https://api.cnpjapi.com.br/{cnpj}",
        headers={"Authorization": f"Bearer {api_key}"},
        timeout=10,
    )
    resposta.raise_for_status()
    return resposta.json()


def main() -> None:
    cnpj = sys.argv[1] if len(sys.argv) > 1 else "00776574000156"
    api_key = os.environ.get("CNPJAPI_KEY", "cnpj_sua_chave")

    empresa = consultar(cnpj, api_key)
    print(empresa["RazaoSocial"], "-", empresa["SituacaoCadastral"]["Descricao"])
    print(resumo_sancoes(empresa))


def resumo_sancoes(empresa: dict) -> str:
    """Sanções e embargos (CGU e IBAMA): vêm em toda consulta nos planos pagos."""
    sancoes = empresa.get("Sancoes")
    if sancoes is None:
        if empresa.get("SancoesDisponivelEm"):
            return "Sanções e embargos: disponível nos planos pagos"
        return "Sanções e embargos: não foi possível verificar agora"
    if not sancoes["PossuiRegistros"]:
        return "Sanções e embargos: nada consta"
    linhas = [f"Sanções e embargos: {len(sancoes['Itens'])} registro(s)"]
    for item in sancoes["Itens"]:
        linhas.append(f"  {item['Fonte']}: {item['Tipo']} (CNPJ {item['CNPJRegistrado']})")
    return "\n".join(linhas)


if __name__ == "__main__":
    main()
