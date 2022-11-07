package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RotationTestData []struct {
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
	cases := RotationTestData{
		{"9*6=42", "4*1=97"},
		{"4*1=97", "9*6=42"},
	}
	ensureRotateCiphersWork(t, Rot5(), cases)
}

func TestCipher_Rotate_Rot13(t *testing.T) {
	cases := RotationTestData{
		{"DECERNO", "QRPREAB"},
		{"QRPREAB", "DECERNO"},
	}
	ensureRotateCiphersWork(t, Rot13(), cases)
}

func ensureRotateCiphersWork(t *testing.T, cipher Cipher, cases RotationTestData) {
	for _, c := range cases {
		actual := cipher.Rotate(c.cleartext)
		assert.Equal(t, c.expected, actual)
	}
}

func TestMakeRing(t *testing.T) {
	runes := []rune("0123456789")
	ring := makeRing("0123456789")

	lastIndex := len(runes) - 1

	for i := 0; i < lastIndex; i++ {
		curr, next := runes[i], runes[i+1]
		assert.Equal(t, ring[curr].Value, curr)
		assert.Equal(t, ring[curr].Next, ring[next])
	}

	assert.Equal(t, runes[lastIndex], ring[runes[lastIndex]].Value)
	assert.Equal(t, ring[runes[0]], ring[runes[lastIndex]].Next)
}
