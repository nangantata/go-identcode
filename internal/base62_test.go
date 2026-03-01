package internal

import (
	"math"
	rand "math/rand/v2"
	"testing"
)

func TestBase62Uint64(t *testing.T) {
	var data []int64 = []int64{
		0, 1, 2, 0x7F0F1F2F3F4F5F6F, math.MaxInt64,
	}
	for range 256 {
		data = append(data, rand.Int64())
	}
	for testIndex, dV := range data {
		b := make([]byte, 0, Uint64Base62EncodedLen)
		enc := Base62AppendEncodeUint64(b, uint64(dV))
		decV, err := Base62DecodeUint64(enc)
		if err != nil {
			t.Fatalf("unexpected error on decode: %v", err)
		}
		if int64(decV) != dV {
			t.Fatalf("unexpected decoded value for test index %d: got %d, want %d", testIndex, decV, dV)
		}
	}
}

func TestBase62Uint32(t *testing.T) {
	var data []int32 = []int32{
		0, 1, 2, 0x7F0F1F2F, math.MaxInt32,
	}
	for range 256 {
		data = append(data, rand.Int32())

	}
	for testIndex, dV := range data {
		b := make([]byte, 0, Uint32Base62EncodedLen)
		enc := Base62AppendEncodeUint32(b, uint32(dV))
		decV, err := Base62DecodeUint32(enc)
		if err != nil {
			t.Fatalf("unexpected error on decode: %v", err)
		}
		if int32(decV) != dV {
			t.Fatalf("unexpected decoded value for test index %d: got %d, want %d", testIndex, decV, dV)
		}
	}
}
