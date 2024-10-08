package evaluator

import (
	"sort"
	"testing"

	"github.com/beorn7/floats"
	"github.com/joshuahenriques/cixac/lexer"
	"github.com/joshuahenriques/cixac/object"
	"github.com/joshuahenriques/cixac/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"4 % 10", 4},
		{"-4 % 10", 6},
		{"10 % 4", 2},
		{"10 % -4", -2},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.5", 5.5},
		{"5.", 5.0},
		{".5", 0.5},
		{"10.0", 10.0},
		{"-5.0", -5.0},
		{"-10.0", -10.0},
		{"5.1 + 5.2 + 5.1 + 5.2 - 10", 10.6},
		{"2.2 * 2.2 * 2.2 * 2.2 * 2.2", 51.536320},
		{"-50.55 + 100.55 + -50.55", -0.55},
		{"5.5 * 2.2 + 10.10", 22.20},
		{"5.5 + 2.2 * 10.10", 27.72},
		{"20.2 + 2.2 * -10.1", -2.02},
		{"4.5 % 10.5", 4.50},
		{"-4.5 % 10.55", 6.05},
		{"10.5 % 4.4", 1.70},
		{"10.5 % -4.5", -3.00},
		{"55.55 / 2 * 2 + 10", 65.55},
		{"2 * (5 + 10.55)", 31.10},
		{"3 * 33.3 * 3 + 10", 309.70},
		{"3 * (3.33 * 3.3) + 10", 42.967},
		{"(5 + 10.95 * 2 + 15 / 3) * 2 + -10.55", 53.25},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, i, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 <= 2", true},
		{"1 >= 2", false},
		{"1.1 >= 1.1", true},
		{"1.1 >= 1.2", false},
		{"1.1 <= 1.1", true},
		{"1.1 <= 0.2", false},
		{"1.1 == 1.1", true},
		{"1.1 == 1.2", false},
		{"1.1 != 1.2", true},
		{"1.1 != 1.1", false},
		{"1.1 > 1.2", false},
		{"1.1 > 0.2", true},
		{"1.1 < 1.2", true},
		{"1.1 < 0.2", false},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 < 2) && true", true},
		{"(1 > 2) && true", false},
		{"(1 > 2) || false", false},
		{"(1 > 2) || true", true},
		{`"foobar" == "foobar"`, true},
		{`"foobar" == "foo"`, false},
		{`"foobar" != "foo"`, true},
		{`"foobar" != "foobar"`, false},
		{`null == null`, true},
		{`null != null`, false},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)
	}
}

func TestEvalNullExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`let a = null; a`, nil},
		{`if (5 == 5) { return null }`, nil},
		{`null`, nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNullObject(t, evaluated)
	}
}

func TestIncrDecrOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let i = 5; i++; i", 6},
		{"let j = 4; j--; j", 3},
		{"let k = 5; k++", 5},
		{"let l = 5; l--", 5},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else if (1 == 1) { 1 + 11 } else { 20 }", 12},
		{"if (null) { 10 } else { 20 }", 20},
		{"if (1 < 2) { null } else { 20 }", nil},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return 10;
				}
        // this is a comment
				return 1;
			}`, 10,
		},
		{
			`
      let f = fn(x) {
        /* multi
          multi-line comment
        */
        return x;
        x + 10;
      };
      f(10);`, 10,
		},
		{
			`
      let f = fn(x) {
        let result = x + 10;
        return result;
        return 10;
      };
      f(10);`,
			20,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}
				return 1;
			}`, "unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"Identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
		{
			`999[1]`,
			"index operator not supported: INTEGER",
		},
		{
			`null + null`,
			"unknown operator: NULL + NULL",
		},
		{
			`{"name": "Hello"}[fn(x) { x }];`,
			"unusable as hash key: FUNCTION",
		},
		{
			`let four = 4; let four = 5`,
			"Identifier four has already been declared",
		},
		{
			`fn five() { 5 } fn five() { 6 }`,
			`Function five has already been declared`,
		},
		{
			`fn decl() { 5 } let decl = 6`,
			`Identifier decl has already been declared`,
		},
		{
			`print = "foo"`,
			`Can't reassign print builtin function`,
		},
		{
			`object = "Person"`,
			`Identifier object doesn't exists`,
		},
		{
			`fn adder(x, y) { x + y } adder = 8`,
			`Identifier adder is const and can't be reassigned`,
		},
		{
			`const a = 7; a = 10`,
			`Identifier a is const and can't be reassigned`,
		},
		{
			`const i = 5; i++`,
			`Identifier i is const and can't be reassigned`,
		},
		{
			`5++`,
			`Invalid left-hand expression for postfix operation`,
		},
		{
			`break`,
			`break not in for statement`,
		},
		{
			`continue`,
			`continue not in for statement`,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("[test: %d] no error object returned. got=%T(%+v)", i, evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestForInLoopStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`let j = 0; for (i, ele in [1, 2, 3]) { j += i + ele }; j`, 9},
		{`let j = 0; for (key, val in {0: 1, 1: 2, 2: 3}) { j += key + val}; j`, 9},
		{`let j = 0; for (i, ch in "string") { j += i; }; j`, 15},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestForLoopStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`let j = 0; for (let i = 0; i < 5; i += 1) { j++ }; j`, 5},
		{`let j = 0; for (let i = 0; i < 5; i = i + 1) { j++ }; j`, 5},
		{`let j = 0; for (let i = 0; i < 5; i++) { j++ }; j`, 5},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestWhileStatement(t *testing.T) {
	input := `let i = 0; while (i < 5) { i++ }; i`
	testIntegerObject(t, 1, testEval(input), 5)
}

func TestBreakStatement(t *testing.T) {
	input := `let j = 0; for (let i = 0; i < 5; i++) { j++; break; }; j`
	testIntegerObject(t, 1, testEval(input), 1)
}

func TestContinueStatement(t *testing.T) {
	input := `let j = 0; for (let i = 0; i < 5; i++) { continue; j++; } j`
	testIntegerObject(t, 1, testEval(input), 0)
}

func TestReassignStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"let a = 5; a = 10; a;", 10},
		{"let a = 5 * 5; a = 5 * 11; a;", 55},
		{"let a = 5; let b = a; a = a * b; a;", 25},
		{"let a = 5; let b = a; let c = a + b + 5; c = c + 100; c;", 115},
		{"let a = 5; a += 2", 7},
		{"let a = 5; a -= 2", 3},
		{"let a = 5; a *= 2", 10},
		{"let a = 6; a /= 2", 3.0},
		{"let a = 5; a += 2.2", 7.2},
		{"let a = 5; a -= 2.2", 2.8},
		{"let a = 5; a *= 2.2", 11.0},
		{"let a = 11; a /= 2.2", 5.0},
	}

	for i, tt := range tests {
		switch v := tt.expected.(type) {
		case float64:
			testFloatObject(t, i, testEval(tt.input), v)
		case int:
			testIntegerObject(t, i, testEval(tt.input), int64(v))
		default:
			t.Errorf("Unexpected type for test case %d: %T", i, v)
		}
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
		{"fn adder(x, y) { return x + y }; adder(5, 5)", 10},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestEnclosingEnvironments(t *testing.T) {
	input := `
let x = 10;
let y = 10;
let z = 10;
// this is a comment

let ourFunction = fn(x) {
  // this is a comment
  let y = 20;

  x + y + z;
};

ourFunction(20) + x + y;`

	testIntegerObject(t, 0, testEval(input), 70)
}

func TestClosures(t *testing.T) {
	input := `
    let newAdder = fn(x) {
      fn(y) { x + y }
    }

    let addTwo = newAdder(2)
    addTwo(2)`

	testIntegerObject(t, 0, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", str, str)
	}

	if str.Value != "Hello World!" {
		t.Fatalf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"Hello" + " " + 5`, "Hello 5"},
		{`"Hello" + " " + 5.5`, "Hello 5.5"},
		{`"Hello" + " " + "World!"`, "Hello World!"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}

		if str.Value != tt.expected {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`len([1, 2, 3])`, 3},
		{`len({})`, 0},
		{`len([])`, 0},
		{`len({"key1": 5, "key2": 10})`, 2},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, i, evaluated, int64(expected))
		case nil:
			testNullObject(t, evaluated)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		case []int:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			for i, expectedElem := range expected {
				testIntegerObject(t, i, array.Elements[i], int64(expectedElem))
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, 0, result.Elements[0], 1)
	testIntegerObject(t, 1, result.Elements[1], 4)
	testIntegerObject(t, 2, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]", 1,
		},
		{
			"[1, 2, 3][1]", 2,
		},
		{
			"[1, 2, 3][2]", 3,
		},
		{
			"let i = 0; [1][i];", 1,
		},
		{
			"[1, 2, 3][1 + 1];", 3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];", 3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];", 6,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]", 2,
		},
		{
			"[1, 2, 3][3]", nil,
		},
		{
			"[1, 2, 3][-1]", nil,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
    {
      "one": 10 - 9,
      two: 1 + 1,
      "thr" + "ee": 6 / 2,
      4: 4,
      true: 5,
      false: 6
    }`
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}
	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}
	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}
	testNum := 0
	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}
		testIntegerObject(t, testNum, pair.Value, expectedValue)
		testNum++
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`, 5,
		},
		{
			`{"foo": 5}["bar"]`, nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`, 5,
		},
		{
			`{}["foo"]`, nil,
		},
		{
			`{5: 5}[5]`, 5,
		},
		{
			`{true: 5}[true]`, 5,
		},
		{
			`{false: 5}[false]`, 5,
		},
	}
	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashBuiltinExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`{"key": 5}.get("key")`, 5},
		{`let map = {"key": 5}; map.get("key")`, 5},
		{`{}.get("key")`, "key doesn't exist in HASH"},
		{`1.get("key")`, "Identifier not found: 1.get"},
		{`let map = {"key": 5}; map.set("new", 8); map.get("new")`, 8},
		{`1.set("key")`, "Identifier not found: 1.set"},
		{`let map = {"key": 5}; map.delete("key"); map.get("key")`, "key doesn't exist in HASH"},
		{`let map = {"key1": 5, "key2": 10, "key3": 15}; map.values()`, []int{5, 10, 15}},
		{`{"key1": 5, "key2": 10, "key3": 15}.keys()`, []string{"key1", "key2", "key3"}},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, i, evaluated, int64(expected))
		case nil:
			testNullObject(t, evaluated)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		case []int:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			sort.Slice(array.Elements, func(i, j int) bool {
				vali := array.Elements[i].(*object.Integer).Value
				valj := array.Elements[j].(*object.Integer).Value
				return vali <= valj
			})

			for i, expectedElem := range expected {
				testIntegerObject(t, i, array.Elements[i], int64(expectedElem))
			}
		case []string:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			sort.Slice(array.Elements, func(i, j int) bool {
				vali := array.Elements[i].(*object.String).Value
				valj := array.Elements[j].(*object.String).Value
				return vali <= valj
			})

			for i, expectedElem := range expected {
				arrEle, ok := array.Elements[i].(*object.String)
				if !ok {
					t.Errorf("array element not string. got=%T (%+v)", array.Elements[i], array.Elements[i])
				}

				if arrEle.Value != expectedElem {
					t.Errorf("string has wrong value. got=%s, want=%s", arrEle, expectedElem)
				}
			}
		default:
			t.Errorf("no test for given expected type")
		}
	}
}

func TestArrayBuiltinExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`[1, 2, 3].first()`, 1},
		{`[].first()`, nil},
		{`1.first()`, "Identifier not found: 1.first"},
		{`[1, 2, 3].last()`, 3},
		{`[].last()`, nil},
		{`1.last()`, "Identifier not found: 1.last"},
		{`[1, 2, 3].rest()`, []int{2, 3}},
		{`[].rest()`, nil},
		{"let myArray = [1, 2, 3]; myArray.push(1 + 1); len(myArray)", 4},
		{`[].push(1)`, []int{1}},
		{`1.push(1)`, "Identifier not found: 1.push"},
		{`[].pushleft(1)`, []int{1}},
		{`[1, 2, 3].pushleft(4)`, []int{4, 1, 2, 3}},
		{`1.pushleft(1)`, "Identifier not found: 1.pushleft"},
		{`[1, 2, 3].pop()`, 3},
		{`let arr = [1, 2, 3]; arr.pop(); arr`, []int{1, 2}},
		{`[].pop()`, "ARRAY must have elements for `pop`"},
		{`1.pop()`, "Identifier not found: 1.pop"},
		{`[1, 2, 3].popleft()`, 1},
		{`let arr = [1, 2, 3]; arr.popleft(); arr`, []int{2, 3}},
		{`[].popleft()`, "ARRAY must have elements for `popleft`"},
		{`1.popleft()`, "Identifier not found: 1.popleft"},
		{`let arr = [1, 2, 3, 4, 5]; arr.slice(1, 3)`, []int{2, 3}},
		{`[1, 2, 3, 4, 5].slice(1, 3)`, []int{2, 3}},
		{`[1, 2, 3, 4, 5].slice(1, 8)`, "slice bounds out of range, [1:8] with array len of 5"},
		{`[1, 2, 3, 4, 5].slice(3, 1)`, "slice bounds out of range, [3:1] with array len of 5"},
		{`[1, 2, 3, 4, 5].slice(1, 3, 5)`, "wrong number of arguments. got=3, want=2 or 3"},
		{`[].slice(1, 3)`, "array must have elements"},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, i, evaluated, int64(expected))
		case nil:
			testNullObject(t, evaluated)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		case []int:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			for i, expectedElem := range expected {
				testIntegerObject(t, i, array.Elements[i], int64(expectedElem))
			}
		default:
			t.Errorf("no test for given expected type")
		}
	}
}

func TestStringIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"\"jump\"[0]", "j",
		},
		{
			"\"blender\"[1]", "l",
		},
		{
			"let str = \"string\"; let length = len(str); str[length-1];", "g",
		},
		{
			"\"elden\"[1 + 1];", "d",
		},
		{
			"let str = \"finished\"; str[0] + str[1] + str[2];", "fin",
		},
		{
			"\"string\"[9]", nil,
		},
		{
			"\"string\"[-1]", nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		char, ok := tt.expected.(string)
		if ok {
			testStringObject(t, evaluated, string(char))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, i int, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("[test: %d] object is not Integer. got=%T (%+v)", i, obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("[test: %d] object has wrong value. got=%d, want=%d", i, result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, i int, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("[test: %d] object is not Float. got=%T (%+v)", i, obj, obj)
		return false
	}

	if !floats.AlmostEqual(result.Value, expected, 0.00001) {
		t.Errorf("[test: %d] object has wrong value. got=%f, want=%f", i, result.Value, expected)
		return false
	}

	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, i int, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("[test id: %d] object is not Boolean. go=%T (%+v)", i, obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("[test id: %d] object has wrong value. got=%t, want=%t",
			i, result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got =%T (%+v)", obj, obj)
		return false
	}
	return true
}
