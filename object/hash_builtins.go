package object

var HashBuiltins = map[string]Builtin{
	"get": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `get` must be HASH, got %s", args[0].Type())
			}

			key, ok := args[1].(Hashable)
			if !ok {
				return newError("argument key to `get` must be Hashable")
			}

			hash := args[0].(*Hash)

			val, ok := hash.Pairs[key.HashKey()]
			if !ok {
				return newError("key doesn't exists in HASH")
			}

			return val.Value
		},
	},
}
