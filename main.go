package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const (
	characters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{};':\"\\|,.<>/?"
	defaultLength    = 32
	generationPrompt = "Press enter to generate password: "
)

func generatePassword(length int) string {
	characterSetLen := big.NewInt(int64(len(characters)))
	password := make([]rune, length)

	for i := range password {
		index, err := rand.Int(rand.Reader, characterSetLen)
		if err != nil {
			panic(err)
		}
		password[i] = rune(characters[index.Int64()])
	}

	return string(password)
}

func main() {
	for {
		fmt.Print(generationPrompt)
		fmt.Scanln()
		password := generatePassword(defaultLength)
		fmt.Println(password)
		time.Sleep(time.Millisecond)
	}
}
