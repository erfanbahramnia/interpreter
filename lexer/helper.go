package lexer

import "interpreter/token"

// this method helps to create new TokenType(type and literal)
func newToken(tokenType token.TokenType, cu byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(cu)}
}

// this methods check char is a alphabet (underline is acceptable)
func isLetter(char byte) bool {
	return 'a' <= char && 'z' >= char || 'A' <= char && 'Z' >= char || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && '9' >= char
}
