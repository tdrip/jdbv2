package jdb

import (
	"errors"

	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

// save one item to the database
type saveStorageItem struct {
	item  i.IKeyedItem
	saved chan i.IKeyedItem
	exit  chan error
}

func newSaveStorageItem(item i.IKeyedItem) *saveStorageItem {
	return &saveStorageItem{
		item:  item,
		saved: make(chan i.IKeyedItem, 1),
		exit:  make(chan error, 1),
	}
}

func (ssi saveStorageItem) ReadOnly() bool {
	return false
}

func (ssi saveStorageItem) Exit() chan error {
	return ssi.exit
}

func (ssi saveStorageItem) Run(items map[string]i.IKeyedItem) (map[string]i.IKeyedItem, error) {
	item := ssi.item
	if item == nil {
		err := errors.New("item to be saved was nil")
		return nil, err
	}

	id := item.GetID()
	if len(id) == 0 {
		err := errors.New("item id was empty")
		return nil, err
	}

	items[id] = item
	ssi.saved <- item
	return items, nil
}
