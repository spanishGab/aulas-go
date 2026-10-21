package main

import "aulas/basics/errorsclass"

func main() {
	err := errorsclass.ErrInternal
	errorsclass.HandleHTTPError(err)
}
