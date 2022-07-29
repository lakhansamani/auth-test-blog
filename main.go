package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	storedChallenge := `Y4kZcRUTVPzyVZMclRB0KTDp1J9Mf5WBVHN_U2VSVTs`
	storedCode := `hello_world`
	fmt.Println("challenge:", storedChallenge)
	fmt.Println("code:", storedCode)
	verifier := `jHO1VPYVNiaA9M1Mcd0Ide2OeQ~DHp8uWVWfMx3Xisl`
	fmt.Println("verifier:", verifier)

	// create hash for verifier
	hash := sha256.New()
	n, err := hash.Write([]byte(verifier))
	fmt.Println("n:", n)
	fmt.Println("err:", err)
	encryptedCode := strings.ReplaceAll(base64.URLEncoding.EncodeToString(hash.Sum(nil)), "+", "-")
	encryptedCode = strings.ReplaceAll(encryptedCode, "/", "_")
	encryptedCode = strings.ReplaceAll(encryptedCode, "=", "")
	fmt.Println(storedChallenge == encryptedCode)
}
