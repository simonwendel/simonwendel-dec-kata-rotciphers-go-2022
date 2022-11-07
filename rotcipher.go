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
