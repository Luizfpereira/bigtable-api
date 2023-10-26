package entity

import "time"

type BigtableOutput struct {
	Key     string    `json:"key"`
	Created time.Time `json:"created"`
	Value   string    `json:"value"`
}
