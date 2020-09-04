package goami2

import "fmt"

type goami2Error struct {
	s string
	e string
}

func errorNew(ctx string) *goami2Error {
	return &goami2Error{s: ctx}
}

func (e *goami2Error) msg(msg string, args ...interface{}) *goami2Error {
	txt := fmt.Sprintf(msg, args...)
	e.e = fmt.Sprintf(": %s", txt)
	return e
}

func (e *goami2Error) Error() string {
	return fmt.Sprintf("goami2 error: %s%s", e.s, e.e)
}
