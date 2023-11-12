package kvnuts

type ErrNotAPointerType struct{}

func (ErrNotAPointerType) Error() string {
	return "passed should be a pointer type"
}
