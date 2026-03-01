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
	Uint64Base62EncodedLen = 11
	Uint32Base62EncodedLen = 6
)

func init() {
	copy(b62DecodeMap[:], b62DecodeMapInitialize)
	for idx, ch := range b62EncoderChars {
		b62DecodeMap[ch] = byte(idx)
	}
}

func Base62DecodeUint64(src []byte) (r uint64, err error) {
	if len(src) < Uint64Base62EncodedLen {
		err = fmt.Errorf("%w: %d", identcode.ErrInvalidBase62EncodedLen, len(src))
		return
	}
	if v := b62DecodeMap[src[0]]; v == b62InvalidIndex {
		err = fmt.Errorf("%w: invalid character 0x%02X at position 0", identcode.ErrInvalidBase62EncodedBytes, src[0])
		return
	} else {
		r = uint64(v)
	}
	for idx := 1; idx < Uint64Base62EncodedLen; idx++ {
		ch := src[idx]
		v := b62DecodeMap[ch]
		if v == b62InvalidIndex {
			err = fmt.Errorf("%w: invalid character 0x%02X at position %d", identcode.ErrInvalidBase62EncodedBytes, ch, idx)
			return
		}
		r = r*62 + uint64(v)
	}
	return
}

func Base62AppendEncodeUint64(dst []byte, src uint64) []byte {
	var b [11]byte
	for idx := (Uint64Base62EncodedLen - 1); idx > 0; idx-- {
		cIndex := src % 62
		cByte := b62EncoderChars[cIndex]
		b[idx] = cByte
		src /= 62
	}
	b[0] = b62EncoderChars[src%62]
	dst = append(dst, b[:]...)
	return dst
}

func Base62DecodeUint32(src []byte) (r uint32, err error) {
	if len(src) < Uint32Base62EncodedLen {
		err = fmt.Errorf("%w: %d", identcode.ErrInvalidBase62EncodedLen, len(src))
		return
	}
	if v := b62DecodeMap[src[0]]; v == b62InvalidIndex {
		err = fmt.Errorf("%w: invalid character 0x%02X at position 0", identcode.ErrInvalidBase62EncodedBytes, src[0])
		return
	} else {
		r = uint32(v)
	}
	for idx := 1; idx < Uint32Base62EncodedLen; idx++ {
		ch := src[idx]
		v := b62DecodeMap[ch]
		if v == b62InvalidIndex {
			err = fmt.Errorf("%w: invalid character 0x%02X at position %d", identcode.ErrInvalidBase62EncodedBytes, ch, idx)
			return
		}
		r = r*62 + uint32(v)
	}
	return
}

func Base62AppendEncodeUint32(dst []byte, src uint32) []byte {
	var b [6]byte
	for idx := (Uint32Base62EncodedLen - 1); idx > 0; idx-- {
		cIndex := src % 62
		cByte := b62EncoderChars[cIndex]
		b[idx] = cByte
		src /= 62
	}
	b[0] = b62EncoderChars[src%62]
	dst = append(dst, b[:]...)
	return dst
}
