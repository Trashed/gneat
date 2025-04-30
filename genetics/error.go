package genetics

import "errors"

var (
	ErrNilInitialGenome error = errors.New("initial genome can't be nil")
)
