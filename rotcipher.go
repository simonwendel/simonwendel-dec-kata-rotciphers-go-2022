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

func Rot5() Cipher {
	return RotN(5, "0123456789")
}

func Rot13() Cipher {
	return RotN(13, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func RotN(n int, alphabet string) Cipher {
	return Cipher{
		Rotation: n,
		Ring:     makeRing(alphabet),
	}
}

func (c *Cipher) Rotate(cleartext string) string {
	runes := []rune(cleartext)
	for i, rune := range runes {
		if rotated, hasValue := c.Ring[rune]; hasValue {
			for n := c.Rotation; n > 0; n-- {
				rotated = rotated.Next
			}

			runes[i] = rotated.Value
		}
	}

	return string(runes)
}

func makeRing(alphabet string) Ring {
	runes := []rune(alphabet)
	ring := make(Ring)

	for _, number := range runes {
		ring[number] = &Node{Value: number}
	}

	for i := 0; i < len(runes); i++ {
		curr := runes[i]

		var next rune
		if i == len(runes)-1 {
			next = runes[0]
		} else {
			next = runes[i+1]
		}

		ring[curr].Next = ring[next]
	}

	return ring
}
