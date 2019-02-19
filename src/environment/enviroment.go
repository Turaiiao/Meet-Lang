package environment

const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	LIST_OBJ    = "LIST"
	BOOL_OBJ    = "BOOL"
	FUN_OBJ     = "FUN"
)

type ObjectType string

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object), outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}

func (e Environment) All() map[string]Object {
	return e.store
}

// --------- Object Interface ---------

type Object interface {
	Type() ObjectType
	Inspect() string
}