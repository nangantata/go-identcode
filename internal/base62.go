package internal

import (
	"fmt"

	identcode "github.com/nangantata/go-identcode"
)

const b62EncoderChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var b62DecodeMap [256]byte

const (
	b62DecodeMapInitialize = "" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"
	b62InvalidIndex = '\xff'
)

const (
	Int64Base62EncodedLen = 11
	Int32Base62EncodedLen = 6
)

func init() {
	copy(b62DecodeMap[:], b62DecodeMapInitialize)
	for idx, ch := range b62EncoderChars {
		b62DecodeMap[ch] = byte(idx)
	}
}

func Base62DecodeInt64(src []byte) (r int64, err error) {
	if len(src) < Int64Base62EncodedLen {
		err = fmt.Errorf("%w: %d", identcode.ErrInvalidBase62EncodedLen, len(src))
		return
	}
	if v := b62DecodeMap[src[0]]; v == b62InvalidIndex {
		err = fmt.Errorf("%w: invalid character 0x%02X at position 0", identcode.ErrInvalidBase62EncodedBytes, src[0])
		return
	} else {
		r = int64(v)
	}
	for idx := 1; idx < Int64Base62EncodedLen; idx++ {
		ch := src[idx]
		v := b62DecodeMap[ch]
		if v == b62InvalidIndex {
			err = fmt.Errorf("%w: invalid character 0x%02X at position %d", identcode.ErrInvalidBase62EncodedBytes, ch, idx)
			return
		}
		r = r*62 + int64(v)
	}
	return
}

func Base62AppendEncodeInt64(dst []byte, src int64) ([]byte, error) {
	var b [11]byte
	for idx := (Int64Base62EncodedLen - 1); idx > 0; idx-- {
		cIndex := src % 62
		cByte := b62EncoderChars[cIndex]
		b[idx] = cByte
		src /= 62
	}
	b[0] = b62EncoderChars[src%62]
	dst = append(dst, b[:]...)
	return dst, nil
}

func Base62DecodeInt32(src []byte) (r int32, err error) {
	if len(src) < Int32Base62EncodedLen {
		err = fmt.Errorf("%w: %d", identcode.ErrInvalidBase62EncodedLen, len(src))
		return
	}
	if v := b62DecodeMap[src[0]]; v == b62InvalidIndex {
		err = fmt.Errorf("%w: invalid character 0x%02X at position 0", identcode.ErrInvalidBase62EncodedBytes, src[0])
		return
	} else {
		r = int32(v)
	}
	for idx := 1; idx < Int32Base62EncodedLen; idx++ {
		ch := src[idx]
		v := b62DecodeMap[ch]
		if v == b62InvalidIndex {
			err = fmt.Errorf("%w: invalid character 0x%02X at position %d", identcode.ErrInvalidBase62EncodedBytes, ch, idx)
			return
		}
		r = r*62 + int32(v)
	}
	return
}

func Base62AppendEncodeInt32(dst []byte, src int32) ([]byte, error) {
	var b [6]byte
	for idx := (Int32Base62EncodedLen - 1); idx > 0; idx-- {
		cIndex := src % 62
		cByte := b62EncoderChars[cIndex]
		b[idx] = cByte
		src /= 62
	}
	b[0] = b62EncoderChars[src%62]
	dst = append(dst, b[:]...)
	return dst, nil
}
