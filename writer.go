package oututils

type Writer interface {
	Write(interface{}) error
}
