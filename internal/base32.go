package internal

import (
	"encoding/base32"
)

const b32EncoderChars = "AbCdeFGHijkLMnopQrstuVwxyz234567"

var b32Encoding = base32.NewEncoding(b32EncoderChars).WithPadding(base32.NoPadding)

const (
	Int6432Base32EncodedLen = 20
	Int3232Base32EncodedLen = 13
)

func Base32AppendDecode(dst, src []byte) ([]byte, error) {
	return b32Encoding.AppendDecode(dst, src)
}

func Base32AppendEncode(dst, src []byte) []byte {
	return b32Encoding.AppendEncode(dst, src)
}
