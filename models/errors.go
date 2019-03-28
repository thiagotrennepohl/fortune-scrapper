package models

type ErrCouldNotRetrieveMessages struct {
	Message string
}

func (err *ErrCouldNotRetrieveMessages) Error() string {
	return err.Message
}
