package jdb

import (
	"fmt"

	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

// delete one item from the database
type deleteStorageItem struct {
	id   string
	item chan i.IKeyedItem
	exit chan error
}

func newDeleteStorageItem(id string) *deleteStorageItem {
	return &deleteStorageItem{
		id:   id,
		item: make(chan i.IKeyedItem, 1),
		exit: make(chan error, 1),
	}
}

func (dsi deleteStorageItem) ReadOnly() bool {
	return false
}

func (dsi deleteStorageItem) Exit() chan error {
	return dsi.exit
}

func (dsi deleteStorageItem) Run(items map[string]i.IKeyedItem) (map[string]i.IKeyedItem, error) {
	_, exists := items[dsi.id]

	if !exists {
		err := fmt.Errorf("id %s does not exist", dsi.id)
		return nil, err
	}

	delete(items, dsi.id)
	return items, nil
}
