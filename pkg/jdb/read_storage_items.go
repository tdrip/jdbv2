package jdb

import i "github.com/tdrip/jdbv2/pkg/interfaces"

// read all items from the database
type readStorageItems struct {
	items chan map[string]i.IKeyedItem
	exit  chan error
}

func newReadStorageItems() *readStorageItems {
	return &readStorageItems{
		items: make(chan map[string]i.IKeyedItem, 1),
		exit:  make(chan error, 1),
	}
}

func (rsi readStorageItems) ReadOnly() bool {
	return true
}

func (rsi readStorageItems) Exit() chan error {
	return rsi.exit
}

func (rsi readStorageItems) Run(items map[string]i.IKeyedItem) (map[string]i.IKeyedItem, error) {
	rsi.items <- items
	return nil, nil
}
