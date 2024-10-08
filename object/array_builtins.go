package object

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
				return newError("wrong number of parameters. got=%d, want=0", len(args))
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
				return newError("wrong number of parameters. got=%d, want=0", len(args))
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
				return newError("wrong number of arguments. got=%d, want=0", len(args))
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
				return newError("wrong number of arguments. got=%d, want=0", len(args))
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
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)

			if len(arr.Elements) == 0 {
				return NULL
			}

			newElements := arr.Elements[1:]

			return &Array{Elements: newElements}
		},
	},
	"slice": {
		Fn: func(args ...Object) Object {
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `slice` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)

			if len(arr.Elements) == 0 {
				return newError("array must have elements")
			}

			lenArgs := len(args)
			lenArr := len(arr.Elements)
			switch lenArgs {
			case 2:
				if args[1].Type() != INTEGER_OBJ {
					return newError("arguments to slice must be INTEGER, got %s", args[1].Type())
				}

				idx := args[1].(*Integer)

				if idx.Value >= int64(lenArr-1) || idx.Value < 0 {
					return newError("slice bounds out of range, [:%d] with array len of %d", idx.Value, lenArr)
				}

				return &Array{Elements: arr.Elements[:idx.Value]}
			case 3:
				if args[1].Type() != INTEGER_OBJ && args[2].Type() != INTEGER_OBJ {
					return newError("arguments to slice must be INTEGER, got %s", args[1].Type())
				}

				idx1 := args[1].(*Integer)
				idx2 := args[2].(*Integer)

				if idx1.Value > idx2.Value || idx1.Value < 0 || idx2.Value >= int64(lenArr-1) {
					return newError("slice bounds out of range, [%d:%d] with array len of %d", idx1.Value, idx2.Value, lenArr)
				}

				return &Array{Elements: arr.Elements[idx1.Value:idx2.Value]}
			default:
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args)-1)
			}
		},
	},
	"clear": {
		Fn: func(args ...Object) Object {
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `slice` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			arr.Elements = []Object{}

			return arr
		},
	},
}
