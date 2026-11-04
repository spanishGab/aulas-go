# Criando um serviodr HTTP pela primeira vez

Para este exercício você vai precisar criar um servidor HTTP do zero. Este servidor será de uma aplicação de  organização financeira.

## Requisitos

O seu servidor deve:

* Definir suas rotas através de um roteador (multiplexer)
* Definir um contexto cancelável
* Aceitar sinais de interrupção do sistema operacional
* Executar shutdown com delicadeza

## Rotas

* `POST /api/sign-up` 
* `POST /api/expenses`
* `GET /api/expenses`
* `GET /api/expenses/{id}`
* `PATCH /api/expenses/{id}`
* `DELETE /api/expenses/{id}`

Por enquanto as rotas não devem executar nenhuma ação, apenas logar o que acontece
