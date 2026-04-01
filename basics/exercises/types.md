## Avaliador de despesas

Suponha que você precisa construir a seguinte feature de um sistema de finanças pessoais:

O usuário do sistema precisa poder verificar quais de suas categorias de despesas registradas está dentro do limite percentual máximo definido por ele no sistema.

### Exemplo

Carlos definiu as seguintes metas:

|   Categoria    | Limite Percentual |
| -------------- | ----------------- |
|  Custos fixos  |        25%        |
|  Lazer         |        10%        |
|  Educação      |        5%         |
|  Investimentos |        25%        |
|  Transporte    |        15%        |
|  Alimentação   |        20%        |
|  Total         |        100%       |

Ele ganha R$ 10.000,00 por mês e teve o seguinte total de despesas registradas:

|   Categoria    | Valor da despesa  |
| -------------- | ----------------- |
|  Custos fixos  |     R$ 2.500,00   |
|  Lazer         |     R$ 1.000,00   |
|  Educação      |     R$ 500,00     |
|  Investimentos |     R$ 2.500,00   |
|  Transporte    |     R$ 1.500,00   |
|  Alimentação   |     R$ 2.000,00   |
|  Total         |     R$ 10.000,00  |

Neste cenário, Carlos está dentro do limite estabelecido para todas as despesas. Porém, se ele, por exemplo, tivesse gastado R$ 1.500,00 reais em **Lazer**, ele estaria fora do limite máximo estabelecido para esta categoria. Caso não tivesse gastado nada, nesta ou em qualquer outra categoria, ele estaria dentro do limite máximo estabelecido.

### Feature

Dado este cenário, desenvolva um programa que receba como entrada três valores:

* A soma de receitas do usuário - valor numérico
* As metas definidas pelo usuário - mapa onde as chaves são os nomes das categorias e os valores são as porcentagens (não é necessário verificar se a soma das porcentagens é igual a 100%)
* As despesas do usuário até aquele momento - slice de strings no formato CSV respeitando a seguinte ordem: `categoria,descricao,valor`

Como saída, o programa deve retornar um mapa onde a as chaves são os nomes de cada categoria e os valores são uma das seguintes opções: ("Dentro do limite", "Estourou o limite")

> Bônus: Imprima o resultado em formato de tabela no terminal, com os espaçamentos adequados.
