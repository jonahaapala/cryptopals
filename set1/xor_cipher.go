package set1

import (
	uni "unicode"
)

// 						 e	   t	 a	   o	 i	   n	 	   s	 h	   r	 d	   l	 u
var mostCommon = []byte{0x65, 0x74, 0x61, 0x6f, 0x69, 0x6e, 0x20, 0x73, 0x68, 0x72, 0x64, 0x6c, 0x75}

// Decrypt - return decrypted string on a Single-byte XOR cipher
func Decrypt(hex Hex) map[byte]string {
	results := make(map[byte]string)
	mostFreq := getMostFreq(hex)

	for _, val := range mostCommon {
		key := mostFreq ^ val
		guess := DecToHexString(key) // val to repeat for single byte key

		var fullKey Hex // build key from guess
		for i := len(hex) / 2; i > 0; i-- {
			fullKey = append(fullKey, guess...)
		}

		plainHex := HeXOr(hex, fullKey)   // Decrypt with our guessed key
		plainText := HexToBytes(plainHex) // Translate into Plain text

		if allGoodChars(plainText) { // Test if decrypted text uses valid characters
			results[key] = string(plainText)
		}
	}

	return results
}

func getMostFreq(hex Hex) (max byte) {
	freqMap := make(map[byte]int)
	for i := 0; i < len(hex); i += 2 {
		val := hex[i]<<4 + hex[i+1]
		freqMap[val]++
		if freqMap[max] <= freqMap[val] {
			max = val
		}
	}

	return
}

func allGoodChars(test []byte) bool {
	for _, char := range test {
		ch := rune(char)
		if !(uni.IsPunct(ch) || uni.IsLetter(ch) ||
			uni.IsDigit(ch) || uni.IsSpace(ch)) {
			return false
		}
	}

	return true
}
