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
