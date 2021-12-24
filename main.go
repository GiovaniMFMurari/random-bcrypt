package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// props to https://stackoverflow.com/users/141555/sberry
func rangeIn(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}

func main() {
	var pass = fmt.Sprintf("%08d", rangeIn(0, 99999999))
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Password:", pass)
	fmt.Println("Hash:", string(hash))
}
