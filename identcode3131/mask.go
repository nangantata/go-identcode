package identcode3131

import (
	"encoding/binary"

	"github.com/nangantata/go-identcode/internal"
)

func MakeIdentMask(seed []byte) (identMasks [16]uint32) {
	chksum := internal.Sum512(seed, 6)
	b := chksum[:]
	for idx := range 16 {
		identMasks[idx] = uint32(binary.LittleEndian.Uint32(b[idx*4:]) & 0x7FFFFFFF)
	}
	return
}
