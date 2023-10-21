// Package defines the basic data structure and helper functions that are required for the translation 
// of monkey source code into Tokens that is conducted during the lexical analysis (lexer).
package token

type TokenType string

type Token struct {
	Type    TokenType		// e.g. "ASSIGN"  	or "EQ" 
	Literal string			// e.g. "=" 		or "=="
}

// Definition and assignment of TokenType to constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 134321

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT = "<"
	GT = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Initialize hash table that contains a mapping between identifier strings and there 
// corresponding TokenType (Token).
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent receives an identifier string (e.g. "fn", "let", etc.) and lookup and returns 
// it's appropriate TokenType from a hash table. If there occures no match at all, the default 
// TokenType IDENT is returned.
func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
