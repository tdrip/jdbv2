package jdb

import (
	"encoding/json"

	uuid "github.com/google/uuid"
	i "github.com/tdrip/jdbv2/pkg/interfaces"
)

type TodoData struct {
	i.IKeyedItem `json:"-"`
	ID           string `json:"id"`
	Name         string `json:"name"`
}

func (sd TodoData) GetID() string {
	return sd.ID
}

func (sd TodoData) NewID() string {
	id := uuid.New()
	return id.String()
}

func encTodoData(items map[string]i.IKeyedItem) ([]byte, error) {
	converted := make(map[string]TodoData, 0)
	for k, v := range items {
		kp := v.(TodoData)
		converted[k] = kp
	}
	b, err := json.Marshal(converted)
	if err != nil {
		return nil, err
	}
	return b, err
}

func decTodoData(raw []byte) (map[string]i.IKeyedItem, error) {
	items := make(map[string]TodoData, 0)
	converted := make(map[string]i.IKeyedItem, 0)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		return nil, err
	}
	for k, v := range items {
		converted[k] = v
	}
	return converted, nil
}
