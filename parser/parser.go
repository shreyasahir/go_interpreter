package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"go_interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curtoken  token.Token
	peektoken token.Token
}

func new(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curtoken = p.peektoken
	p.peektoken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
