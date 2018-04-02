package set1

import (
	"bufio"
	"io"
	"os"
)

// DecryptFromFile - return lines you may have been able to decrypt, or error if can't open file
func DecryptFromFile(filename string) (decrypted []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var line string
	scanner := bufio.NewReader(file)
	for readErr := error(nil); readErr != io.EOF; {
		line, readErr = scanner.ReadString('\n')
		if len(line) > 0 && line[len(line)-1] == '\n' {
			line = line[:len(line)-1] // ignore '\n'
		}

		hex, err := GetHexFromString(line)
		if err != nil {
			return nil, err
		}

		plainMap := Decrypt(hex)

		for _, plainText := range plainMap {
			decrypted = append(decrypted, string(plainText))
		}
	}

	return
}
