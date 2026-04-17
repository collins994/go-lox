package main

type token struct {
	kind  tokenKind
	start uint // an index into the source code for the first character of the token
	end   uint // an index into the source code for the last character of the token
}

type tokenKind int

const (
	// single character tokens
	LEFT_PAREN tokenKind = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	ILLEGAL
	EOF
)

// import (
// 	"fmt"
// )
//
// type tokenKind int
//
// const (
// 	// single character tokens
// 	LEFT_PAREN tokenKind = iota
// 	RIGHT_PAREN
// 	LEFT_BRACE
// 	RIGHT_BRACE
// 	COMMA
// 	DOT
// 	MINUS
// 	PLUS
// 	SEMICOLON
// 	SLASH
// 	STAR
//
// 	// One or two character tokens.
// 	BANG
// 	BANG_EQUAL
// 	EQUAL
// 	EQUAL_EQUAL
// 	GREATER
// 	GREATER_EQUAL
// 	LESS
// 	LESS_EQUAL
//
// 	// literals
// 	IDENTIFIER
// 	STRING
// 	NUMBER
//
// 	// keywords
// 	AND
// 	CLASS
// 	ELSE
// 	FALSE
// 	FUN
// 	FOR
// 	IF
// 	NIL
// 	OR
// 	PRINT
// 	RETURN
// 	SUPER
// 	THIS
// 	TRUE
// 	VAR
// 	WHILE
//
// 	ILLEGAL
// 	EOF
// )
//
// type token struct {
// 	kind    tokenKind
// 	lexeme  string
// 	literal any
// 	line    uint
// }
//
// func (t *token) toString() string {
// 	return fmt.Sprintf("token{kind: %v, lexeme: %v, literal: %v, line: %v}", t.kind, t.lexeme, t.literal, t.line)
// }
