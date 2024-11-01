package main

import (
	"bufio"
	"cryptography/module"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Pilih mode (1 untuk enkripsi, 2 untuk dekripsi): ")
	modeChoice, _ := reader.ReadString('\n')
	modeChoice = strings.TrimSpace(modeChoice)

	fmt.Print("Masukkan teks: ")
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSpace(inputText)

	fmt.Print("Masukkan jumlah pergeseran: ")
	inputShift, _ := reader.ReadString('\n')
	inputShift = strings.TrimSpace(inputShift)
	shift, err := strconv.Atoi(inputShift)
	if err != nil {
		fmt.Println("Jumlah pergeseran harus berupa angka!")
		return
	}

	cipher := module.CaesarCipher{
		Text:  inputText,
		Shift: shift,
	}

	if modeChoice == "1" {
		cipher.Encrypt()
		fmt.Println("Teks terenkripsi:", cipher.GetEncryptedText())
	} else if modeChoice == "2" {
		cipher.Decrypt()
		fmt.Println("Teks terdekripsi:", cipher.GetDecryptedText())
	} else {
		fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.")
	}
}
