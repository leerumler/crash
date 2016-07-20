package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	answer := strings.TrimSuffix(input, "\n")
	return answer
}

func parsePrompt(flagname, usage, question string) string {
	var answer string
	flag.StringVar(&answer, flagname, "", usage)
	flag.Parse()
	if answer == "" {
		answer = prompt(question)
	}
	return answer
}

func main() {
	// flag testing for days
	sip := parsePrompt("sip", "Source IP address", "Source IP: ")
	fmt.Println(sip)
}
