package storage

type Storage interface {
	Write(i interface{}) (err error)
	Close()
}
