package errorsclass

import (
	"errors"
	"fmt"
)

func HandleExactError(err error) {
	if errors.Is(err, ErrInvalidQueryParam) {
		fmt.Printf("O query param enviado não é válido: %s", err)
	} else if errors.Is(err, ErrInvalidRequestBody) {
		fmt.Printf("O request body enviado contém erros: %s", err)
	} else {
		fmt.Printf("Um erro inesperado aconteceu: %s", err)
	}
}

func HandleHTTPError(err error) {
	// var tErr *HTTPError
	// err, ok := errors.As(err, &tErr)
	if err, ok := errors.AsType[*HTTPError](err); ok {
		if err.Code() == 400 {
			fmt.Printf("Reavalie os dados de entrada da requisição. O erro %q ocorreu com o código %d e tipo %q. Causa: %s",
				err.Unwrap(),
				err.Code(),
				err.Kind(),
				err.Error(),
			)
		} else if err.Code() == 422 {
			fmt.Printf("Não foi possível concluir a requisição. O erro %q ocorreu com o código %d e tipo %q. Causa: %s",
				err.Unwrap(),
				err.Code(),
				err.Kind(),
				err.Error(),
			)
		} else if err.Code() >= 500 {
			fmt.Printf("O erro %q ocorreu com o código %d e tipo %q",
				err.Unwrap(),
				err.Code(),
				err.Kind(),
			)
		}

	} else {
		fmt.Printf("Um erro inesperado aconteceu: %s", err)
	}
}
