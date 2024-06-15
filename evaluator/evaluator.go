package evaluator

import (
	"fmt"

	"github.com/joshuahenriques/cixac-interpreter/ast"
	"github.com/joshuahenriques/cixac-interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
		fmt.Sprintf("")
	}

	return result
}
