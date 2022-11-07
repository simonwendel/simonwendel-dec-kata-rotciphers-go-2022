package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCases []struct {
	cleartext Cleartext
	expected  Ciphertext
}

func TestRot5(t *testing.T) {
	rot5 := Rot5()
	assert.Equal(t, 5, rot5.Rotation)
	assert.Equal(t, 10, len(rot5.Ring))
}

func TestRot13(t *testing.T) {
	rot13 := Rot13()
	assert.Equal(t, 13, rot13.Rotation)
	assert.Equal(t, 26, len(rot13.Ring))
}

func TestCleartext_Rotate_Rot5(t *testing.T) {
	testCases := TestCases{
		{"9*6=42", "4*1=97"},
		{"4*1=97", "9*6=42"},
	}
	ensureRotateCiphersWork(t, testCases, Rot5())
}

func TestCleartext_Rotate_Rot13(t *testing.T) {
	testCases := TestCases{
		{"DECERNO", "QRPREAB"},
		{"QRPREAB", "DECERNO"},
	}
	ensureRotateCiphersWork(t, testCases, Rot13())
}

func TestCleartext_Rotate_Rot5Rot13(t *testing.T) {
	testCases := TestCases{
		{"DECERNO 9*6=42", "QRPREAB 4*1=97"},
		{"QRPREAB 4*1=97", "DECERNO 9*6=42"},
	}
	ensureRotateCiphersWork(t, testCases, Rot5(), Rot13())
}

func ensureRotateCiphersWork(t *testing.T, testCases TestCases, ciphers ...Cipher) {
	for _, testCase := range testCases {
		actual := testCase.cleartext.Rotate(ciphers...)
		assert.Equal(t, testCase.expected, actual)
	}
}

func TestMakeRing(t *testing.T) {
	alphabetRunes := []rune("0123456789")
	ring := makeRing("0123456789")

	for index := 0; index < len(alphabetRunes); index++ {
		currentRune, nextRune :=
			alphabetRunes[index],
			alphabetRunes[(index+1)%len(alphabetRunes)]
		assert.Equal(t, ring[currentRune].Value, currentRune)
		assert.Equal(t, ring[currentRune].Next, ring[nextRune])
	}
}
