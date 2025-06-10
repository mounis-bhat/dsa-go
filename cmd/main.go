package main

import (
	"fmt"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func main() {
	isValid := neetcode150.ValidAnagram("art", "rat")
	fmt.Println(isValid)
}
