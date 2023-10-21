// Simple Lexer for the Monkey Programming Language
// Takes the source code and returns corresponding tokens
// Lexer only supports ASCII characters
package lexer

import "github.com/dansch87/monkey/token"

type Lexer struct {
	input        string		// source code
	position     int 		// current index of position at input string (points to current char)
	readPosition int 		// current index of reading position in input (points one step ahead after current char)
	ch           byte		// character
}

// Constructor function that creates a new Lexer object in memory and returns a pointer to that lexer object
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // requried to set the Lexer attributes .position and .readPosition
	return l
}

// readChar reads the character out of the source code where the Lexer.readPosition is currently pointing at.
// In case that the end of the source code is already reached, the NUL character is stored into Lexer attribute .ch.
// After the character is stored, the current Lexer.position is updated with the readPosition and then the readPosition 
// gets advanced by 1.
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

// newToken is a helper function that creates and returns a tokenType object 
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter is a helper function that checks if a character is a letter
// If character is a letter, true is returned, elsewise false
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit is a helper function that return true when a character is a digit, else it returns false
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readIdentifier Lexer method reads identfiers as consecutive characters and returns it as a string
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber Lexer method reads consecutive numeric characters and returns it as a string
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace Lexer method is a helper function that is used to skip white space as well as tab, carriage return and newline escape sequences
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// peekChar Lexer method returns the character that is one position ahead of time as byte
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// NextToken is a Lexer method that reads the literal from soruce code and identifies the corresponding token type
// Based on the token type and literal, a Token object is returned 
func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	// Skip white space
	l.skipWhitespace()

	// Create a Token object
	switch l.ch {
	case '=':
		// Check for equal operator '=='
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// Check for unequal operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}		
	case '/':
		tok = newToken(token.SLASH, l.ch)		
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
	// If NUL character then EOF token
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Ability to read identifiers (variables, keywords) or numeric values
		// If not possible ILLEGAL token is returned
		if isLetter(l.ch) {
			// read consecutive letters as literal
			tok.Literal = l.readIdentifier()
			// return keyword Token or IDENT token
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Read char on next position and update the attributes .position and .readPosition of Lexer obj
	l.readChar()

	return tok
}
