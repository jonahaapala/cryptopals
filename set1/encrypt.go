package set1

// Encrypt - encrypt plaintext using XOr cipher
func Encrypt(plain, key []byte) (cipher []byte) {
	plainHex := TextToHex(plain)
	keyHex := TextToHex(key)

	for i := 0; i < len(plainHex); i += 2 {
		j := i % len(keyHex)
		ch, key := plainHex[i:i+2], keyHex[j:j+2]
		tmp := HexToBytes(HeXOr(ch, key))
		cipher = append(cipher, tmp...)
	}

	return
}
