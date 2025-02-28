package ex42

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func Ex42() {
	var hashType string
	flag.StringVar(&hashType, "hash", "sha256", "Hash algorithm to use (sha256, sha384, sha512)")
	flag.Parse()
	var hash []byte
	switch hashType {
	case "sha256":
		hash = computeHash(sha256.New, os.Stdin)
	case "sha384":
		hash = computeHash(sha512.New384, os.Stdin)
	case "sha512":
		hash = computeHash(sha512.New, os.Stdin)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported hash algorithm: %s\n", hashType)
		os.Exit(1)
	}

	// Print the hash in hexadecimal format
	fmt.Println(hex.EncodeToString(hash))
}

func computeHash(newHash func() hash.Hash, r io.Reader) []byte {
	h := newHash()
	io.Copy(h, r)
	return h.Sum(nil)
}
