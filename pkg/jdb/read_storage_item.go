package jdb

import (
	"fmt"

	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

// read one item from the database
type readStorageItem struct {
	id   string
	item chan i.IKeyedItem
	exit chan error
}

func newReadStorageItem(id string) *readStorageItem {
	return &readStorageItem{
		id:   id,
		item: make(chan i.IKeyedItem, 1),
		exit: make(chan error, 1),
	}
}

func (rsi readStorageItem) ReadOnly() bool {
	return true
}

func (rsi readStorageItem) Exit() chan error {
	return rsi.exit
}

func (rsi readStorageItem) Run(items map[string]i.IKeyedItem) (map[string]i.IKeyedItem, error) {
	_, exists := items[rsi.id]

	if !exists {
		err := fmt.Errorf("id %s does not exist", rsi.id)
		return nil, err
	}

	rsi.item <- items[rsi.id]
	return items, nil
}
