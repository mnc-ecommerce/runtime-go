package contract

type Provide func(container any) error

type Application interface {
	Boot(providers ...Provide) error
	Start() error
	Shutdown()
}
