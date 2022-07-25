package evaluator

import (
	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/object"
)

var (
	NULL      = &object.Null{}
	UNDEFINED = &object.Undefined{}
	TRUE      = &object.Boolean{Value: true}
	FALSE     = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.ClosureStatement:
		return evalClosureStatement(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.OperationExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalOperationExpression(node.Operator, left, right)
	case *ast.DotExpression:
		return evalDotExpression(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.WhileExpression:
		return evalWhileExpression(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.AssignExpression:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.FunctionExpression:
		params := node.Parameters
		body := node.Expression
		name := node.Name
		fn := &object.Function{
			Parameters: params,
			Env:        env,
			Body:       body,
			Name:       name,
		}
		if name != nil {
			env.Set(node.Name.Value, fn)
		}
		return fn
	case *ast.FunctionApplication:
		return evalFunction(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.Null:
		return NULL
	case *ast.Undefined:
		return UNDEFINED
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalDotExpression(node *ast.DotExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env).(*object.BuiltinObj)
	if isError(left) {
		return left
	}
	clEnv := object.NewEnclosedEnvironment(env)
	for k, v := range left.Funcs {
		clEnv.Set(k, v)
	}
	right := Eval(node.Right, clEnv)
	if isError(right) {
		return right
	}
	return right
}
