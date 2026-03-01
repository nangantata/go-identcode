package identcode3131

import (
	"math"
	"math/rand/v2"
	"testing"
)

func TestPackUnpackB32(t *testing.T) {
	prefixText := "test-"
	seed := []byte("test-seed")
	expectIdentCodeTextLen := len(prefixText) + IdentCodeB32Len
	identMasks := MakeIdentMask(seed)
	dV := []int32{
		0,
		1,
		2,
		1234567890,
		math.MaxInt32,
	}
	for range 256 {
		dV = append(dV, rand.Int32())
	}
	for _, identValue := range dV {
		identCodeText, randomKey := PackB32(prefixText, &identMasks, identValue)
		if len(identCodeText) != expectIdentCodeTextLen {
			t.Errorf("expected identCodeText length %d, got %d", expectIdentCodeTextLen, len(identCodeText))
		}
		unpackedIdentValue, unpackedRandomKey, err := UnpackB32(prefixText, &identMasks, identCodeText)
		if err != nil {
			t.Fatalf("UnpackB32 failed: %v", err)
		}
		if identValue != unpackedIdentValue {
			t.Errorf("expected identValue %d, got %d", identValue, unpackedIdentValue)
		}
		if randomKey != unpackedRandomKey {
			t.Errorf("expected randomKey %d, got %d", randomKey, unpackedRandomKey)
		}
	}
}
