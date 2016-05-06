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
func Fit(path string, procs int) map[string]*Sample {
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
				splitedStr := strings.Split(line, "#")
				category, description := splitedStr[0], splitedStr[1]
				total := ngrams.Make(description, 3)
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
							Probs:    make(map[uint]float64),
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
	catsLen := float64(len(newCats.cats))

	for _, sample := range newSamples.m {
		sample.toTfIdf(catsLen)
		sample.scale()
	}

	SaveModel(modelPath, newSamples.m)
	return newSamples.m
}

func bestOption(votes map[uint]uint, freq map[float64]uint) uint {
	var maxFreqKey uint = 1
	var maximum uint
	var maxKey uint
	maxFreq := 0.0
	equals := true

	for key, value := range freq {
		if key > maxFreq {
			maxFreq = key
			maxFreqKey = value
		}
	}

	for key, count := range votes {
		if count > maximum {
			maximum = count
			maxKey = key
		}

		if count == maximum {
			equals = false
		}
	}

	if equals {
		return maxFreqKey
	}
	return maxKey
}

// BestCategory ...
type BestCategory struct {
	ID    uint    // category id
	Name  string  // category name
	Score float64 // category prob
}

// Best ...
type Best struct {
	Prob  float64
	Key   uint
	Ngram string
}

func (b *Best) resetValues() {
	b.Prob = 0.0
	b.Key = 0
	b.Ngram = ""
}

func (b *Best) setValues(prob float64, key uint, ngram string) {
	b.Prob = prob
	b.Key = key
	b.Ngram = ngram
}

// Predict function to get best category
func Predict(m map[string]*Sample, test string, cats map[uint]string) *BestCategory {
	total := ngrams.Make(test, 3)
	freq := make(map[float64]uint)
	votes := make(map[uint]uint)

	bestOpt := &Best{
		Prob:  0.0,
		Key:   0,
		Ngram: "",
	}

	for _, value := range total {
		sample := m[value]
		if sample != nil {
			maxKey := sample.maxKey()
			prob := sample.Classes[maxKey]
			bestOpt.setValues(prob, maxKey, value)
			freq[sample.Freq] = maxKey
		}
		if bestOpt.Prob > 0.0 {
			votes[bestOpt.Key]++
		}
		bestOpt.resetValues()
	}

	best := bestOption(votes, freq)

	return &BestCategory{
		ID:    best,
		Name:  cats[best],
		Score: bestOpt.Prob,
	}
}
