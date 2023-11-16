package response

type InternalError struct {
	Msg string
}

func (err *InternalError) Error() string {
	return err.Msg
}

type InvalidInput struct {
	Msg string
}

func (err *InvalidInput) Error() string {
	return err.Msg
}

type NotFound struct {
	Msg string
}

func (err *NotFound) Error() string {
	return err.Msg
}
