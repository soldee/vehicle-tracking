package response

type InternalError struct {
	Msg  string
	Code int
}

func (err *InternalError) Error() string {
	return err.Msg
}

type InvalidInput struct {
	Msg  string
	Code int
}

func (err *InvalidInput) Error() string {
	return err.Msg
}

type NotFound struct {
	Msg  string
	Code int
}

func (err *NotFound) Error() string {
	return err.Msg
}
