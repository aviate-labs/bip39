package bip39

import (
	"encoding/hex"
	"testing"
)

func TestEnglish_NewSeed(t *testing.T) {
	for _, test := range testVectors {
		words, err := English.NewMnemonic(test.ent)
		if err != nil {
			t.Fatal(err)
		}
		seed := hex.EncodeToString(NewSeed(words, "TREZOR"))
		testSeed := hex.EncodeToString(test.seed)
		if testSeed != seed {
			t.Errorf("expected: %s, got %s", testSeed, seed)
		}
	}
}
