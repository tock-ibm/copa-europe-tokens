package tokens

import "net/http"

type TokenErr interface {
	error
	Status() int
}

type DefaultTokenErr struct {
	ErrMsg string
}

func (e *DefaultTokenErr) Error() string {
	return e.ErrMsg
}

func (e *DefaultTokenErr) Status() int {
	return http.StatusInternalServerError
}

type ErrExist DefaultTokenErr
type ErrInvalid DefaultTokenErr
type ErrNotFound DefaultTokenErr
type ErrPermission DefaultTokenErr

func (e *ErrExist) Error() string {
	return e.ErrMsg
}

func (e *ErrExist) Status() int {
	return http.StatusConflict
}

func (e *ErrInvalid) Error() string {
	return e.ErrMsg
}

func (e *ErrInvalid) Status() int {
	return http.StatusBadRequest
}

func (e *ErrNotFound) Error() string {
	return e.ErrMsg
}

func (e *ErrNotFound) Status() int {
	return http.StatusNotFound
}

func (e *ErrPermission) Error() string {
	return e.ErrMsg
}

func (e *ErrPermission) Status() int {
	return http.StatusForbidden
}

func ErrorStatus(err error) int {
	switch e := err.(type) {
	case TokenErr:
		return e.Status()
	default:
		return http.StatusInternalServerError
	}
}
