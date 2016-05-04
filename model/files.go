package model

import (
	"encoding/gob"
	"os"
)

// SaveModel ...
func SaveModel(file string, m map[string]*Samples) {
	f, err := os.Create(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(m); err != nil {
		panic("cant encode")
	}
}

// SaveCats ...
func SaveCats(file string, cats map[uint]float64) {
	f, err := os.Create(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(cats); err != nil {
		panic("cant encode")
	}
}

// Load ...
func Load(modelPath string, catsPath string) (m map[string]*Samples, cats map[uint]float64) {
	f, err := os.Open(modelPath)
	if err != nil {
		panic("cant open model")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err = enc.Decode(&m); err != nil {
		panic("cant decode model")
	}

	f, err = os.Open(catsPath)
	if err != nil {
		panic("cant open categories")
	}
	defer f.Close()

	enc = gob.NewDecoder(f)
	if err = enc.Decode(&cats); err != nil {
		panic("cant decode categories")
	}

	return m, cats
}

// Exists ...
func Exists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}
