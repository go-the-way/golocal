package golocal

var (
	defContext = newDefaultContext[string]()

	Set = defContext.Set
	Get = defContext.Get
)
