package main

import "crypto/sha512"
import "fmt"

func SHA512(message []byte) []byte {
        c := sha512.New()
        c.Write(message)
        return c.Sum(nil)
}

func main() {
	message := []byte("Hello, world.")
	digest := SHA512(message)
	fmt.Printf("SHA512('%s') = %v\n", string(message), digest)
}
