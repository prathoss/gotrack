package pkg

type ErrorNotFound struct {
}

func (e ErrorNotFound) Error() string {
	// TODO implement me
	panic("implement me")
}

type ErrorInvalidData struct {
}

func (e ErrorInvalidData) Error() string {
	// TODO implement me
	panic("implement me")
}

type ErrorUnauthorized struct {
}

func (e ErrorUnauthorized) Error() string {
	// TODO implement me
	panic("implement me")
}
