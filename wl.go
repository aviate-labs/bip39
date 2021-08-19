package bip39

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"

	_ "embed"
)

func init() {
	if cs := strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(english))), 16); cs != "c1dbd296" {
		panic(fmt.Sprintf("Invalid checksum: %s", cs))
	}
}

// Source: https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
//go:embed testdata/english.txt
var english string
var English = NewWordList(
	strings.Split(
		strings.TrimSpace(english),
		"\n",
	),
)

func NewWordList(list []string) WordList {
	lookup := make(map[string]int)
	for i, l := range list {
		lookup[l] = i
	}
	return WordList{
		list:   list,
		lookup: lookup,
	}
}

type WordList struct {
	list   []string
	lookup map[string]int
}

func (l WordList) Get(w string) (int, bool) {
	i, ok := l.lookup[w]
	return i, ok
}

func (l WordList) List() []string {
	return l.list
}
