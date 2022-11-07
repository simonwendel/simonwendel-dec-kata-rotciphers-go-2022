package main

type Ring map[rune]*Node

type Node struct {
	Value rune
	Next  *Node
}

type Cipher struct {
	Rotation int
	Ring     Ring
}

func (cipher *Cipher) Rotate(cleartext string) string {
	cleartextRunes := []rune(cleartext)
	for index, singleRune := range cleartextRunes {
		if nodeToRotate, hasValue := cipher.Ring[singleRune]; hasValue {
			for n := cipher.Rotation; n > 0; n-- {
				nodeToRotate = nodeToRotate.Next
			}

			cleartextRunes[index] = nodeToRotate.Value
		}
	}

	return string(cleartextRunes)
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

func makeRing(alphabet string) Ring {
	alphabetRunes := []rune(alphabet)
	ring := make(Ring)

	for _, currentRune := range alphabetRunes {
		ring[currentRune] = &Node{Value: currentRune}
	}

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
