package parser

import (
	"github.com/krageon/php/ast"
	"github.com/krageon/php/lexer"
	"github.com/krageon/php/token"
)

func (p *Parser) parseFunctionStmt(inMethod bool) *ast.FunctionStmt {
	stmt := &ast.FunctionStmt{}
	stmt.FunctionDefinition = p.parseFunctionDefinition()
	if !inMethod {
		p.namespace.Functions[stmt.Name] = stmt
	}
	p.scope = ast.NewScope(p.scope, p.FileSet.GlobalScope, p.FileSet.SuperGlobalScope)
	stmt.Body = p.parseBlock()
	p.scope = p.scope.EnclosingScope
	return stmt
}

func (p *Parser) parseFunctionDefinition() *ast.FunctionDefinition {
	def := &ast.FunctionDefinition{}
	if p.peek().Typ == token.AmpersandOperator {
		// This is a function returning a reference ... ignore this for now
		p.next()
	}
	if !p.accept(token.Identifier) {
		p.next()
		if !lexer.IsKeyword(p.current.Typ, p.current.Val) {
			p.errorf("bad function name: %s", p.current.Val)
		}
	}
	def.Name = p.current.Val
	def.Arguments = make([]*ast.FunctionArgument, 0)
	p.expect(token.OpenParen)
	if p.peek().Typ == token.CloseParen {
		p.expect(token.CloseParen)
		return def
	}
	def.Arguments = append(def.Arguments, p.parseFunctionArgument())
	for {
		switch p.peek().Typ {
		case token.Comma:
			p.expect(token.Comma)
			def.Arguments = append(def.Arguments, p.parseFunctionArgument())
		case token.CloseParen:
			p.expect(token.CloseParen)
			return def
		default:
			p.errorf("unexpected argument separator: %s", p.current)
			return def
		}
	}
}

func (p *Parser) parseFunctionArgument() *ast.FunctionArgument {
	arg := &ast.FunctionArgument{}
	switch p.peek().Typ {
	case token.Identifier, token.Array, token.Self:
		p.next()
		arg.TypeHint = p.current.Val
	}
	if p.peek().Typ == token.AmpersandOperator {
		p.next()
	}
	p.expect(token.VariableOperator)
	p.next()
	arg.Variable = ast.NewVariable(p.current.Val)
	if p.peek().Typ == token.AssignmentOperator {
		p.expect(token.AssignmentOperator)
		p.next()
		arg.Default = p.parseExpression()
	}
	return arg
}

func (p *Parser) parseFunctionCall(callable ast.Expr) *ast.FunctionCallExpr {
	expr := &ast.FunctionCallExpr{}
	expr.FunctionName = callable
	return p.parseFunctionArguments(expr)
}

func (p *Parser) parseFunctionArguments(expr *ast.FunctionCallExpr) *ast.FunctionCallExpr {
	expr.Arguments = make([]ast.Expr, 0)
	p.expect(token.OpenParen)
	if p.peek().Typ == token.CloseParen {
		p.expect(token.CloseParen)
		return expr
	}
	expr.Arguments = append(expr.Arguments, p.parseNextExpression())
	for p.peek().Typ != token.CloseParen {
		p.expect(token.Comma)
		arg := p.parseNextExpression()
		if arg == nil {
			break
		}
		expr.Arguments = append(expr.Arguments, arg)
	}
	p.expect(token.CloseParen)
	return expr

}

func (p *Parser) parseAnonymousFunction() ast.Expr {
	f := &ast.AnonymousFunction{}
	f.Arguments = make([]*ast.FunctionArgument, 0)
	f.ClosureVariables = make([]*ast.FunctionArgument, 0)
	p.expect(token.OpenParen)
	if p.peek().Typ != token.CloseParen {
		f.Arguments = append(f.Arguments, p.parseFunctionArgument())
	}

Loop:
	for {
		switch p.peek().Typ {
		case token.Comma:
			p.expect(token.Comma)
			f.Arguments = append(f.Arguments, p.parseFunctionArgument())
		case token.CloseParen:
			break Loop
		default:
			p.errorf("unexpected argument separator: %s", p.current)
			return f
		}
	}
	p.expect(token.CloseParen)

	// Closure variables
	if p.peek().Typ == token.Use {
		p.expect(token.Use)
		p.expect(token.OpenParen)
		f.ClosureVariables = append(f.ClosureVariables, p.parseFunctionArgument())
	ClosureLoop:
		for {
			switch p.peek().Typ {
			case token.Comma:
				p.expect(token.Comma)
				f.ClosureVariables = append(f.ClosureVariables, p.parseFunctionArgument())
			case token.CloseParen:
				break ClosureLoop
			default:
				p.errorf("unexpected argument separator: %s", p.current)
				return f
			}
		}
		p.expect(token.CloseParen)
	}

	p.scope = ast.NewScope(p.scope, p.FileSet.GlobalScope, p.FileSet.SuperGlobalScope)
	f.Body = p.parseBlock()
	p.scope = p.scope.EnclosingScope
	return f
}
