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
	var err error
	// ignore whitespace in between tokens
	// increment l.line at newlines
	for {
		nextCharacter, err = l.peekChar()
		if err != nil {
			//TODO: write a handler for out of range access
			newToken = token{kind: EOF, lexeme: "", line: l.line, literal: nil}
			return newToken
		}
		if nextCharacter == '\n' {
			l.line++
		}

		if nextCharacter == '\n' || nextCharacter == ' ' || nextCharacter == '\t'{
			l.readChar() // read the character but ignore it
			continue
		} else {
			break
		}
	}

	switch nextCharacter {
	// scanning single character tokens
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
	case '/':
		{
			char, _ := l.readChar()
			newToken = token{kind: SLASH, lexeme: string(char), line: l.line, literal: nil}
		}

	case '*': 
		{
			char, _ := l.readChar()
			newToken = token{kind: STAR, lexeme: string(char), line: l.line, literal: nil}
		}
	case ';':
		{
			char, _ := l.readChar()
			newToken = token{kind: SEMICOLON, lexeme: string(char), line: l.line, literal: nil}
		}
	// scanning double character tokens
	case '!': // !=, ! operators
		{
			char, _ := l.readChar() // read the character so we can peek the next
			nextChar, err := l.peekChar()
			// TODO: handle peekChar error
			if err != nil || nextChar != '=' {
				// assuming the error indicates end of input, ie there's no more characters after char
				newToken = token{kind: BANG, lexeme: string(char), line: l.line, literal: nil}
			} else {
				var lexeme_bytes = make([]byte, 0, 2)
				lexeme_bytes = append(lexeme_bytes, char)
				char, _ = l.readChar()
				lexeme_bytes = append(lexeme_bytes, char)
				newToken = token{kind: BANG_EQUAL, lexeme: string(lexeme_bytes), line: l.line, literal: nil}
			}
			return newToken
		}

	case '=': // ==, = operators
		{
			char, _ := l.readChar() // read the character so we can peek the next
			nextChar, err := l.peekChar()
			// TODO: handle peekChar error
			if err != nil || nextChar != '=' {
				// assuming the error indicates end of input, ie there's no more characters after char
				newToken = token{kind: EQUAL, lexeme: string(char), line: l.line, literal: nil}
			} else {
				var lexeme_bytes = make([]byte, 0, 2)
				lexeme_bytes = append(lexeme_bytes, char)
				char, _ = l.readChar()
				lexeme_bytes = append(lexeme_bytes, char)
				newToken = token{kind: EQUAL_EQUAL, lexeme: string(lexeme_bytes), line: l.line, literal: nil}
			}
			return newToken
		}

	case '<': // <=, < operators
		{
			char, _ := l.readChar() // read the character so we can peek the next
			nextChar, err := l.peekChar()
			// TODO: handle peekChar error
			if err != nil || nextChar != '=' {
				// assuming the error indicates end of input, ie there's no more characters after char
				newToken = token{kind: LESS, lexeme: string(char), line: l.line, literal: nil}
			} else {
				var lexeme_bytes = make([]byte, 0, 2)
				lexeme_bytes = append(lexeme_bytes, char)
				char, _ = l.readChar()
				lexeme_bytes = append(lexeme_bytes, char)
				newToken = token{kind: LESS_EQUAL, lexeme: string(lexeme_bytes), line: l.line, literal: nil}
			}
			return newToken
		}

	case '>': // >=, > operators
		{
			char, _ := l.readChar() // read the character so we can peek the next
			nextChar, err := l.peekChar()
			// TODO: handle peekChar error
			if err != nil || nextChar != '=' {
				// assuming the error indicates end of input, ie there's no more characters after char
				newToken = token{kind: GREATER, lexeme: string(char), line: l.line, literal: nil}
			} else {
				var lexeme_bytes = make([]byte, 0, 2)
				lexeme_bytes = append(lexeme_bytes, char)
				char, _ = l.readChar()
				lexeme_bytes = append(lexeme_bytes, char)
				newToken = token{kind: GREATER_EQUAL, lexeme: string(lexeme_bytes), line: l.line, literal: nil}
			}
			return newToken
		}

	default:
		{
			char, _ := l.readChar()
			newToken = token{kind: ILLEGAL, lexeme: string(char), line: l.line, literal: nil}
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
// remove all comments (lines starting with '//')
func cleanSrc(src []byte) []byte {
	var cleanSrc = make([]byte, 0, len(src))
	var counter uint32 = 0
	var inComment bool = false; // indicates when we are in a comment line
	for i, b := range src {
		//NOTE: we start by eliminating comments before eliminating carriage returns - comments contain carriage returns anyway
		if b == '/' {
			if !((i + 1) >= len(src)) && src[i + 1] == '/' { // so we know we are on a comment - we skip all characters upto a new line
				inComment = true;
			}
		}
		if inComment {
			if !((i + 1) >= len(src)) && src[i + 1] == '\n' { // the last newline should be left in the source code to mark the line
				inComment = false; // the next character will not be in the comment line
			}
			continue
		}

		if b == '\r' {
			continue // skip '\r' bytes
		}
		cleanSrc = append(cleanSrc, b)
		counter++
	}

	return cleanSrc
}

func scanTokens(source string) []token {
	var tokens = []token{}
	var lex = lexer{}
	lex.source = cleanSrc([]byte(source))
	lex.sourceLength = uint(len(lex.source))
	lex.line = 1

	for {
		var newToken = lex.scanToken()
		tokens = append(tokens, newToken)
		if newToken.kind == EOF {
			break
		}
	}

	return tokens
}
