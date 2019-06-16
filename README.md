## KDF Implementation
Implementation of Key Derivation Function (KDF) algorithm for password stretching in GoLang 

### Package Used
- `crypto/sha1` : Used to compute the hash inside the Key function inside the kdf key generation.
- `crypto/sha256` : Used to generate 32 bytes of the input data
- `github.com/howeyc/gopass` : Used to get the password from the user in the terminal.
- `golang.org/x/crypto/pbkdf2` : Used to generate the derived key.

### Clone The Repository
From your terminal, changen to the directory of your choice and run 
`git clone https://github.com/princesegzy01/kdf-implementation.git` 

### Installing Pluging
` go get github.com/howeyc/gopass`

` go get golang.org/x/crypto/pbkdf2`


### Run The App
To run the application, Execute the below command from your terminal.

`go run main.go`