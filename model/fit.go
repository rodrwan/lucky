package model

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/rodrwan/lucky/ngrams"
)

// Fit create a map of ngrams from file
func Fit(path string, procs int, words []string) map[string]*Sample {
	modelPath := "model.bin"
	if Exists(modelPath) {
		m := Load(modelPath)
		return m
	}

	newSamples := NewSet()
	newCats := NewCats()

	inFile, _ := os.Open(path)
	defer inFile.Close()

	fileBytes, _ := ioutil.ReadFile(path)
	fileAsString := string(fileBytes)
	lines := strings.Split(fileAsString, "\n")
	lines = lines[0 : len(lines)-1]
	totalLines := len(lines)
	chunkSize := totalLines / procs
	rest := totalLines % procs
	wg := &sync.WaitGroup{}

	for idx := 0; idx < procs; idx++ {
		chunk := lines[chunkSize*(idx) : chunkSize*(idx+1)+rest]
		wg.Add(1)

		go func(chunk []string) {
			for _, line := range chunk {
				if len(line) == 0 {
					continue
				}
				splitedStr := strings.Split(line, "#")
				category, description := splitedStr[0], splitedStr[1]
				total := ngrams.MakeN(description, 3, words)
				i, err := strconv.Atoi(category)
				if err != nil {
					continue
				}

				for _, value := range total {
					newCats.Add(uint(i))
					ok := newSamples.Has(value)
					if ok {
						newSamples.Update(value, uint(i))
					} else {
						// init
						sample := &Sample{
							Ngram:    value,
							Freq:     1.0,
							Classes:  make(map[uint]float64),
							Tfidf:    make(map[uint]float64),
							Maximum:  0.0,
							Minimum:  100000000.0,
							Weighted: false,
						}
						newSamples.Add(value, sample, uint(i))
					}
				}
			}
			wg.Done()
		}(chunk)
	}
	wg.Wait()

	for _, sample := range newSamples.m {
		sample.toProb()
	}

	SaveModel(modelPath, newSamples.m)
	return newSamples.m
}
