package set1

import "errors"

// Hex - each entry is really a 4-bit "nibble"
type Hex []byte

const b16 = "0123456789ABCDEF"

// GetHexFromString - return a valid Hexidecimal number
func GetHexFromString(in string) (Hex, error) {
	if len(in)&1 != 0 {
		return nil, errors.New("Invalid Hex: Length must be even")
	}

	var offset int
	hex := make(Hex, len(in))
	for i, cur := range in {
		if cur >= '0' && cur <= '9' {
			offset = '0'
		} else if cur >= 'a' && cur <= 'f' {
			offset = 'a' - 10
		} else if cur >= 'A' && cur <= 'F' {
			offset = 'A' - 10
		} else {
			return nil, errors.New("Invalid Hex: Character " + string(cur))
		}
		hex[i] = byte(int(cur) - offset)
	}

	return hex, nil
}

// HeXOr - return input Hex values XOr'd together
func HeXOr(h1, h2 Hex) Hex {
	diff := len(h1) - len(h2)
	if diff > 0 {
		h1, h2 = h2, h1 // force h1 shorter than h2
	} else {
		diff *= -1 // Ensure diff is positive
	}

	xor := make(Hex, 0, len(h2))
	xor = append(xor, h2[:diff]...)
	for i, nib := range h1 {
		val := nib ^ h2[i+diff]
		xor = append(xor, val)
	}
	return xor
}

// TextToHex - return hex version of plain text
func TextToHex(text []byte) (hex Hex) {
	for _, char := range text {
		hex = append(hex, DecToHexString(char)...)
	}

	return
}

// HexToBytes - return plain text of given Hex value
func HexToBytes(hex Hex) []byte {
	out := make([]byte, len(hex)/2)
	for i := 0; i < len(hex); i += 2 {
		num := hex[i]<<4 + hex[i+1]
		out[i/2] = num
	}

	return out
}

// DecToHexString - translate decimal number into hex representation
func DecToHexString(dec byte) Hex {
	return Hex{dec >> 4, dec & 0x0F}
}

func (hex Hex) String() string {
	base16 := make([]byte, len(hex))
	for i, nib := range hex {
		base16[i] = b16[nib]
	}
	return string(base16)
}
