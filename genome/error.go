package genome

import "errors"

var (
	ErrNilInitialGenome error = errors.New("initial genome can't be nil")
)
