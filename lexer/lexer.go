package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSING, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // ASCII, 0 => NULL
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, cu byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(cu)}
}
