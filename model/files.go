package model

import (
	"bufio"
	"encoding/gob"
	"os"
	"strconv"
	"strings"
)

// SaveModel ...
func SaveModel(file string, m map[string]*Sample) {
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
func Load(modelPath string) (m map[string]*Sample) {
	f, err := os.Open(modelPath)
	if err != nil {
		panic("cant open model")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err = enc.Decode(&m); err != nil {
		panic("cant decode model")
	}

	return m
}

// Exists ...
func Exists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}

// LoadLabels ...
func LoadLabels(subCatPath string) map[uint]string {
	f, err := os.Open(subCatPath)
	if err != nil {
		panic("cant open labels file")
	}
	defer f.Close()

	m := make(map[uint]string)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		splitedStr := strings.Split(line, "#")
		id, name := splitedStr[0], splitedStr[1]
		i, err := strconv.Atoi(id)
		if err != nil {
			panic("cannot convert string to int")
		}
		m[uint(i)] = name
	}

	return m
}
