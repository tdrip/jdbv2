package jdb

import (
	i "github.com/tdrip/jdbv2/pkg/interfaces"
	s "github.com/tdrip/jdbv2/pkg/storage"
)

type Database struct {
	Storage chan i.IStorage
}

func BuildDatabase(be i.IBackend, encdata i.EncodeKeyItems, decdata i.DecodeKeyItems) (*Database, error) {

	err := be.Intiliase(encdata)
	if err != nil {
		return nil, err
	}

	// create storage channel to communicate over
	store := make(chan i.IStorage)

	// start watching storage channel for work
	go s.ProcessStorage(store, be, encdata, decdata)

	db := &Database{Storage: store}

	return db, nil
}

func (d *Database) SaveItem(todo i.IKeyedItem) (i.IKeyedItem, error) {
	job := newSaveStorageItem(todo)
	d.Storage <- job
	if err := <-job.Exit(); err != nil {
		return nil, err
	}
	return <-job.saved, nil
}

func (d *Database) CountItems() (int, error) {
	job := newCountStorageItems()
	d.Storage <- job

	if err := <-job.Exit(); err != nil {
		return -1, err
	}

	total := <-job.total
	return total, nil
}

func (d *Database) GetItems() ([]i.IKeyedItem, error) {
	arr := make([]i.IKeyedItem, 0)
	job := newReadStorageItems()
	d.Storage <- job

	if err := <-job.Exit(); err != nil {
		return arr, err
	}

	items := <-job.items
	for _, value := range items {
		arr = append(arr, value)
	}
	return arr, nil
}

func (d *Database) GetItem(id string) (i.IKeyedItem, error) {
	job := newReadStorageItem(id)
	d.Storage <- job
	if err := <-job.Exit(); err != nil {
		return nil, err
	}
	item := <-job.item
	return item, nil
}

func (d *Database) DeleteItem(id string) error {
	job := newDeleteStorageItem(id)
	d.Storage <- job

	if err := <-job.Exit(); err != nil {
		return err
	}
	return nil
}
