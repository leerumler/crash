package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	// Checking zero values for tls.Config.NextProtos
	config := new(tls.Config)
	fmt.Println("Zero value of tls.Config:", config.NextProtos)

	// Simple way to define at least one protocol
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}
	fmt.Println("New protocol:", config.NextProtos)
}
