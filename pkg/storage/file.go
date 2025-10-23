package storage

import (
	"errors"
	"os"

	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

type FileStorgae struct {
	Path string
	Perm os.FileMode
}

func (fs FileStorgae) Read() ([]byte, error) {
	return os.ReadFile(fs.Path)
}

func (fs FileStorgae) Write(data []byte) error {
	return os.WriteFile(fs.Path, data, fs.Perm)
}

func (fs FileStorgae) Intiliase(encdata i.EncodeKeyItems) error {
	if len(fs.Path) == 0 {
		return errors.New("file path for json database missing")
	}

	if _, err := os.ReadFile(fs.Path); err != nil {
		empty := make(map[string]i.IKeyedItem, 0)
		b, err := encdata(empty)
		if err != nil {
			return err
		}
		if err = os.WriteFile(fs.Path, b, 0644); err != nil {
			return err
		}
	}

	return nil
}
