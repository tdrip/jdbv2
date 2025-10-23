package jdb

import i "github.com/tdrip/jdbv2/pkg/interfaces"

// read all items from the database
type countStorageItems struct {
	total chan int
	exit  chan error
}

func newCountStorageItems() *countStorageItems {
	return &countStorageItems{
		total: make(chan int, 1),
		exit:  make(chan error, 1),
	}
}

func (csi countStorageItems) ReadOnly() bool {
	return true
}

func (csi countStorageItems) Exit() chan error {
	return csi.exit
}

func (csi countStorageItems) Run(items map[string]i.IKeyedItem) (map[string]i.IKeyedItem, error) {
	csi.total <- len(items)
	return items, nil
}
