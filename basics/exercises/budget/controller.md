
## Controlador de budget

Um budget é uma quantidade determinada de dinheiro destinada a uma operação.

O budget deverá ter um valor monetário (incluindo o tipo de moeda em que este valor está) do tipo inteiro que representa seu valor atual; e um valor monetário que representa seu valor máximo.

> utilizamos o tipo inteiro para valores monetários porque ainda não chegamos a falar do pacote decimal, que é a forma correta de armazenar valores monetários. Exemplo: R$ 10,64 == 1064 no seu budget.

O budget também deverá conter um estado, que pode ser um dos seguintes:

* CREATED: budget foi criado mas ainda não pode ser usado
* ACTIVE: budget está ativo e pronto para uso
* PAUSED: budget está pausado e não pode ser usado
* EMPTY: budget está expirado e precisa de uma recarga de valor para ser utilizado
* FINISHED: budget foi encerrado e não pode ser mais usado

Você deve construir o programa de modo que ele possa ser executado por funções de teste. Lembre-se de separar os conceitos. Buget é uma coisa, o estado de um budget é outra, etc.

Abaixo segue o diagrama de máquina de estados de um budget e os componentes de software que você deve declarar previamente com o mesmo nome para que o teste funcione:

* tipo `State`: um tipo customizado que tem como tipo subjacente uma string
* constantes de estado `CreatedState`, `ActiveState`, `PausedState`, `EmptyState`, `FinishedState`: constantes do tipo state que guardam os nomes dos estados em upper case
* tipo Budget: uma struct que deve ter pelo menos os campos `amount: int`, `maxAmount: int`, `state: State`, `currency: string`; e os seguintes métodos:
    * Activate() error: ativa o budget para uso
    * Pause() error: pausa o budget
    * Spend(int) error: retira dinheiro do budget
    * Recharge(int) error: recarrega o budget
    * Finish() error: desabilita o budget para uso
    * String() string: retorna a representação em string do budget; p.ex.: ""Budget amount: R$ 10.00; maxAmount: R$ 50.00; state: ACTIVE""

[Budget state machine](budget_state.excalidraw)
