## Parser de Arquivos

Imagine que você tem dois sistemas externos (um legado e um atual) que seu sistema utiliza para coletar dados sobre carros e utilizá-los em determinado processo.

O sistema **legado-a** retorna dados em `xml` de acordo com o seguinte exemplo:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<cars>
    <car>
        <brand>Honda</brand>
        <model>Prelude</model>
        <year>2026</year>
        <fuel_types>
            <fuel_type>gas</fuel_type>
            <fuel_type>ethanol</fuel_type>
            <fuel_type>electric</fuel_type>
        </fuel_types>
        <price>250.000</price>
        <transmission>automatic</transmission>
    </car>
</cars>
```

Já o sistema **atual-b** retorna dados em `json` de acordo com o seguinte exemplo:

```json
{
    "cars": [
        {
            "brand": "Honda",
            "model": "Prelude",
            "year": 2026,
            "fuel_types": [
                "gas",
                "ethanol",
                "electric"
            ],
            "price": "250.000",
            "transmission": "automatic"
        }
    ]
}
```

Sua tarefa é construir:

* Uma função chamada `ParseCarsData`, que fará o parse dos dados dos sistemas externos
* Uma interface adaptadora que determinará o contrato dos parsers
* Um parser JSON
* Um parser XML
* Uma struct chamada `Cars` que representará os dados internamente


A função `ParseCarsData` tem a seguinte assinatura:
```go
func ParseCarsData(data string, parser DataParser) (Cars, error) {...}
```

> Pattern Adapter: https://refactoring.guru/design-patterns/adapter
> Dica: entender sobre as libs nativas `encoding/json` e `enconding/xml`

## Coloridor de mensagem

### O Problema

Construa uma ferramenta de linha de comando em Go que receba uma mensagem e a imprima no terminal com uma formatação visual (cores) definida pela pessoa usuária via flag. Para isso, você precisará modelar diferentes comportamentos de formatação e integrá-los ao pacote `flag` da linguagem através de uma flag customizada.

### Contexto Técnico (Códigos ANSI)

Para modificar a aparência do texto no terminal, utilizamos **códigos de escape ANSI**. Seguem os principais para este exercício:

- **Vermelho (red)**: `"\033[31m"`
- **Verde (green)**: `"\033[32m"`
- **Amarelo (yellow)**: `"\033[33m"`
- **Azul (blue)**: `"\033[34m"`
- **Rosa / Magenta (magenta)**: `"\033[35m"`
- **Ciano (cyan)**: `"\033[36m"`
- **Reset (volta todos os formatos ao padrão)**: `"\033[0m"`

O padrão de uso é: `[codigo_ANSI] + mensagem + [codigo_reset]`.

---

### Requisitos Obrigatórios

**1. Crie sua própria interface de formatação**  
Defina uma interface (sugestão de nome: `Colorizer`) que represente a capacidade de transformar uma string simples em uma string com códigos ANSI embutidos. Pense em qual método é essencial para essa transformação.

**2. Implemente múltiplos formatos concretos**  
Crie implementações concretas da sua interface para **todas as seis cores listadas acima** (red, green, yellow, blue, magenta e cyan), além de uma implementação para o comportamento **Reset** (que retorna a mensagem sem formatação ou aplica apenas o reset, garantindo que nenhuma cor "vaze" para o restante do terminal).

**3. Construa uma flag customizada**  
O pacote `flag` do Go exige que valores customizados implementem a interface `flag.Value`. Para atender a este requisito, você deverá:
- Declarar um novo tipo que armazene internamente uma das suas implementações da interface criada no item 1.
- Implementar os métodos necessários para que este tipo satisfaça `flag.Value`.
- O método de *configuração* (`Set`) deve receber a string passada pelo usuário (que estará em inglês, ex: `"red"`, `"blue"`, `"cyan"`) e, a partir dela, decidir qual implementação concreta deve ser atribuída ao campo interno da struct. Caso a string não corresponda a nenhuma das cores conhecidas, o método deve retornar um erro (e não apenas ignorar o valor).

**4. A aplicação principal**  
No `main`, a aplicação deve:
- Registrar a flag customizada (ex: `-color`).
- Fazer o parsing dos argumentos.
- Pegar os argumentos restantes como sendo a mensagem a ser exibida (caso nenhum seja passado, use uma mensagem padrão como `"Hello, world!"`).
- Aplicar a formatação escolhida através da interface criada e exibir o resultado no terminal.

**Comportamento esperado:**
```bash
$ go run main.go -color red "Texto em vermelho"
# (exibe o texto em vermelho no terminal)

$ go run main.go -color cyan "Texto em ciano"
# (exibe o texto em ciano)

$ go run main.go -color magenta "Mensagem rosa"
# (exibe o texto em magenta)

$ go run main.go -color yellow "Atenção!"
# (exibe o texto em amarelo)

$ go run main.go -color orange "Teste"
# Deve exibir uma mensagem de erro clara (retornada pelo Set) indicando que a cor é inválida, por exemplo: "invalid color: orange"
```

---

### Desafio Bônus (para aprofundar ainda mais o conceito)

Crie uma nova implementação da sua interface chamada `MultiColorizer` (ou `CompositeColorizer`).

Ela deve ser capaz de **combinar múltiplas formatações** em uma única string, aplicando-as em sequência (por exemplo, primeiro ativar o amarelo, depois ativar o negrito — lembrando que negrito é `"\033[1m"`, então você pode criar a implementação `Bold` para isso).

Para testá-la, modifique a lógica da sua flag customizada para que ela aceite **combinações separadas por vírgula** na linha de comando, como:

```bash
$ go run main.go -color yellow,bold "Alerta!"
```

Aqui, sua flag deve interpretar a string, instanciar cada formatador individual e combiná-los dentro de uma única instância do `MultiColorizer`, que por sua vez também implementa a sua interface `Colorizer`.

---

### Critérios de Avaliação (o que deve ser entregue)

- O código deve compilar sem erros.
- A flag customizada deve ser registrada com `flag.Var`.
- A validação de entradas inválidas deve ser tratada no método `Set` com retorno de erro.
- A separação entre a "lógica de formatação" (interface + structs) e a "lógica de entrada de dados" (flag + main) deve ser clara e bem estruturada.
- O desafio bônus exige pensar em uma solução de composição (um formatador que contém uma lista de outros formatadores) e na ordem correta de aplicação dos códigos ANSI (especialmente onde colocar o `\033[0m` no final).

---

### Dica geral (a única permitida!)

Pense em como os códigos ANSI funcionam na prática: eles são "estados" que o terminal vai acumulando. Se você aplicar vermelho e depois azul, o texto final fica azul porque o azul sobrescreve o vermelho (já que ambos são cores de texto). Porém, se você aplicar negrito e depois vermelho, o texto fica vermelho *e* negrito. Como sua interface pode representar essa composição sem saber quais combinações específicas existirão? Essa é a beleza de projetar com interfaces!

Boa sorte e divirta-se com as cores! 🎨
