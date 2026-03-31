package main

import (
	"errors"
)

type lexer struct {
	current      uint // index of the character being scanned
	line         uint // the line of the source code where the current character is
	source       []byte
	sourceLength uint
}

/** =========================================== LEXER METHODS ======================== */
// returns the next token in the source code
// returns EOF token at the end of input
func (l *lexer) scanToken() token {
	var newToken = token{}
	var nextCharacter byte
	// TODO: skip any newlines in between tokens
	nextCharacter, err := l.peekChar()
	if err != nil {
		//TODO: write a handler for out of range access
		newToken = token{kind: EOF, lexeme: "", line: l.line, literal: nil}
		return newToken
	}

	switch nextCharacter {
	case '(':
		{
			char, _ := l.readChar()
			newToken = token{kind: LEFT_PAREN, lexeme: string(char), line: l.line, literal: nil}
		}
	case ')':
		{
			char, _ := l.readChar()
			newToken = token{kind: RIGHT_PAREN, lexeme: string(char), line: l.line, literal: nil}
		}
	case '{':
		{
			char, _ := l.readChar()
			newToken = token{kind: LEFT_BRACE, lexeme: string(char), line: l.line, literal: nil}
		}
	case '}':
		{
			char, _ := l.readChar()
			newToken = token{kind: RIGHT_BRACE, lexeme: string(char), line: l.line, literal: nil}
		}
	case ',':
		{
			char, _ := l.readChar()
			newToken = token{kind: COMMA, lexeme: string(char), line: l.line, literal: nil}
		}
	case '.':
		{
			char, _ := l.readChar()
			newToken = token{kind: DOT, lexeme: string(char), line: l.line, literal: nil}
		}
	case '-':
		{
			char, _ := l.readChar()
			newToken = token{kind: MINUS, lexeme: string(char), line: l.line, literal: nil}
		}
	case '+':
		{
			char, _ := l.readChar()
			newToken = token{kind: PLUS, lexeme: string(char), line: l.line, literal: nil}
		}
	case ';':
		{
			char, _ := l.readChar()
			newToken = token{kind: SEMICOLON, lexeme: string(char), line: l.line, literal: nil}
		}
	default:
		{
			char, _ := l.readChar()
			newToken = token{kind: STAR, lexeme: string(char), line: l.line, literal: nil}
		}
	}

	return newToken
}

// returns the next character in the source code
// it returns an error if there is not enough characters in the source code
func (l *lexer) readChar() (byte, error) {
	if (l.current) >= l.sourceLength {
		// TODO: define an error for  end of input
		return byte(0), errors.New("not enough characters")
	}
	var b = l.source[l.current]
	l.current++ // the next call to readChar will return the next character
	return b, nil
}

// peekChar returns the character that the next call to readChar is gonna return
// it returns an error if there's not enough character in the source code
// eg source == "helo", current == 0, peekChar(0) => 'h'
// unless readChar is called, peekChar will always return the same character, i.e it does not change the value of l.current
func (l *lexer) peekChar() (byte, error) {
	if l.current >= l.sourceLength {
		// TODO: define an error for out of range access
		return byte(0), errors.New("not enough characters")
	}
	return l.source[l.current], nil
}

/** ============================== UTILITY FUNCTIONS ============================*/
// removes all \r characters
func cleanSrc(src []byte) []byte {
	var cleanSrc = make([]byte, len(src), len(src))
	var counter uint32 = 0
	for _, b := range src {
		if b == '\r' {
			continue // skip '\r' bytes
		}
		cleanSrc[counter] = b
		counter++
	}

	return cleanSrc
}

func scanTokens(source string) []token {
	var tokens = []token{}
	var lex = lexer{}
	lex.source = cleanSrc([]byte(source))
	lex.sourceLength = uint(len(lex.source))

	for {
		var newToken = lex.scanToken()
		tokens = append(tokens, newToken)
		if newToken.kind == EOF {
			break
		}
	}

	return tokens
}
