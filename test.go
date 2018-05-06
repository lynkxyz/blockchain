package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	x := sha256.Sum256([]byte("I like donutsca07ca"))
	fmt.Printf("%x\n", x)
}