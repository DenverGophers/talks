package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"io"
)

func SHA512Reader(r io.Reader) []byte {
	// Not checking err for brevity. Do check err in real code.
	c := sha512.New()

	for {
		buf := make([]byte, sha512.BlockSize)

		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return []byte{}
		}
		c.Write(buf[:n])
		if err == io.EOF {
			err = nil
			break
		}
	}
	return c.Sum(nil)
}

func main() {
	message := []byte("Hello, world.")
	buf := bytes.NewBuffer(message)
	digest := SHA512Reader(buf)
	fmt.Printf("SHA512('%s') = %v\n", string(message), digest)
}

