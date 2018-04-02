package set1

const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// HexTo64 - convert input hex value to base64
func HexTo64(hex Hex) (base64 string) {
	var pad int // pad with zeros to make len divisible by 3
	switch len(hex) % 3 {
	case 1:
		pad = 1 // change last char to be =
		hex = append(hex, 0, 0)
	case 2:
		pad = 2 // append == to final result
		hex = append(hex, 0)
	}

	for len(hex) > 0 {
		cur := hex[:3] // get first three bytes
		base64 += encode(cur)
		hex = hex[3:] // remove first three bytes
	}

	if pad == 1 {
		base64 = base64[:len(base64)-1] + "="
	} else if pad == 2 {
		base64 += "=="
	}

	return
}

func encode(n Hex) string {
	x := n[0]<<2 + n[1]>>2
	y := (n[1]&3)<<4 + n[2]
	return string([]byte{b64[x], b64[y]})
}

// Decode64 - return decoded text
func Decode64(text string) (decoded []byte) {
	if len(text)%4 == 0 {
		for len(text) > 0 {

			index := make(map[byte]int, len(b64))
			for i := range b64 {
				index[b64[i]] = i
			}

			for len(text) > 0 {
				a := index[text[0]]
				b := index[text[1]]
				c := index[text[2]]
				d := index[text[3]]

				decoded = append(decoded, byte(a<<2+b>>4))
				decoded = append(decoded, byte((b&15)<<4+c>>2))
				decoded = append(decoded, byte((c&3)<<6+d))

				text = text[4:] // remove first four
			}
		}
	}

	return
}
