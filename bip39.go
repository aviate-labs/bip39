package bip39

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

var (
	bi2047 = big.NewInt(0x07FF) // 0111_1111_1111
	bi2048 = big.NewInt(0x0800) // 1000_0000_0000
)

type Mnemonic []string

func (l WordList) NewMnemonic(ent Entropy) (Mnemonic, error) {
	if err := ent.Valid(); err != nil {
		return nil, err
	}
	// ENT
	data := new(big.Int).SetBytes(ent)
	// Add CS to ENT.
	e := sha256.Sum256(ent)[0]
	for i := 0; i < len(ent)/4; i++ {
		b := ((e << i) & 128) >> 7
		data.Lsh(data, 1)
		data.Or(data, big.NewInt(int64(b)))
	}
	var (
		word  = new(big.Int)
		words Mnemonic
	)
	for i := 0; i < 33*len(ent)/44; i++ {
		// Split into groups of 11 bits.
		word.And(data, bi2047)
		data.Div(data, bi2048)
		words = append(
			// Convert the index into a word.
			[]string{l.list[binary.BigEndian.Uint16(
				padLeft(word.Bytes(), 2),
			)]},
			words...,
		)
	}
	return words, nil
}

func padLeft(bs []byte, n int) []byte {
	for len(bs) != n {
		bs = append([]byte{0x00}, bs...)
	}
	return bs
}
