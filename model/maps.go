package model

import "sync"

// Set ...
type Set struct {
	sync.RWMutex
	m map[string]*Sample
}

// Add add
func (s *Set) Add(key string, sample *Sample, cKey uint) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = sample
	s.m[key].Classes[cKey]++
}

// Has ...
func (s *Set) Has(key string) (ok bool) {
	s.RLock()
	defer s.RUnlock()
	_, ok = s.m[key]
	return
}

// Update ...
func (s *Set) Update(key string, cKey uint) {
	s.Lock()
	defer s.Unlock()
	s.m[key].Classes[cKey]++
	s.m[key].add()
}

// NewSet ...
func NewSet() *Set {
	return &Set{
		m: make(map[string]*Sample),
	}
}

// Cats ...
type Cats struct {
	sync.RWMutex
	cats map[uint]float64
}

// Add add
func (c *Cats) Add(key uint) {
	c.Lock()
	defer c.Unlock()
	c.cats[key]++
}

// NewCats ...
func NewCats() *Cats {
	return &Cats{
		cats: make(map[uint]float64),
	}
}
