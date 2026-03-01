package identcode

import (
	"errors"
)

var (
	ErrPrefixNotMatch        = errors.New("prefix not match")
	ErrEncodedLengthNotMatch = errors.New("encoded length not match")
	ErrInvalidUnpackedBytes  = errors.New("invalid unpacked bytes")

	ErrInvalidBase62EncodedBytes = errors.New("invalid base62 encoded bytes")
	ErrInvalidBase62EncodedLen   = errors.New("invalid base62 encoded length")
)
