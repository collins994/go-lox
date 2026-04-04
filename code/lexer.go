package main

import (
	"errors"
	"strconv"
	"fmt"
)

type lexer struct {
	current      uint // index of the character being scanned
	line         uint // the line of the source code where the current character is
	source       []byte
	sourceLength uint
	hadError     bool // indicates whether or not the lexer encountered an error during scanning
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
		nextCharacter, err = l.peek()
		if err != nil {
			//end of input
			newToken = token{kind: EOF, lexeme: "", line: l.line, literal: nil}
			return newToken
		}
		if nextCharacter == '\n' {
			l.line++
		}

		if nextCharacter == '\n' || nextCharacter == ' ' || nextCharacter == '\t' {
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
			nextChar, err := l.peek()
			// TODO: handle peek error
			if err != nil || nextChar != '=' {
				// The error indicates end of input, ie there's no more characters after char
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
			nextChar, err := l.peek()
			// TODO: handle peek error
			if err != nil || nextChar != '=' {
				// The error indicates end of input, ie there's no more characters after char
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
			nextChar, err := l.peek()
			// TODO: handle peek error
			if err != nil || nextChar != '=' {
				// The error indicates end of input, ie there's no more characters after char
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
			nextChar, err := l.peek()
			// TODO: handle peek error
			if err != nil || nextChar != '=' {
				// The error indicates end of input, ie there's no more characters after char
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
	

	// scan strings
	case '"': {
		// read untill we get the closing " or until the end of input (in which case we stop reading, report an error, then return an ILLEGAL token)
		// since Lox supports newlines in strings, we need to update l.line at newlines
		var string_bytes []byte;
		var nextChar byte;
		var char byte;
		var startLine = l.line; // mark the line where the string begins (for error reporting)
		l.readChar(); // discard the opening "
		for {
			nextChar, err = l.peek();
			if err != nil { // err indicates end of input (before the closing ")
				l.reportError("Unterminated string", startLine);
				newToken = token{kind: ILLEGAL, lexeme: string(string_bytes), line: l.line, literal: nil}
				break;
			}
			if nextChar == '\n' {
				l.line++;
			}
			if nextChar != '"' { // only append up to the last character before the closing "
				char, _ = l.readChar() 
				string_bytes  = append(string_bytes, char);
			} else {
				newToken = token{kind: STRING, lexeme: string(string_bytes), line: l.line, literal: string(string_bytes)}
				l.readChar(); // discard the closing " before breaking
				break;
			}
		}
	}


	// scan numbers
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		{
			var number_bytes []byte;
			var ok bool = true; // false iff the number lexeme has an error
			// we read until we hit the end of the number (a whitespace, end of input, or a non digit character)
			for {
				nextChar, err := l.peek();
				if err != nil { // end of input
					break;
				}

				if !(nextChar >= '0' && nextChar <= '9') { // if nextChar is not a number
					if nextChar != '.' { // we've hit the end of the number
						break;
					}

					if nextChar == '.' {
						char, _ := l.readChar(); // read the character so we can peek the next
						number_bytes = append(number_bytes, char);
						// check if there is a character after the . and that it is a digit. if it is not - that is a trailing . error
						// report the error and declare the token ILLEGAL
						nextChar, err = l.peek(); 
						if err != nil || !(nextChar >= '0' && nextChar <= '9') { 
							l.reportError("Trailing .", l.line);
							ok = false;
							break
						}
					}
				}

				// the character is a digit - just read it 
				char, _ := l.readChar();
				number_bytes = append(number_bytes, char);
			}

			if ok {
			 number, _ := strconv.ParseFloat(string(number_bytes), 64);
			 newToken = token{kind: NUMBER, lexeme: string(number_bytes), line: l.line, literal: number}
			}
			if !ok {
			 newToken = token{kind: ILLEGAL, lexeme: string(number_bytes), line: l.line, literal: nil}
			}
		}

	default:
		{
			char, _ := l.readChar()
			l.reportError(fmt.Sprintf("Unexpected character %c", char), l.line)
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

// peek returns the character that the next call to readChar is gonna return
// it returns an error if there's not enough character in the source code
// eg source == "helo", current == 0, peek() => 'h'
// unless readChar is called, peek will always return the same character
func (l *lexer) peek() (byte, error) {
	if l.current >= l.sourceLength {
		// TODO: define an error for out of range access
		return byte(0), errors.New("not enough characters")
	}
	return l.source[l.current], nil
}


// peekNext returns the character that the second call to readChar is gonna return
// it returns an error if there's not enough character in the source code
// eg source == "helo", current == 0, peekNext() => 'e'
// unless readChar is called, peekNext will always return the same character
func (l *lexer) peekNext() (byte, error) {
	if (l.current + 1) >= l.sourceLength {
		// TODO: define an error for out of range access
		return byte(0), errors.New("not enough characters")
	}
	return l.source[l.current + 1], nil
}

// reports an error to the user
func (l *lexer) reportError(message string, line uint) {
	fmt.Printf("[line %v] Error: %v\n", line, message)
	l.hadError = true
}

/** ============================== UTILITY FUNCTIONS ============================*/
// removes all \r characters
// remove all comments (lines starting with '//')
func cleanSrc(src []byte) []byte {
	var cleanSrc = make([]byte, 0, len(src))
	var counter uint32 = 0
	var inComment bool = false // indicates when we are in a comment line
	for i, b := range src {
		//NOTE: we start by eliminating comments before eliminating carriage returns - comments contain carriage returns anyway
		if b == '/' {
			if !((i + 1) >= len(src)) && src[i+1] == '/' { // so we know we are on a comment - we skip all characters upto a new line
				inComment = true
			}
		}
		if inComment {
			if !((i + 1) >= len(src)) && src[i+1] == '\n' { // the last newline should be left in the source code to mark the line
				inComment = false // the next character will not be in the comment line
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
		if newToken.kind == ILLEGAL {
			continue // skip illegal tokens
		}
		tokens = append(tokens, newToken)
		if newToken.kind == EOF {
			break
		}
	}

	return tokens
}
