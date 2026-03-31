package main

import (
	"fmt"
)

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

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type token struct {
	kind    tokenKind
	lexeme  string
	literal any
	line    uint
}

func (t *token) toString() string {
	return fmt.Sprintf("token{kind: %v, lexeme: %v, literal: %v, line: %v}", t.kind, t.lexeme, t.literal, t.line)
}
