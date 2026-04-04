package main

import (
	"testing"
)

func TestCleanSrc(t *testing.T) {
	var tests = []struct {
		source         string
		bytes_expected []byte
	}{
		{source: "hell\r\no\r\n", bytes_expected: []byte{'h', 'e', 'l', 'l', '\n', 'o', '\n'}},
		{
			source: "hell//this is a comment \r\nco//and this is a comment", 
			bytes_expected: []byte{'h', 'e', 'l', 'l', '\n', 'c', 'o'},
		},
	}

	for testNumber, test := range tests {
		var bytes_got = cleanSrc([]byte(test.source))
		if len(bytes_got) != len(test.bytes_expected) {
			t.Fatalf("TEST NUMBER %v: expected %v bytes, got %v bytes", testNumber, len(test.bytes_expected), len(bytes_got))
		}

		for counter := 0; counter < len(test.bytes_expected); counter++ {
			var byte_expected = test.bytes_expected[counter]
			var byte_got = bytes_got[counter]

			if byte_expected != byte_got {
				t.Fatalf("TEST NUMBER %v: expected %c , got %c", testNumber, byte_expected, byte_got)
			}
		}
	}
}

func TestPeekChar(t *testing.T) {
	var source string = "var hello"
	var lex = lexer{source: cleanSrc([]byte(source))}
	lex.sourceLength = uint(len(lex.source))

	var tests = []struct {
		current       uint
		byte_expected byte
	}{
		{current: 0, byte_expected: 'v'},
		{current: 1, byte_expected: 'a'},
		{current: 1, byte_expected: 'a'},
	}

	for testNumber, test := range tests {
		lex.current = test.current
		byte_got, err := lex.peek()
		if err != nil {
			t.Fatalf("peekChar error: test number: #%v, err: %v\n", testNumber, err)
		}

		if byte_got != test.byte_expected {
			t.Fatalf("TEST NUMBER %v: expected byte, %v, got byte %v", testNumber, test.byte_expected, byte_got)
		}
	}
}

func TestReadChar(t *testing.T) {
	var source string = "var hello"
	var lex = lexer{source: cleanSrc([]byte(source))}
	lex.sourceLength = uint(len(lex.source))

	var bytes_expected = []byte{'v', 'a', 'r'}
	
	for testNumber, byte_expected := range bytes_expected {
		byte_got, err := lex.readChar();
		if err != nil {
			t.Fatalf("readChar error: test number: #%v, err: %v\n", testNumber, err)
		}

		if byte_got != byte_expected {
			t.Fatalf("TEST NUMBER %v: expected byte, %v, got byte %v", testNumber, byte_expected, byte_got)
		}
	}
}

func TestScanTokens (t *testing.T) {
	var source = "(){},.-+;!!====<<=>//this is a comment \r\n>=/*\"this is a string\" 88.8"
	var tokens_got = scanTokens(source)
	var tokens_expected = []token{
		token{kind: LEFT_PAREN, lexeme: "(", line: 1, literal: nil},
		token{kind: RIGHT_PAREN, lexeme: ")", line: 1, literal: nil},
		token{kind: LEFT_BRACE, lexeme: "{", line: 1, literal: nil},
		token{kind: RIGHT_BRACE, lexeme: "}", line: 1, literal: nil},
		token{kind: COMMA, lexeme: ",", line: 1, literal: nil},
		token{kind: DOT, lexeme: ".", line: 1, literal: nil},
		token{kind: MINUS, lexeme: "-", line: 1, literal: nil},
		token{kind: PLUS, lexeme: "+", line: 1, literal: nil},
		token{kind: SEMICOLON, lexeme: ";", line: 1, literal: nil},
		token{kind: BANG, lexeme: "!", line: 1, literal: nil},
		token{kind: BANG_EQUAL, lexeme: "!=", line: 1, literal: nil},
		token{kind: EQUAL_EQUAL, lexeme: "==", line: 1, literal: nil},
		token{kind: EQUAL, lexeme: "=", line: 1, literal: nil},
		token{kind: LESS, lexeme: "<", line: 1, literal: nil},
		token{kind: LESS_EQUAL, lexeme: "<=", line: 1, literal: nil},
		token{kind: GREATER, lexeme: ">", line: 1, literal: nil},
		token{kind: GREATER_EQUAL, lexeme: ">=", line: 2, literal: nil},
		token{kind: SLASH, lexeme: "/", line: 2, literal: nil},
		token{kind: STAR, lexeme: "*", line: 2, literal: nil},
		token{kind: STRING, lexeme: "this is a string", line: 2, literal: "this is a string"},
		token{kind: NUMBER, lexeme: "88.8", line: 2, literal: 88.8},
		token{kind: EOF, lexeme: "", line: 2, literal: nil},
	}

	if len(tokens_got) != len(tokens_expected) {
		t.Fatalf("expected %v tokens, got %v tokens", len(tokens_expected), len(tokens_got))
	}

	for counter, token_expected := range tokens_expected {
		var token_got = tokens_got[counter]
		if token_got != token_expected {
			t.Fatalf("TOKEN NUMBER #%v: expected %v, got %v", counter, token_expected.toString(), token_got.toString())
		}
	}
}
