// Simple Lexer for the Monkey Programming Language
// Lexer only supports ASCII characters
package lexer

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