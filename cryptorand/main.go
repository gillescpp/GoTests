package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	// Example: this will give us a 44 byte, base64 encoded output
	token, err := GenerateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("", err)
	}
	fmt.Println("", token)
}

//links
//	https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
