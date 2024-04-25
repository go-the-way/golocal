package golocal

type Context[T any] interface {
	Set(data T)
	Get() (data T)
}
