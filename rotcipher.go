// RotCiphers - Rotating ciphers using complicated data structures
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

type Cleartext string
type Ciphertext string

type Cipher struct {
	Rotation int
	Ring     Ring
}

type Ring map[rune]*Node

type Node struct {
	Value rune
	Next  *Node
}

func (node *Node) Translate(distance int) Node {
	if distance == 0 {
		return *node
	}

	return node.Next.Translate(distance - 1)
}

func Rot5() Cipher {
	return rotationCipher(5, "0123456789")
}

func Rot13() Cipher {
	return rotationCipher(13, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func rotationCipher(rotations int, alphabet string) Cipher {
	return Cipher{
		Rotation: rotations,
		Ring:     makeRing(alphabet),
	}
}

func (cleartext *Cleartext) Rotate(ciphers ...Cipher) Ciphertext {
	runes := []rune(*cleartext)
	for index, currentRune := range runes {
		runes[index] = rotateRune(currentRune, ciphers...)
	}

	return Ciphertext(runes)
}

func rotateRune(currentRune rune, ciphers ...Cipher) rune {
	for _, currentCipher := range ciphers {
		if currentNode, hasValue := currentCipher.Ring[currentRune]; hasValue {
			currentRune = currentNode.Translate(currentCipher.Rotation).Value
		}
	}

	return currentRune
}

func makeRing(alphabet string) Ring {
	alphabetRunes := []rune(alphabet)
	ring := make(Ring)

	ring.fill(alphabetRunes)

	for index, currentRune := range alphabetRunes {
		var nextRune rune
		if index == len(alphabetRunes)-1 {
			nextRune = alphabetRunes[0]
		} else {
			nextRune = alphabetRunes[index+1]
		}

		ring[currentRune].Next = ring[nextRune]
	}

	return ring
}

func (ring *Ring) fill(alphabetRunes []rune) {
	for _, currentRune := range alphabetRunes {
		(*ring)[currentRune] = &Node{Value: currentRune}
	}
}
