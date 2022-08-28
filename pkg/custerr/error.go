package custerr

type CustomError struct {
	Err     error
	ErrCode int
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}
