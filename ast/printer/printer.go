package printer

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/krageon/php/ast"
)

type Printer struct {
	w         io.Writer
	tabLevel  int
	tabString string
}

func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		w:         w,
		tabLevel:  0,
		tabString: "\t",
	}
}

func (p *Printer) tab() {
	io.WriteString(p.w, strings.Repeat(p.tabString, p.tabLevel))
}

func (p *Printer) entab() {
	p.tabLevel += 1
}

func (p *Printer) detab() {
	p.tabLevel -= 1
}

func (p *Printer) PrintNode(node ast.Node) {
	switch n := node.(type) {
	case *ast.AnonymousFunction:
		p.PrintAnonymousFunction(n)
	case ast.AnonymousFunction:
		p.PrintAnonymousFunction(&n)
	case *ast.ArrayAppendExpr:
		p.PrintArrayAppendExpression(n)
	case ast.ArrayAppendExpr:
		p.PrintArrayAppendExpression(&n)
	case *ast.ArrayExpr:
		p.PrintArrayExpression(n)
	case ast.ArrayExpr:
		p.PrintArrayExpression(&n)
	case *ast.ArrayLookupExpr:
		p.PrintArrayLookupExpression(n)
	case ast.ArrayLookupExpr:
		p.PrintArrayLookupExpression(&n)
	case *ast.ArrayPair:
		p.PrintArrayPair(n)
	case ast.ArrayPair:
		p.PrintArrayPair(&n)
	case *ast.AssignmentExpr:
		p.PrintAssignmentExpression(n)
	case ast.AssignmentExpr:
		p.PrintAssignmentExpression(&n)
	case *ast.BinaryExpr:
		p.PrintBinaryExpression(n)
	case ast.BinaryExpr:
		p.PrintBinaryExpression(&n)
	case *ast.Block:
		p.PrintBlock(n)
	case ast.Block:
		p.PrintBlock(&n)
	case *ast.BreakStmt:
		p.PrintBreakStmt(n)
	case ast.BreakStmt:
		p.PrintBreakStmt(&n)
	case *ast.CatchStmt:
		p.PrintCatchStmt(n)
	case ast.CatchStmt:
		p.PrintCatchStmt(&n)
	case *ast.Class:
		p.PrintClass(n)
	case ast.Class:
		p.PrintClass(&n)
	case *ast.ClassExpr:
		p.PrintClassExpression(n)
	case ast.ClassExpr:
		p.PrintClassExpression(&n)
	case *ast.Constant:
		p.PrintConstant(n)
	case ast.Constant:
		p.PrintConstant(&n)
	case *ast.ConstantExpr:
		p.PrintConstantExpression(n)
	case ast.ConstantExpr:
		p.PrintConstantExpression(&n)
	case *ast.ContinueStmt:
		p.PrintContinueStmt(n)
	case ast.ContinueStmt:
		p.PrintContinueStmt(&n)
	case *ast.DeclareBlock:
		p.PrintDeclareBlock(n)
	case ast.DeclareBlock:
		p.PrintDeclareBlock(&n)
	case *ast.DoWhileStmt:
		p.PrintDoWhileStmt(n)
	case ast.DoWhileStmt:
		p.PrintDoWhileStmt(&n)
	case *ast.EchoStmt:
		p.PrintEchoStmt(n)
	case ast.EchoStmt:
		p.PrintEchoStmt(&n)
	case *ast.EmptyStatement:
		p.PrintEmptyStatement(n)
	case ast.EmptyStatement:
		p.PrintEmptyStatement(&n)
	case *ast.ExitStmt:
		p.PrintExitStmt(n)
	case ast.ExitStmt:
		p.PrintExitStmt(&n)
	case *ast.ExprStmt:
		p.PrintExpressionStmt(n)
	case ast.ExprStmt:
		p.PrintExpressionStmt(&n)
	case *ast.ForStmt:
		p.PrintForStmt(n)
	case ast.ForStmt:
		p.PrintForStmt(&n)
	case *ast.ForeachStmt:
		p.PrintForeachStmt(n)
	case ast.ForeachStmt:
		p.PrintForeachStmt(&n)
	case *ast.FunctionArgument:
		p.PrintFunctionArgument(n)
	case ast.FunctionArgument:
		p.PrintFunctionArgument(&n)
	case *ast.FunctionCallExpr:
		p.PrintFunctionCallExpression(n)
	case ast.FunctionCallExpr:
		p.PrintFunctionCallExpression(&n)
	case *ast.FunctionCallStmt:
		p.PrintFunctionCallStmt(n)
	case ast.FunctionCallStmt:
		p.PrintFunctionCallStmt(&n)
	case *ast.FunctionDefinition:
		p.PrintFunctionDefinition(n)
	case ast.FunctionDefinition:
		p.PrintFunctionDefinition(&n)
	case *ast.FunctionStmt:
		p.PrintFunctionStmt(n)
	case ast.FunctionStmt:
		p.PrintFunctionStmt(&n)
	case *ast.GlobalDeclaration:
		p.PrintGlobalDeclaration(n)
	case ast.GlobalDeclaration:
		p.PrintGlobalDeclaration(&n)
	case *ast.Identifier:
		p.PrintIdentifier(n)
	case ast.Identifier:
		p.PrintIdentifier(&n)
	case *ast.IfStmt:
		p.PrintIfStmt(n)
	case ast.IfStmt:
		p.PrintIfStmt(&n)
	case *ast.Include:
		p.PrintInclude(n)
	case ast.Include:
		p.PrintInclude(&n)
	case *ast.IncludeStmt:
		p.PrintIncludeStmt(n)
	case ast.IncludeStmt:
		p.PrintIncludeStmt(&n)
	case *ast.Interface:
		p.PrintInterface(n)
	case ast.Interface:
		p.PrintInterface(&n)
	case *ast.ListStatement:
		p.PrintListStatement(n)
	case ast.ListStatement:
		p.PrintListStatement(&n)
	case *ast.Literal:
		p.PrintLiteral(n)
	case ast.Literal:
		p.PrintLiteral(&n)
	case *ast.Method:
		p.PrintMethod(n)
	case ast.Method:
		p.PrintMethod(&n)
	case *ast.MethodCallExpr:
		p.PrintMethodCallExpression(n)
	case ast.MethodCallExpr:
		p.PrintMethodCallExpression(&n)
	case *ast.NewCallExpr:
		p.PrintNewExpression(n)
	case ast.NewCallExpr:
		p.PrintNewExpression(&n)
	case *ast.Property:
		p.PrintProperty(n)
	case ast.Property:
		p.PrintProperty(&n)
	case *ast.PropertyCallExpr:
		p.PrintPropertyExpression(n)
	case ast.PropertyCallExpr:
		p.PrintPropertyExpression(&n)
	case *ast.ReturnStmt:
		p.PrintReturnStmt(n)
	case ast.ReturnStmt:
		p.PrintReturnStmt(&n)
	case *ast.ShellCommand:
		p.PrintShellCommand(n)
	case ast.ShellCommand:
		p.PrintShellCommand(&n)
	case *ast.StaticVariableDeclaration:
		p.PrintStaticVariableDeclaration(n)
	case ast.StaticVariableDeclaration:
		p.PrintStaticVariableDeclaration(&n)
	case *ast.SwitchCase:
		p.PrintSwitchCase(n)
	case ast.SwitchCase:
		p.PrintSwitchCase(&n)
	case *ast.SwitchStmt:
		p.PrintSwitchStmt(n)
	case ast.SwitchStmt:
		p.PrintSwitchStmt(&n)
	case *ast.TernaryCallExpr:
		p.PrintTernaryExpression(n)
	case ast.TernaryCallExpr:
		p.PrintTernaryExpression(&n)
	case *ast.ThrowStmt:
		p.PrintThrowStmt(n)
	case ast.ThrowStmt:
		p.PrintThrowStmt(&n)
	case *ast.TryStmt:
		p.PrintTryStmt(n)
	case ast.TryStmt:
		p.PrintTryStmt(&n)
	case *ast.UnaryCallExpr:
		p.PrintUnaryExpression(n)
	case ast.UnaryCallExpr:
		p.PrintUnaryExpression(&n)
	case *ast.Variable:
		p.PrintVariable(n)
	case ast.Variable:
		p.PrintVariable(&n)
	case *ast.WhileStmt:
		p.PrintWhileStmt(n)
	case ast.WhileStmt:
		p.PrintWhileStmt(&n)
	default:
		fmt.Fprintf(p.w, `/* Unsupported node type: %T */`, n)
	}
}

func (p *Printer) PrintIdentifier(i *ast.Identifier) {
	io.WriteString(p.w, i.Value)
}

func (p *Printer) PrintVariable(v *ast.Variable) {
	io.WriteString(p.w, "$")
	p.PrintNode(v.Name)
}

func (p *Printer) PrintGlobalDeclaration(g *ast.GlobalDeclaration) {
	io.WriteString(p.w, "global ")
	for i, id := range g.Identifiers {
		p.PrintNode(id)
		if i+1 < len(g.Identifiers) {
			io.WriteString(p.w, ", ")
		}
	}
}

func (p *Printer) PrintEmptyStatement(e *ast.EmptyStatement) {}

func (p *Printer) PrintBinaryExpression(b *ast.BinaryExpr) {
	p.PrintNode(b.Antecedent)
	fmt.Fprintf(p.w, "%s", b.Operator)
	p.PrintNode(b.Subsequent)
}

func (p *Printer) PrintTernaryExpression(t *ast.TernaryCallExpr) {
	p.PrintNode(t.Condition)
	fmt.Fprintf(p.w, " ? ")
	p.PrintNode(t.True)
	fmt.Fprintf(p.w, " : ")
	p.PrintNode(t.False)
}

func (p *Printer) PrintUnaryExpression(u *ast.UnaryCallExpr) {
	if u.Preceding {
		fmt.Fprintf(p.w, "%s", u.Operator)
		p.PrintNode(u.Operand)
	}
	p.PrintNode(u.Operand)
	fmt.Fprintf(p.w, "%s", u.Operator)
}

func (p *Printer) PrintEchoStmt(e *ast.EchoStmt) {
	io.WriteString(p.w, "echo ")
	for i, expr := range e.Expressions {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(expr)
	}
	io.WriteString(p.w, ";")
}

func (p *Printer) PrintReturnStmt(r *ast.ReturnStmt) {
	io.WriteString(p.w, "return ")
	if r.Expr != nil {
		p.PrintNode(r.Expr)
	}
	io.WriteString(p.w, ";")
}
func (p *Printer) PrintBreakStmt(b *ast.BreakStmt) {
	io.WriteString(p.w, "break")
	if b.Expr != nil {
		p.PrintNode(b.Expr)
	}
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintContinueStmt(b *ast.ContinueStmt) {
	io.WriteString(p.w, "continue")
	if b.Expr != nil {
		p.PrintNode(b.Expr)
	}
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintThrowStmt(b *ast.ThrowStmt) {
	io.WriteString(p.w, "throw")
	if b.Expr != nil {
		p.PrintNode(b.Expr)
	}
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintInclude(e *ast.Include) {
	io.WriteString(p.w, "include ")
	for i, expr := range e.Expressions {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(expr)
	}
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintExitStmt(b *ast.ExitStmt) {
	io.WriteString(p.w, "exit")
	if b.Expr != nil {
		p.PrintNode(b.Expr)
	}
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintNewExpression(b *ast.NewCallExpr) {
	io.WriteString(p.w, "new ")
	p.PrintNode(b.Class)
	io.WriteString(p.w, "(")
	for i, arg := range b.Arguments {
		if i > 0 {
			io.WriteString(p.w, ",")
		}
		p.PrintNode(arg)
	}
	io.WriteString(p.w, ")")

}
func (p *Printer) PrintAssignmentExpression(a *ast.AssignmentExpr) {
	p.PrintNode(a.Assignee)
	io.WriteString(p.w, " ")
	io.WriteString(p.w, a.Operator)
	io.WriteString(p.w, " ")
	p.PrintNode(a.Value)
}

func (p *Printer) PrintFunctionCallStmt(f *ast.FunctionCallStmt) {
	p.PrintNode(f.FunctionCallExpr)
	io.WriteString(p.w, ";")

}
func (p *Printer) PrintFunctionCallExpression(f *ast.FunctionCallExpr) {
	p.PrintNode(f.FunctionName)
	io.WriteString(p.w, "(")
	for i, arg := range f.Arguments {
		if i > 0 {
			io.WriteString(p.w, ",")
		}
		p.PrintNode(arg)
	}
	io.WriteString(p.w, ")")

}

func (p *Printer) PrintBlock(b *ast.Block) {
	io.WriteString(p.w, "{\n")
	p.entab()
	for _, s := range b.Statements {
		p.tab()
		p.PrintNode(s)
		io.WriteString(p.w, "\n")
	}
	p.detab()
	p.tab()
	io.WriteString(p.w, "}")
}

func (p *Printer) PrintFunctionStmt(f *ast.FunctionStmt) {
	p.PrintNode(f.FunctionDefinition)
	p.PrintNode(f.Body)
}

func (p *Printer) PrintAnonymousFunction(a *ast.AnonymousFunction) {
	io.WriteString(p.w, "function (")
	for i, arg := range a.Arguments {
		if i > 0 {
			io.WriteString(p.w, ",")
		}
		p.PrintNode(arg)
	}
	io.WriteString(p.w, ") ")
	if len(a.ClosureVariables) > 0 {
		fmt.Fprint(p.w, "use (")
		for i, arg := range a.ClosureVariables {
			if i > 0 {
				io.WriteString(p.w, ",")
			}
			p.PrintNode(arg)
		}
		io.WriteString(p.w, ") ")
	}
	p.PrintNode(a.Body)
}

func (p *Printer) PrintFunctionDefinition(fd *ast.FunctionDefinition) {
	io.WriteString(p.w, "function ")
	io.WriteString(p.w, fd.Name)
	io.WriteString(p.w, "(")
	for i, arg := range fd.Arguments {
		p.PrintNode(arg)
		if i+1 < len(fd.Arguments) {
			io.WriteString(p.w, ",")
		}
	}
	io.WriteString(p.w, ") ")

}

func (p *Printer) PrintFunctionArgument(fa *ast.FunctionArgument) {
	buf := &bytes.Buffer{}
	if fa.TypeHint != "" {
		fmt.Fprint(buf, fa.TypeHint, "")
	}
	p.PrintNode(fa.Variable)
	if fa.Default != nil {
		io.WriteString(p.w, " =")
		p.PrintNode(fa.Default)
	}

}

func (p *Printer) PrintClass(c *ast.Class) {
	io.WriteString(p.w, "class ")
	io.WriteString(p.w, c.Name)
	if c.Extends != "" {
		fmt.Fprintf(p.w, " extends %s", c.Extends)
	}
	for i, imp := range c.Implements {
		if i > 0 {
			io.WriteString(p.w, ",")
		} else {
			io.WriteString(p.w, "implements ")
		}
		io.WriteString(p.w, imp)
	}
	io.WriteString(p.w, " {\n")
	p.entab()
	for _, c := range c.Constants {
		p.tab()
		p.PrintNode(c)
		io.WriteString(p.w, "\n")
	}
	for _, pr := range c.Properties {
		p.tab()
		p.PrintNode(pr)
		io.WriteString(p.w, "\n")
	}
	for _, m := range c.Methods {
		p.tab()
		p.PrintNode(m)
		io.WriteString(p.w, "\n")
	}
	p.detab()
	p.tab()
	io.WriteString(p.w, "}")
}

func (p *Printer) PrintInterface(i *ast.Interface) {
	io.WriteString(p.w, "interface ")
	io.WriteString(p.w, i.Name)

	for i, imp := range i.Inherits {
		if i > 0 {
			io.WriteString(p.w, ", ")
		} else {
			io.WriteString(p.w, "implements ")
		}
		io.WriteString(p.w, imp)
	}

	io.WriteString(p.w, " {")
	for _, c := range i.Constants {
		p.PrintNode(c)
	}

	for _, m := range i.Methods {
		p.PrintNode(m)
	}

	io.WriteString(p.w, "}")
}

func (p *Printer) PrintProperty(pr *ast.Property) {
	buf := &bytes.Buffer{}
	p.PrintVisibility(pr.Visibility)
	fmt.Fprintf(buf, " %s", pr.Name)
	if pr.Initialization != nil {
		p.PrintNode(pr.Initialization)
	}
	io.WriteString(p.w, ";")

}

func (p *Printer) PrintPropertyExpression(pr *ast.PropertyCallExpr) {
	p.PrintNode(pr.Receiver)
	io.WriteString(p.w, "->")
	p.PrintNode(pr.Name)
}

func (p *Printer) PrintClassExpression(c *ast.ClassExpr) {
	p.PrintNode(c.Receiver)
	io.WriteString(p.w, "::")
	p.PrintNode(c.Expr)
}

func (p *Printer) PrintMethod(m *ast.Method) {
	p.PrintVisibility(m.Visibility)
	io.WriteString(p.w, " ")
	p.PrintNode(m.FunctionStmt)
}

func (p *Printer) PrintMethodCallExpression(m *ast.MethodCallExpr) {
	p.PrintNode(m.Receiver)
	io.WriteString(p.w, "->")
	p.PrintNode(m.FunctionCallExpr)
}

func (p *Printer) PrintIfStmt(i *ast.IfStmt) {
	if len(i.Branches) == 0 {
		return
	}

	io.WriteString(p.w, "if (")
	p.PrintNode(i.Branches[0].Condition)
	io.WriteString(p.w, ") {\n")
	p.PrintNode(i.Branches[0].Block)
	io.WriteString(p.w, "\n}")

	for _, branch := range i.Branches[1:] {
		io.WriteString(p.w, " else if (")
		p.PrintNode(branch.Condition)
		io.WriteString(p.w, ") {\n")
		p.PrintNode(branch.Block)
		io.WriteString(p.w, "\n}")
	}
	if i.ElseBlock != nil {
		io.WriteString(p.w, " else {\n")
		p.PrintNode(i.ElseBlock)
		io.WriteString(p.w, "\n}")
	}
}

func (p *Printer) PrintSwitchStmt(s *ast.SwitchStmt) {
	fmt.Fprintf(p.w, "switch (%s) {\n", s.Expr)
	for _, c := range s.Cases {
		p.PrintNode(c)
		io.WriteString(p.w, "\n")
	}
	if s.DefaultCase != nil {
		fmt.Fprintf(p.w, "default:\n")
		p.PrintNode(s.DefaultCase)
	}
	io.WriteString(p.w, "}")
}

func (p *Printer) PrintSwitchCase(s *ast.SwitchCase) {
	io.WriteString(p.w, "case ")
	p.PrintNode(s.Expr)
	io.WriteString(p.w, ":\n")
	p.PrintNode(s.Block)
	io.WriteString(p.w, "\n")
}
func (p *Printer) PrintForStmt(f *ast.ForStmt) {
	fmt.Fprintf(p.w, "for (")
	for i, e := range f.Initialization {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(e)
	}
	for i, e := range f.Termination {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(e)
	}
	for i, e := range f.Iteration {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(e)
	}
	io.WriteString(p.w, ") ")
	p.PrintNode(f.LoopBlock)

}
func (p *Printer) PrintWhileStmt(wh *ast.WhileStmt) {
	fmt.Fprintf(p.w, "while (%s) %s", wh.Termination, wh.LoopBlock)
}
func (p *Printer) PrintDoWhileStmt(wh *ast.DoWhileStmt) {
	fmt.Fprintf(p.w, "do %s while (%s);", wh.LoopBlock, wh.Termination)
}
func (p *Printer) PrintTryStmt(t *ast.TryStmt) {
	fmt.Fprintf(p.w, "try ")
	p.PrintNode(t.TryBlock)
	for _, c := range t.CatchStmts {
		p.PrintNode(c)
	}
	if t.FinallyBlock != nil {
		fmt.Fprintf(p.w, "finally ")
		p.PrintNode(t.FinallyBlock)
	}

}
func (p *Printer) PrintCatchStmt(c *ast.CatchStmt) {
	fmt.Fprintf(p.w, "catch (%s ", c.CatchType)
	p.PrintNode(c.CatchVar)
	io.WriteString(p.w, ") ")
	p.PrintNode(c.CatchBlock)
}

func (p *Printer) PrintLiteral(l *ast.Literal) {
	switch l.Type {
	case ast.String:
		io.WriteString(p.w, l.Value)
	case ast.Integer, ast.Float:
		io.WriteString(p.w, l.Value)
	case ast.Boolean:
		io.WriteString(p.w, l.Value)
	case ast.Null:
		io.WriteString(p.w, "null")
	default:
		io.WriteString(p.w, l.Value)
	}
}

func (p *Printer) PrintForeachStmt(f *ast.ForeachStmt) {
	fmt.Fprintf(p.w, "foreach (%s as ", f.Source)
	if f.Key != nil {
		fmt.Fprintf(p.w, "%s => ", f.Key)
	}
	p.PrintNode(f.Value)
	io.WriteString(p.w, ") ")
	p.PrintNode(f.LoopBlock)
}

func (p *Printer) PrintArrayExpression(a *ast.ArrayExpr) {
	fmt.Fprintf(p.w, "array(")
	for i, pair := range a.Pairs {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(pair)
	}
	io.WriteString(p.w, ")")
}

func (p *Printer) PrintArrayPair(pr *ast.ArrayPair) {
	if pr.Key != nil {
		p.PrintNode(pr.Key)
		fmt.Fprintf(p.w, " => ")
		p.PrintNode(pr.Value)
	}
	p.PrintNode(pr.Value)
}

func (p *Printer) PrintArrayLookupExpression(a *ast.ArrayLookupExpr) {
	p.PrintNode(a.Array)
	io.WriteString(p.w, "[")
	p.PrintNode(a.Index)
	io.WriteString(p.w, "]")
}

func (p *Printer) PrintArrayAppendExpression(a *ast.ArrayAppendExpr) {
	fmt.Fprintf(p.w, "%s[]", a.Array)
}

func (p *Printer) PrintShellCommand(s *ast.ShellCommand) {
	io.WriteString(p.w, s.Command)
}

func (p *Printer) PrintListStatement(l *ast.ListStatement) {
	fmt.Fprintf(p.w, "list(")
	for i, a := range l.Assignees {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(a)
	}
	io.WriteString(p.w, ") =")
	p.PrintNode(l.Value)
}

func (p *Printer) PrintStaticVariableDeclaration(s *ast.StaticVariableDeclaration) {
	fmt.Fprintf(p.w, "static ")
	for i, d := range s.Declarations {
		if i > 0 {
			io.WriteString(p.w, ", ")
		}
		p.PrintNode(d)
	}
	io.WriteString(p.w, ";")
}
func (p *Printer) PrintDeclareBlock(d *ast.DeclareBlock) {
	io.WriteString(p.w, "declare (")
	for i, decl := range d.Declarations {
		if i > 0 {
			io.WriteString(p.w, ",")
		}
		io.WriteString(p.w, decl)
	}
	io.WriteString(p.w, ") {")
	p.PrintNode(d.Statements)
	io.WriteString(p.w, "}")
}

func (p *Printer) PrintConstant(c *ast.Constant) {
	io.WriteString(p.w, c.Name)
}

func (p *Printer) PrintConstantExpression(c *ast.ConstantExpr) {
	p.PrintNode(c.Name)
}

func (p *Printer) PrintExpressionStmt(c *ast.ExprStmt) {
	p.PrintNode(c.Expr)
	io.WriteString(p.w, ";")
}

func (p *Printer) PrintIncludeStmt(c *ast.IncludeStmt) {
	p.PrintInclude(&c.Include)
	io.WriteString(p.w, ";")
}

func (p *Printer) PrintVisibility(v ast.Visibility) {
	switch v {
	case ast.Public:
		io.WriteString(p.w, "public")
	case ast.Protected:
		io.WriteString(p.w, "protected")
	case ast.Private:
		io.WriteString(p.w, "private")
	}
}
