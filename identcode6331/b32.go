package identcode6331

import (
	"bytes"
	"encoding/binary"
	"fmt"

	identcode "github.com/nangantata/go-identcode"
	"github.com/nangantata/go-identcode/internal"
)

func PackB32(prefixText string, identMasks *[16]uint64, identValue int64) (identCodeText string, randomKey int32) {
	randomKey = internal.GenerateNonZeroRandomKey31()
	prefixBytes := []byte(prefixText)
	prefixLen := len(prefixBytes)
	resultBuf := make([]byte, prefixLen, prefixLen+internal.Int6432Base32EncodedLen)
	copy(resultBuf, prefixBytes)
	var rawIdentCodeBuf [8 + 4]byte
	maskedIdentValue := uint64(identValue) ^ (*identMasks)[int(randomKey&0xF)]
	binary.LittleEndian.PutUint64(rawIdentCodeBuf[:], maskedIdentValue)
	binary.LittleEndian.PutUint32(rawIdentCodeBuf[8:], uint32(randomKey))
	resultBuf = internal.Base32AppendEncode(resultBuf, rawIdentCodeBuf[:])
	identCodeText = string(resultBuf)
	return
}

func UnpackB32(prefixText string, identMasks *[16]uint64, identCodeText string) (identValue int64, randomKey int32, err error) {
	prefixBytes := []byte(prefixText)
	identCodeBytes := []byte(identCodeText)
	if !bytes.HasPrefix(identCodeBytes, prefixBytes) {
		err = identcode.ErrPrefixNotMatch
		return
	}
	prefixLen := len(prefixBytes)
	if (len(identCodeBytes) - prefixLen) != internal.Int6432Base32EncodedLen {
		err = identcode.ErrEncodedLengthNotMatch
		return
	}
	decBuf := make([]byte, 0, (8 + 4))
	if decBuf, err = internal.Base32AppendDecode(decBuf, identCodeBytes[prefixLen:]); err != nil {
		err = fmt.Errorf("cannot decode ident code: %w", err)
		return
	}
	if len(decBuf) != 8+4 {
		err = fmt.Errorf("%w: %d", identcode.ErrInvalidUnpackedBytes, len(decBuf))
		return
	}
	maskedIdentValue := binary.LittleEndian.Uint64(decBuf)
	randomKey = int32(binary.LittleEndian.Uint32(decBuf[8:]))
	identValue = int64(maskedIdentValue ^ (*identMasks)[int(randomKey&0xF)])
	return
}
