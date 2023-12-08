package in_memory

import "sync"

func NewMemory() *sync.Map {
	return &sync.Map{}
}
