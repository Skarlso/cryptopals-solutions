package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: main <string>")
		os.Exit(1)
	}

	hexSrc := []byte(os.Args[1])
	hexDest := make([]byte, hex.DecodedLen(len(hexSrc)))
	hex.Decode(hexDest, hexSrc)
	dest := make([]byte, base64.StdEncoding.EncodedLen((len(hexDest))))
	base64.StdEncoding.Encode(dest, hexDest)
	fmt.Println(string(dest))
}
