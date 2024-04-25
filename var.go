package golocal

var (
	defContext = NewDefaultContext[string]()

	Set = defContext.Set
	Get = defContext.Get
)
