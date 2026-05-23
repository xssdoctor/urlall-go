package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	out, err := exec.Command("pbpaste").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pbpaste failed: %v\n", err)
		os.Exit(1)
	}

	encoded := make([]byte, 0, len(out)*3)
	for _, b := range out {
		encoded = append(encoded, '%')
		encoded = append(encoded, hex(b>>4), hex(b&0x0F))
	}

	os.Stdout.Write(encoded)
}

func hex(n byte) byte {
	if n < 10 {
		return '0' + n
	}
	return 'A' + (n - 10)
}
