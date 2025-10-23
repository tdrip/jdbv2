package interfaces

type EncodeKeyItems func(map[string]IKeyedItem) ([]byte, error)

type DecodeKeyItems func([]byte) (map[string]IKeyedItem, error)

type IBackend interface {
	Intiliase(encdata EncodeKeyItems) error
	Read() ([]byte, error)
	Write([]byte) error
}
