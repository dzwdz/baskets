package device

type Device interface {
	// TODO Identify
	GetKey(service string) ([]byte, error)
}
