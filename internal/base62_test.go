package internal

import (
	"math"
	rand "math/rand/v2"
	"testing"
)

func TestBase62Int64(t *testing.T) {
	var data []int64 = []int64{
		0, 1, 2, 0x7F0F1F2F3F4F5F6F, math.MaxInt64,
	}
	for range 256 {
		data = append(data, rand.Int64())
	}
	for testIndex, dV := range data {
		b := make([]byte, 0, Int64Base62EncodedLen)
		enc, err := Base62AppendEncodeInt64(b, dV)
		if err != nil {
			t.Fatalf("unexpected error on encode: %v", err)
		}
		decV, err := Base62DecodeInt64(enc)
		if err != nil {
			t.Fatalf("unexpected error on decode: %v", err)
		}
		if decV != dV {
			t.Fatalf("unexpected decoded value for test index %d: got %d, want %d", testIndex, decV, dV)
		}
	}
}

func TestBase62Int32(t *testing.T) {
	var data []int32 = []int32{
		0, 1, 2, 0x7F0F1F2F, math.MaxInt32,
	}
	for range 256 {
		data = append(data, rand.Int32())

	}
	for testIndex, dV := range data {
		b := make([]byte, 0, Int32Base62EncodedLen)
		enc, err := Base62AppendEncodeInt32(b, dV)
		if err != nil {
			t.Fatalf("unexpected error on encode: %v", err)
		}
		decV, err := Base62DecodeInt32(enc)
		if err != nil {
			t.Fatalf("unexpected error on decode: %v", err)
		}
		if decV != dV {
			t.Fatalf("unexpected decoded value for test index %d: got %d, want %d", testIndex, decV, dV)
		}
	}
}
