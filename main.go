package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	fmt.Println(" >>> Supply Password to generate a secure Key")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		fmt.Println("error reading password")
		return
	}

	salt := []byte("mysaltedSalt")

	fmt.Println("Original Password :-> ", string(password))
	derivedKey := pbkdf2.Key(password, salt, 10, 32, sha1.New)
	fmt.Printf("KDF Key :->  %x\n", derivedKey)
	hashDerviedKey := sha1.Sum(derivedKey)
	fmt.Printf("Hashed Derived Key :-> %x\n", hashDerviedKey)

	hashedPassword := sha1.Sum([]byte(password))

	fmt.Printf("Hashed Password :-> %x\n", hashedPassword)
}
