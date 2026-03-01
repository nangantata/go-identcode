package identcode6331

import (
	"encoding/binary"

	"github.com/nangantata/go-identcode/internal"
)

func MakeIdentMask(seed []byte) (identMasks [16]uint64) {
	chksum := internal.Sum512(seed, 3)
	b := chksum[:]
	for idx := range 8 {
		identMasks[idx] = uint64(binary.LittleEndian.Uint64(b[idx*8:]) & 0x7FFFFFFFFFFFFFFF)
	}
	chksum = internal.Sum512(b, 5)
	b = chksum[:]
	for idx := range 8 {
		identMasks[idx+8] = uint64(binary.LittleEndian.Uint64(b[idx*8:]) & 0x7FFFFFFFFFFFFFFF)
	}
	return
}
