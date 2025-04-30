package genetics

import "errors"

var (
	ErrNilInitialGenome error = errors.New("initial genome can't be nil")
	ErrEmptyGenomeFile  error = errors.New("genome file shouldn't be empty")
)
