package main

import (
	"fmt"

	"github.com/k1jo/cryptit/decrypt"
	"github.com/k1jo/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Grunwald")
	fmt.Println(encryptedStr)
	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println(decryptedStr)

}
