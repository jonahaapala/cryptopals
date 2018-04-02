package set1

// DecryptXor - Decrypt Xor ciphers by guessing the key
func DecryptXor(cipher []byte) (plain string) {
	max := 40 // absolute max key length
	if l := len(cipher) / 2; l < max {
		max = l // max key length we can possibly check
	}

	// 1. Assume key length is between 2 and at most 40.
	var keyLen int
	var minAvg = 8.0 // This should be big enough to be overridden
	for k := 2; k <= max; k++ {
		// 2. Calculate Hamming distance and normalize by key length 'k'
		sumAvg := 0.0
		for i := 0; i < len(cipher); i += 2 * k {
			d := hamming(cipher[i:i+k], cipher[i+k:i+2*k])
			sumAvg += float64(d) / float64(k)
		}
		// Average the averages for all blocks of length 'k'
		iter := float64(len(cipher)) / float64(2*k)
		avg := sumAvg / iter

		// 3. The key with the min avg distance will be the true key length
		if avg < minAvg {
			minAvg, keyLen = avg, k
		}
	}

	// 4. Break up cipher text into blocks of length D and transpose to group values
	//    encrypted using the same charater from the key.
	blocks := blockTranspose(cipher, keyLen)

	// 5. Decrypt each new block with the single character XOR decryption function.
	realKey := make([]byte, 0, keyLen)
	for _, block := range blocks {
		hexBlock := TextToHex(block)
		res := Decrypt(hexBlock)

		for keyByte := range res {
			realKey = append(realKey, keyByte)
			break // assume first found value is right
		}
	}

	// 6. Construct the key and decrypt using the repeated key "Encrypt()" function.
	return string(Encrypt(cipher, realKey))
}

func hamming(s1, s2 []byte) (diff int) {
	if len(s1) != len(s2) {
		return -1
	}

	for i := range s1 {
		xor := s1[i] ^ s2[i]
		for xor > 0 {
			diff += int(xor & 1) // count 1's
			xor >>= 1            // bitshift right
		}
	}

	return
}

func blockTranspose(s []byte, l int) [][]byte {
	blocks := make([][]byte, l)

	for i := 0; i < l; i++ {
		for j := i; j < len(s); j += l {
			blocks[i] = append(blocks[i], s[j])
		}
	}

	return blocks
}
