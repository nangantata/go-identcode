package internal

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// generateRandomKey31 create 31 bits random key.
func generateRandomKey31() (randomKey int32) {
	buf := make([]byte, 4)
	rand.Read(buf)
	randomKey = int32(binary.LittleEndian.Uint32(buf) & 0x7FFFFFFF)
	return
}

// GenerateNonZeroRandomKey31 creates non-zero 31 bits random key.
// Return 1 if all attempts fail to generate non-zero random key.
func GenerateNonZeroRandomKey31(maxAttempts int) (randomKey int32) {
	remain := maxAttempts - 1
	for remain > 0 {
		if randomKey = generateRandomKey31(); randomKey != 0 {
			return
		}
		remain--
	}
	if randomKey = generateRandomKey31(); randomKey != 0 {
		return
	}
	if randomKey = int32(uint64(time.Now().UnixNano()) & 0x7FFFFFFF); randomKey != 0 {
		return
	}
	return 1
}
