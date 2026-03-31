package main

import (
	"os"
	"log"
	"bufio"
)

func main() {
	// if a file is provided
	if len(os.Args) == 2 {
		var fileName = os.Args[1]
		file, err := os.ReadFile(fileName)
		if err != nil{
			log.Fatal(err)
		}
		run(string(file))
	}

	// if no file is provided
	if len(os.Args) == 1 {
		var scanner = bufio.NewScanner(os.Stdin)
		print("> ")
		for scanner.Scan() {
			var line = scanner.Text();
			if line == ".exit" {
				break;
			}
			run(line)
			print("> ")
		}
	}
}

func run(source string) {
	var tokens = scanTokens(source)
	for _, token := range tokens {
		println(token.toString())
	}
}
