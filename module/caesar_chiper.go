package module

import "unicode"

type CaesarCipher struct {
	Text        string
	Shift       int
	ShiftedText []rune
}

func (c *CaesarCipher) shiftCharacter(char rune, shift int) rune {
	if unicode.IsLetter(char) {
		shiftBase := rune(shift % 26)
		asciiOffset := 'A'
		if unicode.IsLower(char) {
			asciiOffset = 'a'
		}
		return asciiOffset + (char-asciiOffset+shiftBase)%26
	}
	return char
}

func (c *CaesarCipher) Encrypt() {
	c.ShiftedText = []rune(c.Text)
	for i, char := range c.ShiftedText {
		c.ShiftedText[i] = c.shiftCharacter(char, c.Shift)
	}
}

func (c *CaesarCipher) Decrypt() {
	c.ShiftedText = []rune(c.Text)
	for i, char := range c.ShiftedText {
		c.ShiftedText[i] = c.shiftCharacter(char, -c.Shift)
	}
}

func (c *CaesarCipher) GetEncryptedText() string {
	return string(c.ShiftedText)
}

func (c *CaesarCipher) GetDecryptedText() string {
	return string(c.ShiftedText)
}
