package domain

type ErrorType string

func (u ErrorType) Error() string {
	return string(u)
}
