package set1

import "errors"

// XOrHexString - return XOr of input hex values
func XOrHexString(in1, in2 string) (string, error) {
	if len(in1) != len(in2) {
		return "", errors.New("Cannot XOR values of different length")
	}

	hex1, err1 := GetHexFromString(in1)
	if err1 != nil {
		return "", err1 // Invalid Hex
	}

	hex2, err2 := GetHexFromString(in2)
	if err2 != nil {
		return "", err2 // Invalid Hex
	}

	out := HeXOr(hex1, hex2)
	return out.String(), nil
}
