package password

import (
	"crypto/rand"
	"errors"
	"fmt"
)

const (
	LowerCaseChars    = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseChars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberChars       = "0123456789"
	SpecialChars      = " !\"£$%^&*()-_=+]}[{#~'@;:/?.>,<\\|`¬"
	AlphabetChars     = LowerCaseChars + UpperCaseChars
	AlphaNumericChars = AlphabetChars + NumberChars
	AllChars          = AlphaNumericChars + SpecialChars
)

var ErrorBadCharSet = errors.New("availableCharBytes length must be greater" +
	" than 0 and less than or equal to 256")

type Generator struct {
	letterBytes         string
	bitMask             byte
	availableCharLength int
}

func NewGenerator(charSet string) (*Generator, error) {
	availableCharLength := len(charSet)
	if availableCharLength < 2 || availableCharLength > 256 {
		return nil, ErrorBadCharSet
	}
	// Compute bitMask
	var bitLength byte
	var bitMask byte
	for bits := availableCharLength - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}
	bitMask = 1<<bitLength - 1
	return &Generator{
		letterBytes:         charSet,
		bitMask:             bitMask,
		availableCharLength: availableCharLength,
	}, nil
}

// SecureRandomString returns a string of the requested length,
// made from the byte characters provided (only ASCII allowed).
// Uses crypto/rand for security. Will panic if len(availableCharBytes) > 256.
func (g Generator) SecureRandomString(length int) ([]byte, error) {
	bufferSize := length + length/3
	var err error
	result := make([]byte, length)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			// Random byte buffer is empty, get a new one
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return nil, fmt.Errorf("unable to generate secure random bytes: %s", err)
			}
		}
		// Mask bytes to get an index into the character slice
		if idx := int(randomBytes[j%length] & g.bitMask); idx < g.availableCharLength {
			result[i] = g.letterBytes[idx]
			i++
		}
	}
	return result, nil
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) ([]byte, error) {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}
