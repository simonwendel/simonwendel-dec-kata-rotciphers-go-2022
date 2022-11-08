// Complicated RotCiphers - Rotating ciphers using complicated data structures
// Copyright (C) 2022  Simon Wendel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCases []struct {
	cleartext Cleartext
	expected  Ciphertext
}

func TestMakeRing(t *testing.T) {
	alphabetRunes := []rune("Q024ABC68")
	ring := makeRing("Q024ABC68")

	for index := 0; index < len(alphabetRunes); index++ {
		currentRune, nextRune :=
			alphabetRunes[index],
			alphabetRunes[(index+1)%len(alphabetRunes)]
		assert.Equal(t, currentRune, ring[currentRune].Value)
		assert.Equal(t, ring[nextRune], ring[currentRune].Next)
	}
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
