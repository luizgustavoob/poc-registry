package ports

type Registry interface {
	Register() error
	Unregister()
}
