package bip39

import (
	"crypto/sha512"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type Seed []byte

func NewSeed(mnemonic Mnemonic, password string) Seed {
	return pbkdf2.Key(
		// A mnemonic sentence.
		[]byte(strings.Join(mnemonic, " ")),
		// "mnemonic" + passphrase (again in UTF-8 NFKD) used as the salt.
		[]byte("mnemonic"+password),
		2048,
		64,
		sha512.New,
	)
}
