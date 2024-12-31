package evaluation

import (
	"github.com/singlaanish56/Interpreter-In-Go/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn : func(args ...object.Object) object.Object{
			if len(args)!=1{
				return newError("wrong number of args, expected=1, got=%d", len(args))
			}

			switch arg := args[0].(type){
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument for the len builtin not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn : func(args ...object.Object) object.Object{
			if len(args)!=1{
				return newError("wrong number of args, expected=1, got=%d", len(args))
			}

			switch arg := args[0].(type){
			case *object.Array:
				if len(arg.Elements) > 0{
					return arg.Elements[0]
				}
				return NULL
			default:
				return newError("argument for the first builtin not supported, got %s", args[0].Type())
			}
		},
	},
	"last": &object.Builtin{
		Fn : func(args ...object.Object) object.Object{
			if len(args)!=1{
				return newError("wrong number of args, expected=1, got=%d", len(args))
			}

			switch arg := args[0].(type){
			case *object.Array:
				if len(arg.Elements) > 0{
					return arg.Elements[len(arg.Elements) - 1]
				}
				return NULL
			default:
				return newError("argument for the last builtin not supported, got %s", args[0].Type())
			}
		},
	},
	"rest": &object.Builtin{
		Fn : func(args ...object.Object) object.Object{
			if len(args)!=1{
				return newError("wrong number of args, expected=1, got=%d", len(args))
			}

			switch arg := args[0].(type){
			case *object.Array:
				sz := len(arg.Elements)
				if sz > 0{
					newArr := make([]object.Object, sz-1)
					copy(newArr, arg.Elements[1:sz])
					return &object.Array{Elements: newArr}
				}
				return NULL
			default:
				return newError("argument for the rest builtin not supported, got %s", args[0].Type())
			}
		},
	},
	"push_back": &object.Builtin{
		Fn : func(args ...object.Object) object.Object{
			if len(args)!=2{
				return newError("wrong number of args, expected=2, got=%d", len(args))
			}

			switch arg := args[0].(type){
			case *object.Array:
				sz := len(arg.Elements)
				if sz > 0{
					newArr := make([]object.Object, sz+1)
					copy(newArr, arg.Elements)
					newArr[sz] =args[1]
					return &object.Array{Elements: newArr}
				}
				return NULL
			default:
				return newError("argument for the push_back builtin not supported, got %s", args[0].Type())
			}
		},
	},
}