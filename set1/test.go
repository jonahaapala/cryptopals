package set1

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// RunTests - run set1 tests
func RunTests() {
	testHexTo64, _ := GetHexFromString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	verifyHexTo64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result := HexTo64(testHexTo64)
	assert(result, verifyHexTo64)

	testXOr1 := "1c0111001f010100061a024b53535009181c"
	testXOr2 := "686974207468652062756c6c277320657965"
	verifyXOr := "746865206B696420646F6E277420706C6179"
	result, _ = XOrHexString(testXOr1, testXOr2)
	assert(result, verifyXOr)

	testDecrypt, _ := GetHexFromString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	verifyDecrypt := "Cooking MC's like a pound of bacon" // encrypted using 'X'
	results := Decrypt(testDecrypt)
	assert(results['X'], verifyDecrypt)

	testDecrypt2, _ := GetHexFromString("0C1F0C1B10490C050C190108071D490C081D1A491A0C1F0C07491D00040C1A49190C1B490D0810")
	verifyDecrypt2 := "every elephant eats seven times per day" // encrypted using 'i'
	results = Decrypt(testDecrypt2)
	assert(results['i'], verifyDecrypt2)

	testDecryptFromFile := "set1/pswds.txt"
	verifyDecryptFromFile := "Now that the party is jumping\n"
	results2, _ := DecryptFromFile(testDecryptFromFile)
	assert(results2[0], verifyDecryptFromFile)

	testEncrypt := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	verifyEncrypt, _ := GetHexFromString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	res := TextToHex(Encrypt(testEncrypt, []byte("ICE")))
	assert(res.String(), verifyEncrypt.String())

	raw, _ := ioutil.ReadFile("set1/6.txt")
	text := strings.Join(strings.Split(string(raw), "\n"), "")
	testRepeatDecrypt := Decode64(text)
	verifyRepeatDecrypt, _ := ioutil.ReadFile("set1/6-decrypted.txt")
	result = DecryptXor(testRepeatDecrypt)
	assert(result, string(verifyRepeatDecrypt))
}

func assert(test, verify string) {
	if test != verify {
		for i := range test {
			if test[i] != verify[i] {
				msg := fmt.Sprintf("strings differ at position %d: %b != %b", i, test[i], verify[i])
				msg += fmt.Sprintf("\nLen of result = %d\tLen of actual = %d", len(test), len(verify))
				panic(msg)
			}
		}
		panic(test + " != " + verify)
	}
	fmt.Println("PASSED")
}
