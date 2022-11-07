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

type Cleartext string
type Ciphertext string

func (cleartext *Cleartext) Rotate(ciphers ...Cipher) Ciphertext {
	cleartextRunes := []rune(*cleartext)
	for index, currentRune := range cleartextRunes {
		for _, currentCipher := range ciphers {
			if currentNode, hasValue := currentCipher.Ring[currentRune]; hasValue {
				for n := currentCipher.Rotation; n > 0; n-- {
					currentNode = currentNode.Next
				}

				cleartextRunes[index] = currentNode.Value
			}
		}
	}

	return Ciphertext(cleartextRunes)
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
