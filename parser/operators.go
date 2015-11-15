package parser

import (
	"github.com/krageon/php/ast"
	"github.com/krageon/php/token"
)

type operationType int

const (
	nilOperation operationType = 1 << iota
	unaryOperation
	binaryOperation
	ternaryOperation
	assignmentOperation
	subexpressionBeginOperation
	subexpressionEndOperation
	ignoreErrorOperation
)

func operationTypeForToken(t token.Token) operationType {
	switch t {
	case token.IgnoreErrorOperator:
		return ignoreErrorOperation
	case token.UnaryOperator,
		token.NegationOperator,
		token.CastOperator,
		token.BitwiseNotOperator:
		return unaryOperation
	case token.AdditionOperator,
		token.SubtractionOperator,
		token.ConcatenationOperator,
		token.ComparisonOperator,
		token.MultOperator,
		token.AndOperator,
		token.OrOperator,
		token.AmpersandOperator,
		token.BitwiseXorOperator,
		token.BitwiseOrOperator,
		token.BitwiseShiftOperator,
		token.WrittenAndOperator,
		token.WrittenXorOperator,
		token.WrittenOrOperator,
		token.InstanceofOperator:
		return binaryOperation
	case token.TernaryOperator1:
		return ternaryOperation
	case token.AssignmentOperator:
		return assignmentOperation
	case token.OpenParen:
		return subexpressionBeginOperation
	case token.CloseParen:
		return subexpressionEndOperation
	}
	return nilOperation
}

func (p *Parser) newBinaryOperation(operator token.Item, expr1, expr2 ast.Expr) ast.Expr {
	var t ast.Type = ast.Numeric
	switch operator.Typ {
	case token.AssignmentOperator:
		return p.parseAssignmentOperation(expr1, expr2, operator)
	case token.ComparisonOperator, token.AndOperator, token.OrOperator, token.WrittenAndOperator, token.WrittenOrOperator, token.WrittenXorOperator:
		t = ast.Boolean
	case token.ConcatenationOperator:
		t = ast.String
	case token.AmpersandOperator, token.BitwiseXorOperator, token.BitwiseOrOperator, token.BitwiseShiftOperator:
		t = ast.Unknown
	}
	return ast.BinaryExpr{
		Type:       t,
		Antecedent: expr1,
		Subsequent: expr2,
		Operator:   operator.Val,
	}
}

func (p *Parser) parseBinaryOperation(lhs ast.Expr, operator token.Item, originalParenLevel int) ast.Expr {
	p.next()
	rhs := p.parseOperand()
	currentPrecedence := operatorPrecedence[operator.Typ]
	for {
		nextOperator := p.peek()
		nextPrecedence, ok := operatorPrecedence[nextOperator.Typ]
		if !ok || nextPrecedence < currentPrecedence {
			break
		}
		rhs = p.parseOperation(originalParenLevel, rhs)
	}
	return p.newBinaryOperation(operator, lhs, rhs)
}

func (p *Parser) parseTernaryOperation(lhs ast.Expr) ast.Expr {
	var truthy ast.Expr
	if p.peek().Typ == token.TernaryOperator2 {
		truthy = lhs
	} else {
		truthy = p.parseNextExpression()
	}
	p.expect(token.TernaryOperator2)
	falsy := p.parseNextExpression()
	return &ast.TernaryCallExpr{
		Condition: lhs,
		True:      truthy,
		False:     falsy,
		Type:      truthy.EvaluatesTo().Union(falsy.EvaluatesTo()),
	}
}

func (p *Parser) parseUnaryExpressionRight(operand ast.Expr, operator token.Item) ast.Expr {
	return ast.UnaryCallExpr{
		Operand:  operand,
		Operator: operator.Val,
	}
}

func (p *Parser) parseUnaryExpressionLeft(operand ast.Expr, operator token.Item) ast.Expr {
	return ast.UnaryCallExpr{
		Operand:   operand,
		Operator:  operator.Val,
		Preceding: true,
	}
}
