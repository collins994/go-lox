package main

import "fmt"

func main() {
	var source = "(){},.-+;/##\t*"
	var scanner = newScanner(source)
	var counter = 0

	for {
		counter++
		var nextToken = scanner.next()
		scanner.print_token(nextToken)
		if nextToken.kind == EOF {
			break
		}
	}
	fmt.Printf("number of tokens: %v\n", counter)
}

// func print_token(tok token) {
// 	fmt.Printf("token{kind: %v, start:%v, end:%v}\n", tok.kind, tok.start, tok.end)
// }

// import (
// 	"os"
// 	"log"
// 	"bufio"
// )
//
// func main() {
// 	// if a file is provided
// 	if len(os.Args) == 2 {
// 		var fileName = os.Args[1]
// 		file, err := os.ReadFile(fileName)
// 		if err != nil{
// 			log.Fatal(err)
// 		}
// 		run(string(file))
// 	}
//
// 	// if no file is provided
// 	if len(os.Args) == 1 {
// 		var scanner = bufio.NewScanner(os.Stdin)
// 		print("> ")
// 		for scanner.Scan() {
// 			var line = scanner.Text();
// 			if line == ".exit" {
// 				break;
// 			}
// 			run(line)
// 			print("> ")
// 		}
// 	}
// }
//
// func run(source string) {
// 	var tokens = scanTokens(source)
// 	for _, token := range tokens {
// 		if token.kind != EOF {
// 			println(token.toString())
// 		}
// 	}
// }
