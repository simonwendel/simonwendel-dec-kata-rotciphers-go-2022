package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCases []struct {
	cleartext string
	expected  string
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

func TestCipher_Rotate_Rot5(t *testing.T) {
	testCases := TestCases{
		{"9*6=42", "4*1=97"},
		{"4*1=97", "9*6=42"},
	}
	ensureRotateCiphersWork(t, Rot5(), testCases)
}

func TestCipher_Rotate_Rot13(t *testing.T) {
	testCases := TestCases{
		{"DECERNO", "QRPREAB"},
		{"QRPREAB", "DECERNO"},
	}
	ensureRotateCiphersWork(t, Rot13(), testCases)
}

func ensureRotateCiphersWork(t *testing.T, cipher Cipher, testCases TestCases) {
	for _, testCase := range testCases {
		actual := cipher.Rotate(testCase.cleartext)
		assert.Equal(t, testCase.expected, actual)
	}
}

func TestMakeRing(t *testing.T) {
	alphabetRunes := []rune("0123456789")
	ring := makeRing("0123456789")

	lastIndex := len(alphabetRunes) - 1
	lastRune := alphabetRunes[lastIndex]
	firstRune := alphabetRunes[0]

	for index := 0; index < lastIndex; index++ {
		currentRune, nextRune := alphabetRunes[index], alphabetRunes[index+1]
		assert.Equal(t, ring[currentRune].Value, currentRune)
		assert.Equal(t, ring[currentRune].Next, ring[nextRune])
	}

	assert.Equal(t, lastRune, ring[lastRune].Value)
	assert.Equal(t, ring[firstRune], ring[lastRune].Next)
}
