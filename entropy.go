package bip39

import (
	"crypto/rand"
	"fmt"
)

type Entropy []byte

func (e Entropy) Valid() error {
	switch l := len(e); {
	case l%4 != 0:
		return fmt.Errorf("ENT in a multiple of 32 bits")
	case l < 16 || 32 < l:
		return fmt.Errorf("allowed size of ENT is 128-256 bits")
	default:
		return nil
	}
}

func NewEntropy(ent int) (Entropy, error) {
	if ent%32 != 0 {
		return nil, fmt.Errorf("ENT in a multiple of 32 bits")
	}
	if ent < 128 || 256 < ent {
		return nil, fmt.Errorf("allowed size of ENT is 128-256 bits")
	}
	entropy := make(Entropy, ent/8)
	if _, err := rand.Read(entropy); err != nil {
		return nil, err
	}
	return entropy, nil
}
