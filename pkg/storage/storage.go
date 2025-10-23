package storage

import (
	"errors"

	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

func ProcessStorage(storage chan i.IStorage, be i.IBackend, encdata i.EncodeKeyItems, decdata i.DecodeKeyItems) {
	for {
		s := <-storage

		content, err := be.Read()
		if err != nil {
			s.Exit() <- err
			continue
		}

		converted, err := decdata(content)
		if err != nil {
			s.Exit() <- err
			continue
		}

		modified, err := s.Run(converted)
		if err != nil {
			s.Exit() <- err
			continue
		}
		if !s.ReadOnly() {
			if modified == nil {
				s.Exit() <- errors.New("modified data was nil")
				continue
			}

			b, err := encdata(modified)
			if err != nil {
				s.Exit() <- err
				continue
			}

			err = be.Write(b)
			if err != nil {
				s.Exit() <- err
				continue
			}
		}
		s.Exit() <- err
	}
}
