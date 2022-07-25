package common

type Error struct {
	Note string
}

func (e *Error) Error() string {
	return e.Note
}

func NewErr(note string) *Error {
	err := &Error{
		Note: note,
	}
	return err
}
