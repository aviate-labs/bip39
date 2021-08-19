package bip39

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	_ "embed"
)

func init() {
	var vs map[string][][]string
	if err := json.Unmarshal(vectors, &vs); err != nil {
		panic(err)
	}
	for _, v := range vs["english"] {
		ent, err := hex.DecodeString(v[0])
		if err != nil {
			panic(err)
		}
		seed, err := hex.DecodeString(v[2])
		if err != nil {
			panic(err)
		}
		testVectors = append(testVectors, vector{
			ent:   ent,
			words: strings.Split(v[1], " "),
			seed:  seed,
		})
	}
}

//go:embed testdata/vectors.json
var vectors []byte
var testVectors []vector

type vector struct {
	ent   Entropy
	words Mnemonic
	seed  Seed
}

func TestEnglish_NewMnemonic(t *testing.T) {
	for _, test := range testVectors {
		words, err := English.NewMnemonic(test.ent)
		if err != nil {
			t.Fatal(err)
		}
		for i, w := range test.words {
			if words[i] != w {
				t.Errorf("expected: %s, got %s", w, words[i])
			}
		}
	}
}
