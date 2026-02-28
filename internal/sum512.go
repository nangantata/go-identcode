package internal

import (
	"crypto/sha512"
)

func Sum512(d []byte, iter int) (chksum [64]byte) {
	remain := iter - 1
	for remain > 0 {
		chksum = sha512.Sum512(d)
		d = chksum[:]
		remain--
	}
	chksum = sha512.Sum512(d)
	return
}
