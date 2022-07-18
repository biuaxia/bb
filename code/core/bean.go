package core

type Bean interface {
	Init()
	Cleanup()
	Bootstrap()
	PanicError(err error)
}
