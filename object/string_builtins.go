package object

import (
	"strings"
	"unicode"
)

var StringBuiltins = map[string]Builtin{
	"lower": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != STRING_OBJ {
				return newError("argument to `lower` must be STRING, got %s", args[0].Type())
			}

			str := args[0].(*String)

			if len(str.Value) == 0 {
				return newError("string must have length greater than 0")
			}

			str.Value = strings.ToLower(str.Value)

			return str
		},
	},
	"upper": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != STRING_OBJ {
				return newError("argument to `upper` must be STRING, got %s", args[0].Type())
			}

			str := args[0].(*String)

			if len(str.Value) == 0 {
				return newError("string must have length greater than 0")
			}

			str.Value = strings.ToUpper(str.Value)

			return str
		},
	},
	"capitalize": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != STRING_OBJ {
				return newError("argument to `capitalize` must be STRING, got %s", args[0].Type())
			}

			str := args[0].(*String)

			if len(str.Value) == 0 {
				return newError("string must have length greater than 0")
			}

			r := []rune(str.Value)
			r[0] = unicode.ToUpper(r[0])
			str.Value = string(r)

			return str
		},
	},
	"split": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != STRING_OBJ {
				return newError("argument to `split` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != STRING_OBJ {
				return newError("argument to `split` mus tbe STRING, got %s, args[1].Type()")
			}

			str := args[0].(*String)

			if len(str.Value) == 0 {
				return newError("string must have length greater than 0")
			}

			delim := args[1].(*String)

			split := strings.Split(str.Value, delim.Value)
			splitArr := &Array{}

			for _, str := range split {
				splitArr.Elements = append(splitArr.Elements, &String{Value: str})
			}

			return splitArr
		},
	},
}
