// Simple Lexer for the Monkey Programming Language
// Lexer only supports ASCII characters
package lexer

import "github.com/dansch87/monkey/token"

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// give next character and advance position in the input string

	if l.readPosition >= len(l.input) {
		l.ch = 0 // if last index of input is reached, ch is set to ASCII code for NUL character
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	// Actual creation of Token object based on defined const
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	// Create a Token object
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// Read char on next position and update Lexer obj
	l.readChar()

	return tok
}
