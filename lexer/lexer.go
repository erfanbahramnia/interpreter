package lexer

import (
	"fmt"
	"interpreter/token"
)

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

// this methods help to identify the value and type of token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	default:
		if isLetter(l.currentChar) {
			// get keyword or identifier
			tok.Literal = l.readIdentifier()
			fmt.Println(tok.Literal)
			// check token is keyword or identifier
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.currentChar)
	}

	l.readChar()
	return tok
}

// this method helps to move char by char
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // ASCII, 0 => NULL
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// this method helps to read number corectly without break it
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// this method help to get word instead of char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// remove all space (include: /n. /r, /t)
func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\r' || l.currentChar == '\t' {
		l.readChar()
	}
}
