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
				return newError("key doesn't exist in HASH")
			}

			return val.Value
		},
	},
	"set": {
		Fn: func(args ...Object) Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `set` must be HASH, got %s", args[0].Type())
			}

			hashableKey, ok := args[1].(Hashable)
			if !ok {
				return newError("argument key to `set` must be Hashable")
			}

			hash := args[0].(*Hash)
			key := args[1]
			val := args[2]
			hash.Pairs[hashableKey.HashKey()] = HashPair{Key: key, Value: val}

			return EMPTY
		},
	},
	"delete": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `delete` must be HASH, got %s", args[0].Type())
			}

			key, ok := args[1].(Hashable)
			if !ok {
				return newError("argument key to `delete` must be Hashable")
			}

			hash := args[0].(*Hash)
			delete(hash.Pairs, key.HashKey())

			return EMPTY
		},
	},
	"values": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `values` must be HASH, got %s", args[0].Type())
			}

			hash := args[0].(*Hash)
			values := make([]Object, 0, len(hash.Pairs))

			for _, val := range hash.Pairs {
				values = append(values, val.Value)
			}

			return &Array{Elements: values}
		},
	},
	"keys": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `keys` must be HASH, got %s", args[0].Type())
			}

			hash := args[0].(*Hash)
			keys := make([]Object, 0, len(hash.Pairs))

			for _, val := range hash.Pairs {
				keys = append(keys, val.Key)
			}

			return &Array{Elements: keys}
		},
	},
	"clear": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `clear` must be HASH, got %s", args[0].Type())
			}

			hash := args[0].(*Hash)
			clear(hash.Pairs)

			return hash
		},
	},
	"contains": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != HASH_OBJ {
				return newError("argument to `contains` must be HASH, got %s", args[0].Type())
			}

			hash := args[0].(*Hash)
			key, ok := args[1].(Hashable)
			if !ok {
				return newError("argument to `contains` must the HASHABLE, got %s", args[1].Type())
			}

			_, ok = hash.Pairs[key.HashKey()]
			if !ok {
				return FALSE
			}

			return TRUE
		},
	},
}
