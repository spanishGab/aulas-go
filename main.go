package main

import "aulas/basics/errorsclass"

func main() {
	err := errorsclass.FakeInternalServerError()
	errorsclass.HandleHTTPError(err)
}
