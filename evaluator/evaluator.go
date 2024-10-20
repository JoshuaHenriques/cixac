package evaluator

import (
	"fmt"
	"math"
	"strconv"

	"github.com/beorn7/floats"
	"github.com/joshuahenriques/cixac/ast"
	"github.com/joshuahenriques/cixac/object"
)

// Avoid creating object.Boolean & object.Null every time
var (
	NULL                    = &object.Null{}
	EMPTY                   = &object.Empty{}
	TRUE                    = &object.Boolean{Value: true}
	FALSE                   = &object.Boolean{Value: false}
	BREAK                   = &object.Break{}
	CONTINUE                = &object.Continue{}
	ENV_FOR_FLAG            = "ENV_FOR_FLAG"
	ENV_WHILE_FLAG          = "ENV_WHILE_FLAG"
	ENV_OBJECT_BUILTIN_FLAG = "ENV_OBJECT_BUILTIN_FLAG"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}

		if val != BREAK || val != CONTINUE {
			return &object.ReturnValue{Value: val}
		}

	case *ast.LetStatement:
		if ExistsInBuiltins(node.Name.Value) {
			return newError("Identifier %s has same name as builtin", node.Name.Value)
		}

		if env.ExistsInScope(node.Name.Value) && !env.ExistsInScope(ENV_FOR_FLAG) {
			return newError("Identifier %s has already been declared", node.Name.Value)
		}

		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		obj := object.ObjectMeta{Object: val, Const: node.Name.Const}
		env.Set(node.Name.Value, obj)

	case *ast.FunctionDeclaration:
		if env.ExistsInScope(node.Name.Value) {
			return newError("Function %s has already been declared", node.Name.Value)
		}

		if ExistsInBuiltins(node.Name.Value) {
			return newError("Identifier %s has same name as builtin", node.Name.Value)
		}

		val := Eval(node.Function, env)
		if isError(val) {
			return val
		}

		env.Set(node.Name.Value, object.ObjectMeta{Object: val, Const: node.Name.Const})

	case *ast.ReassignStatement:
		if ExistsInBuiltins(node.Name.Value) {
			return newError("Can't reassign %s builtin function", node.Name.Value)
		}

		obj, ok := env.Get(node.Name.Value)
		if !ok {
			return newError("Identifier %s doesn't exists", node.Name.Value)
		}

		if obj.Const {
			return newError("Identifier %s is const and can't be reassigned", node.Name.Value)
		}

		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		switch node.TokenLiteral() {
		case "+=":
			if obj.Object.Type() == object.INTEGER_OBJ {
				if val.Type() == object.FLOAT_OBJ {
					val.(*object.Float).Value = float64(obj.Object.(*object.Integer).Value) + val.(*object.Float).Value
				} else {
					val.(*object.Integer).Value = obj.Object.(*object.Integer).Value + val.(*object.Integer).Value
				}
			} else if obj.Object.Type() == object.FLOAT_OBJ {
				if val.Type() == object.INTEGER_OBJ {
					val = &object.Float{Value: obj.Object.(*object.Float).Value + float64(val.(*object.Integer).Value)}
				} else {
					val.(*object.Float).Value = obj.Object.(*object.Float).Value + val.(*object.Float).Value
				}
			}
		case "-=":
			if obj.Object.Type() == object.INTEGER_OBJ {
				if val.Type() == object.FLOAT_OBJ {
					val.(*object.Float).Value = float64(obj.Object.(*object.Integer).Value) - val.(*object.Float).Value
				} else {
					val.(*object.Integer).Value = obj.Object.(*object.Integer).Value - val.(*object.Integer).Value
				}
			} else if obj.Object.Type() == object.FLOAT_OBJ {
				if val.Type() == object.INTEGER_OBJ {
					val = &object.Float{Value: obj.Object.(*object.Float).Value - float64(val.(*object.Integer).Value)}
				} else {
					val.(*object.Float).Value = obj.Object.(*object.Float).Value - val.(*object.Float).Value
				}
			}
		case "*=":
			if obj.Object.Type() == object.INTEGER_OBJ {
				if val.Type() == object.FLOAT_OBJ {
					val.(*object.Float).Value = float64(obj.Object.(*object.Integer).Value) * val.(*object.Float).Value
				} else {
					val.(*object.Integer).Value = obj.Object.(*object.Integer).Value * val.(*object.Integer).Value
				}
			} else if obj.Object.Type() == object.FLOAT_OBJ {
				if val.Type() == object.INTEGER_OBJ {
					val = &object.Float{Value: obj.Object.(*object.Float).Value * float64(val.(*object.Integer).Value)}
				} else {
					val.(*object.Float).Value = obj.Object.(*object.Float).Value * val.(*object.Float).Value
				}
			}
		case "/=":
			if obj.Object.Type() == object.INTEGER_OBJ {
				if val.Type() == object.FLOAT_OBJ {
					val.(*object.Float).Value = float64(obj.Object.(*object.Integer).Value) / val.(*object.Float).Value
				} else {
					val = &object.Float{Value: float64(obj.Object.(*object.Integer).Value / val.(*object.Integer).Value)}
				}
			} else if obj.Object.Type() == object.FLOAT_OBJ {
				if val.Type() == object.INTEGER_OBJ {
					val = &object.Float{Value: obj.Object.(*object.Float).Value / float64(val.(*object.Integer).Value)}
				} else {
					val.(*object.Float).Value = obj.Object.(*object.Float).Value / val.(*object.Float).Value
				}
			}
		}

		if env.ExistsInScope(ENV_FOR_FLAG) && !env.ExistsInScope(node.Name.Value) && env.ExistsOutsideScope(node.Name.Value) {
			env.SetOutsideScope(node.Name.Value, object.ObjectMeta{Object: val})
		} else {
			env.Set(node.Name.Value, object.ObjectMeta{Object: val})
		}

		return val

	case *ast.ForLoopStatement:
		return evalForLoopStatement(node, env)

	case *ast.ForInLoopStatement:
		return evalForInLoopStatement(node, env)

	case *ast.WhileStatement:
		return evalWhileStatement(node, env)

	case *ast.BreakStatement:
		if !env.ExistsInScope(ENV_FOR_FLAG) && !env.ExistsInScope(ENV_WHILE_FLAG) {
			return newError("break not in for statement")
		}
		return BREAK

	case *ast.ContinueStatement:
		if !env.ExistsInScope(ENV_FOR_FLAG) && !env.ExistsInScope(ENV_WHILE_FLAG) {
			return newError("continue not in for statement")
		}
		return CONTINUE

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Env: env, Body: body}

	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.HashLiteral:
		return evalHashLiteral(node, env)

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.Null:
		return nativeNulltoNullObject()

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.PostfixExpression:
		ident, ok := node.Left.(*ast.Identifier)
		if !ok {
			return newError("Invalid left-hand expression for postfix operation")
		}

		obj, ok := env.Get(ident.Value)
		if !ok {
			return newError("Identifier %s doesn't exists", ident.Value)
		}

		if obj.Const {
			return newError("Identifier %s is const and can't be reassigned", ident.Value)
		}

		val, retVal := evalPostfixExpression(node.Operator, obj.Object)

		if env.ExistsInScope(ENV_FOR_FLAG) && !env.ExistsInScope(ident.Value) && env.ExistsOutsideScope(ident.Value) {
			env.SetOutsideScope(ident.Value, object.ObjectMeta{Object: val})
		} else {
			env.Set(ident.Value, object.ObjectMeta{Object: val})
		}

		return retVal

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	case *ast.BuiltinExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		methEnv := object.NewEnclosedEnvironment(env)
		methEnv.Set(ENV_OBJECT_BUILTIN_FLAG, object.ObjectMeta{Object: left})

		method := Eval(node.Builtin.Function, methEnv)
		if isError(method) {
			return method
		}

		args := evalExpressions(node.Builtin.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		args = append([]object.Object{left}, args...)
		return applyFunction(method, args)

	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
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

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	forLoopFlag := env.ExistsInScope(ENV_FOR_FLAG)
	whileLoopFlag := env.ExistsInScope(ENV_WHILE_FLAG)
	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}

			if (forLoopFlag || whileLoopFlag) && (rt == object.BREAK_OBJ || rt == object.CONTINUE_OBJ) {
				return result
			}

			if (!forLoopFlag && !whileLoopFlag) && (rt == object.BREAK_OBJ || rt == object.CONTINUE_OBJ) {
				return newError("%s not in loop", rt)
			}
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func nativeNulltoNullObject() *object.Null {
	return NULL
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalPostfixExpression(operator string, left object.Object) (object.Object, object.Object) {
	switch left.Type() {
	case object.INTEGER_OBJ:
		objVal := left.(*object.Integer)
		switch operator {
		case "++":
			return &object.Integer{Value: objVal.Value + 1}, objVal
		case "--":
			return &object.Integer{Value: objVal.Value - 1}, objVal
		}

	case object.FLOAT_OBJ:
		objVal := left.(*object.Float)
		switch operator {
		case "++":
			return &object.Float{Value: objVal.Value + 1}, objVal
		case "--":
			return &object.Float{Value: objVal.Value - 1}, objVal
		}

	default:
		return newError("wrong type for postfix operator"), nil
	}

	return nil, nil
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	var obj object.Object

	switch right.Type() {
	case object.INTEGER_OBJ:
		obj = &object.Integer{Value: -right.(*object.Integer).Value}
	case object.FLOAT_OBJ:
		obj = &object.Float{Value: -right.(*object.Float).Value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}

	return obj
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case (left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ) ||
		(left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ) ||
		(left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ):
		return evalFloatInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.FLOAT_OBJ) ||
		(left.Type() == object.FLOAT_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ) ||
		(left.Type() == object.INTEGER_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.BOOLEAN_OBJ) ||
		(left.Type() == object.BOOLEAN_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.ARRAY_OBJ) ||
		(left.Type() == object.ARRAY_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ && right.Type() == object.HASH_OBJ) ||
		(left.Type() == object.HASH_OBJ && right.Type() == object.STRING_OBJ):
		return evalStringInfixExpression(operator, left, right)

	// Using pointer comparison on object.Object becuase we're using
	// the same TRUE and FALSE pointers that we created above
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)

	case operator == "&&":
		return evalBooleanInfixExpression(operator, left, right)
	case operator == "||":
		return evalBooleanInfixExpression(operator, left, right)

	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: modLikePython(leftVal, rightVal)}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	var leftVal, rightVal float64

	if left.Type() == object.INTEGER_OBJ {
		leftVal = float64(left.(*object.Integer).Value)
	} else {
		leftVal = left.(*object.Float).Value
	}

	if right.Type() == object.INTEGER_OBJ {
		rightVal = float64(right.(*object.Integer).Value)
	} else {
		rightVal = right.(*object.Float).Value
	}

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "%":
		return &object.Float{Value: modLikePython(leftVal, rightVal)}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject((leftVal < rightVal) || floats.AlmostEqual(leftVal, rightVal, 0.00001))
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject((leftVal > rightVal) || floats.AlmostEqual(leftVal, rightVal, 0.00001))
	case "==":
		return nativeBoolToBooleanObject(floats.AlmostEqual(leftVal, rightVal, 0.00001))
	case "!=":
		return nativeBoolToBooleanObject(!floats.AlmostEqual(leftVal, rightVal, 0.00001))
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "&&":
		return nativeBoolToBooleanObject(leftVal && rightVal)
	case "||":
		return nativeBoolToBooleanObject(leftVal || rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalForInLoopStatement(fl *ast.ForInLoopStatement, env *object.Environment) object.Object {
	var result object.Object

	forEnv := object.NewEnclosedEnvironment(env)
	forEnv.Set(ENV_FOR_FLAG, object.ObjectMeta{Object: TRUE})

	iterable := Eval(fl.Iterable, env)

	switch iterable := iterable.(type) {
	case *object.Array:
		for i, ele := range iterable.Elements {
			forEnv.Set(fl.KeyIndex.Value, object.ObjectMeta{Object: &object.Integer{Value: int64(i)}})
			forEnv.Set(fl.ValueElement.Value, object.ObjectMeta{Object: ele})

			result = Eval(fl.Body, forEnv)

			if result.Type() == object.BREAK_OBJ {
				break
			}
		}
	case *object.Hash:
		for _, hashPair := range iterable.Pairs {
			forEnv.Set(fl.KeyIndex.Value, object.ObjectMeta{Object: hashPair.Key})
			forEnv.Set(fl.ValueElement.Value, object.ObjectMeta{Object: hashPair.Value})

			result = Eval(fl.Body, forEnv)

			if result.Type() == object.BREAK_OBJ {
				break
			}
		}
	case *object.String:
		for i, ch := range iterable.Value {
			forEnv.Set(fl.KeyIndex.Value, object.ObjectMeta{Object: &object.Integer{Value: int64(i)}})
			forEnv.Set(fl.ValueElement.Value, object.ObjectMeta{Object: &object.String{Value: string(ch)}})

			result = Eval(fl.Body, forEnv)

			if result.Type() == object.BREAK_OBJ {
				break
			}
		}
	}

	forEnv.Delete(ENV_FOR_FLAG)
	return result
}

func evalForLoopStatement(fl *ast.ForLoopStatement, env *object.Environment) object.Object {
	var result object.Object

	forEnv := object.NewEnclosedEnvironment(env)
	Eval(fl.Initialization, forEnv)

	forEnv.Set(ENV_FOR_FLAG, object.ObjectMeta{Object: TRUE})
	for isTruthy(Eval(fl.Condition, forEnv)) {
		result = Eval(fl.Body, forEnv)

		if result.Type() == object.BREAK_OBJ {
			break
		}

		Eval(fl.Update, forEnv)
	}

	forEnv.Delete(ENV_FOR_FLAG)
	return result
}

func evalWhileStatement(w *ast.WhileStatement, env *object.Environment) object.Object {
	var result object.Object

	env.Set(ENV_WHILE_FLAG, object.ObjectMeta{Object: TRUE})
	for isTruthy(Eval(w.Condition, env)) {
		result = Eval(w.Body, env)

		if result.Type() == object.BREAK_OBJ {
			break
		}
	}

	env.Delete(ENV_WHILE_FLAG)
	return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	for _, con := range ie.Conditions {
		condition := Eval(con.Condition, env)
		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			return Eval(con.Consequence, env)
		}
	}

	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val.Object
	}

	if env.ExistsInScope(ENV_OBJECT_BUILTIN_FLAG) {
		envObj, _ := env.Get(ENV_OBJECT_BUILTIN_FLAG)
		obj, ok := envObj.Object.(object.Methodable)
		if !ok {
			return newError("Object does not implement Methodable")
		}

		if objBuiltin, ok := obj.Methods(node.Value); ok {
			return objBuiltin
		}

		return newError("Method not found in object methods")
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("Identifier not found: " + node.Value)
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
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

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		obj := object.ObjectMeta{Object: args[paramIdx]}
		env.Set(param.Value, obj)
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func convertToString(obj object.Object) object.Object {
	switch obj.Type() {
	case object.INTEGER_OBJ:
		intNum := obj.(*object.Integer)
		string := &object.String{Value: intNum.Inspect()}
		return string
	case object.FLOAT_OBJ:
		floatNum := obj.(*object.Float)
		string := &object.String{Value: strconv.FormatFloat(floatNum.Value, 'f', -1, 64)}
		return string
	case object.BOOLEAN_OBJ:
		boolean := obj.(*object.Boolean)
		string := &object.String{Value: boolean.Inspect()}
		return string
	case object.ARRAY_OBJ:
		array := obj.(*object.Array)
		string := &object.String{Value: array.Inspect()}
		return string
	case object.HASH_OBJ:
		hash := obj.(*object.Hash)
		string := &object.String{Value: hash.Inspect()}
		return string
	case object.STRING_OBJ:
		return obj
	default:
		return nil
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	left = convertToString(left)
	right = convertToString(right)

	switch operator {
	case "+":
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.String).Value
		return &object.String{Value: leftVal + rightVal}
	case "!=":
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.String).Value
		return &object.Boolean{Value: leftVal != rightVal}
	case "==":
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.String).Value
		return &object.Boolean{Value: leftVal == rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.STRING_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalStringIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalStringIndexExpression(str, index object.Object) object.Object {
	stringObject := str.(*object.String)
	idx := index.(*object.Integer).Value
	max := int64(len(stringObject.Value) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	char := string([]rune(stringObject.Value)[idx])

	return &object.String{Value: char}
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

type Number interface {
	~int64 | ~float64
}

func modLikePython[T Number](x, y T) T {
	var res T

	switch any(x).(type) {
	case int64:
		res = T(any(x).(int64) % any(y).(int64))
	case float64:
		res = T(math.Mod(any(x).(float64), any(y).(float64)))
	}

	if res < 0 && y > 0 || (res > 0 && y < 0) {
		return res + y
	}

	return res
}
