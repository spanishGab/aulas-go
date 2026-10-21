# Error handler

Você está programando um servidor HTTP e precisa elaborar uma função que irá tratar todos os erros que acontecerem dentro do servidor e retornar uma resposta com base nas regras de tratamento de erro determinadas.
Crie uma função chamada `handleError` que retorna uma resposta com mensagem e código HTTP com a seguinte assinatura:

```go
type Response struct {
    StatusCode int `json:"status_code"`
    Message string `json:"message"`
}

func handleError(err error) Response {}
```

## Contexto sobre erros

Em um servidor HTTP é importante que **nenhum erro** escape ao tratamento, e o ideal é que o tipo de erro que chegua na função de tratamento de erros seja sempre um tipo conhecido e determinado pelo próprio servidor, para que haja maior controle sobre o que fazer com aquele erro.

Por exemplo, se o erro que chega à função de tratamento for do tipo `error` simplesmente, só o que sabemos é que ele possui o método `Error() string` e que podemos saber qual é a mensagem de erro. Porém, se determinarmos um tipo `HTTPError` em nosso servidor, e esse tipo contiver métodos como `Code() int`, `Kind() string`, `Unwrap() error` e `Error() string`; desta forma teremos muito mais controle e acesso ao tipo de erro e o contexto em que ele ocorreu. Além de sabermos em qual nível de detalhamento o usuário final pode saber dos detalhes desse erro.

## Regras para tratamento de erros

Se o erro for do **tipo `HTTPError`** e o código de erro for:
    * 4XX: retorne um código 400 com a mesma mensagem retornada no erro e faça um log com essas informações
    * 5XX: retorne um código 500 com uma mensagem padrão "internal error occurred" e faça um log dando mais detalhes do erro para o servidor

Se o erro for de qualquer outro tipo que não `HTTPError`:
    * retorne um código 500 com uma mensagem padrão "unexpected error occurred"

O tipo `HTTPError` citada no exercício implementa a seguinte interface:

```go
type IHTTPError interface {
    Code() int
    Kind() string
    Error() string
    Unwrap() error
}
```

Você deve utilizar o seguinte tipo como base para retorno do método `Kind()`:

```go
type HTTPErrorKind int

const (
	InternalError HTTPErrorKind = iota
	BusinessError
	ExternalDependencyError
)

func (h HTTPErrorKind) String() string {
	return [...]string{"InternalError", "BusinessError", "ExternalDependencyError"}[h]
}
```

## Mensagens de erro

Você deve utilizar as seguintes mensagens de erro para cada erro que acontecer no servidor:

* 400: 'start' must be a date before 'end'
* 422: bill amount must be paid in full
* 429: rate limit exceeded
* 500: SQL Statement error: 'SELCT' is an invalid statement

## Referências

* https://www.datadoghq.com/blog/go-error-handling/
* https://pkg.go.dev/errors
