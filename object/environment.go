package object

type Environment struct {
	store map[string]ObjectMeta
	outer *Environment
}

type ObjectMeta struct {
	Object Object
	Const  bool
}

func NewEnvironment() *Environment {
	s := make(map[string]ObjectMeta)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (ObjectMeta, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, obj ObjectMeta) {
	e.store[name] = obj
}

func (e *Environment) ExistsInScope(name string) bool {
	_, ok := e.store[name]
	return ok
}
