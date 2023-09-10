package shared

type ErrorMessage string

func (e ErrorMessage) String() string {
	return string(e)
}

const (
	ErrorMissingParam = ErrorMessage("invalid/missing param in request path")
)
