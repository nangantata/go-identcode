package identcode

import (
	"errors"
)

var (
	ErrInvalidBase62EncodedBytes = errors.New("invalid base62 encoded bytes")
	ErrInvalidBase62EncodedLen   = errors.New("invalid base62 encoded length")
)
