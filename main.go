package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	fmt.Println(" >>> Supply Password to generate a secure Key")

	// get password from the user terminal
	// password is in  bytes
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		fmt.Println("error reading password")
		return
	}

	// we need to add salt to make our
	// derived key stronger to crack
	salt := []byte("mysaltedSalt")

	// print out the original password entered by the user
	fmt.Println("Original Password :-> ", string(password))

	// hash the original password
	// using sha-1
	hashedPassword := sha256.Sum256([]byte(password))

	fmt.Printf("Hashed Password :-> %x\n", hashedPassword)

	// generate the derived key
	// base on the input parameter
	// password to stretch
	// salt to use
	// number of iteration to use while generating the new key
	// length of the new key to be generated
	// hash function to be used while deriving the new key
	derivedKey := pbkdf2.Key(password, salt, 10, 4, sha1.New)

	// print out the derived key
	fmt.Printf("KDF Key :->  %x\n", derivedKey)

	// compute the hash function of the derived key to make it stronger
	hashDerviedKey := sha256.Sum256(derivedKey)

	// print the  hash of the derived key
	fmt.Printf("Hashed Derived Key :-> %x\n", hashDerviedKey)

}
