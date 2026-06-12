
## Controlador de budget

Um budget é uma quantidade determinada de dinheiro destinada a uma operação.

Você deverá construir um programa que, dado um valor de entrada, retorne uma struct capaz de executar os seguintes métodos:

* Activate(): ativa o budget para uso
* Pause(): pausa o budget
* Unpause(): despausa o budget
* Spend(int): retira dinheiro do budget
* Finish(): desabilita o budget para uso
* Recharge(int): recarrega o budget
* FullRecharge(): recarrega o budget até seu valor máximo
* Next(): passe o budget para o próximo estado
* Previous(): passe o budget para o estado anterior
* String() string: retorna a representação em string do budget

O budget deverá ter um valor monetário (incluindo o tipo de moeda em que este valor está) do tipo inteiro que representa seu valor atual; e um valor monetário que representa seu valor máximo.

> utilizamos o tipo inteiro para valores monetários porque ainda não chegamos a falar do pacote decimal, que é a forma correta de armazenar valores monetários. Exemplo: R$ 10,64 == 1064 no seu budget.

O budget também deverá conter um estado, que pode ser um dos seguintes:

* CREATED: budget foi criado mas ainda não pode ser usado
* ACTIVE: budget está ativo e pronto para uso
* PAUSED: budget está pausado e não pode ser usado
* EMPTY: budget está expirado e precisa de uma recarga de valor para ser utilizado
* FINISHED: budget foi encerrado e não pode ser mais usado

Você deve construir o programa de modo que ele possa ser executado por funções de teste. Lembre-se de separar os conceitos buget é uma coisa, o estado de um budget é outra, o valor do budget pode ser outra, etc.

[Budget state machine](budget_state.excalidraw)
