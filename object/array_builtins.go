package object

var (
	NULL  = &Null{}
	EMPTY = &Empty{}
)

var ArrayBuiltins = map[string]Builtin{
	"push": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			arr.Elements = append(arr.Elements, []Object{args[1]}...)

			return arr
		},
	},
	"pop": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of parameters. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `pop` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			if length == 0 {
				return newError("ARRAY must have elements for `pop`")
			}

			popped := arr.Elements[length-1]
			arr.Elements = arr.Elements[:length-1]

			return popped
		},
	},
	"pushleft": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `pushleft` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			arr.Elements = append([]Object{args[1]}, arr.Elements...)

			return arr
		},
	},
	"popleft": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of parameters. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `popleft` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			if length == 0 {
				return newError("ARRAY must have elements for `popleft`")
			}

			popped := arr.Elements[0]
			arr.Elements = arr.Elements[1:length]

			return popped
		},
	},
	"first": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)

			if len(arr.Elements) == 0 {
				return NULL
			}

			return arr.Elements[0]
		},
	},
	"last": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			if length == 0 {
				return NULL
			}

			return arr.Elements[length-1]
		},
	},
	"rest": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)

			if len(arr.Elements) == 0 {
				return NULL
			}

			newElements := arr.Elements[1:]

			return &Array{Elements: newElements}
		},
	},
}
