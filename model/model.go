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
				total := ngrams.Make(description, 3, words)
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
	// catsLen := float64(len(newCats.cats))

	for _, sample := range newSamples.m {
		sample.toProb()
		// sample.toTfIdf(catsLen)
	}

	SaveModel(modelPath, newSamples.m)
	return newSamples.m
}

// Prediction fase ...

// BestCategory ...
type BestCategory struct {
	ID    uint    // category id
	Name  string  // category name
	Score float64 // category prob
}

// Best ...
type Best struct {
	Prob  map[string]float64
	count uint
}

func (b *Best) setValues(prob float64, ngram string) {
	b.Prob[ngram] = prob
	b.count++
}

func maxVote(votes map[uint]*Vote) (uint, float64) {
	var maxKey uint
	var maximum float64

	for key, vote := range votes {
		if vote.Score > maximum {
			maximum = vote.Score
			maxKey = key
		}
	}

	return maxKey, maximum
}

// Vote ...
type Vote struct {
	Count uint
	Score float64
}

// Predict function to get best category
func Predict(m map[string]*Sample, test string, cats map[uint]string, words []string) *BestCategory {
	total := ngrams.Make(test, 3, words)
	votes := make(map[uint]*Vote)

	for _, value := range total {
		sample := m[value]
		if sample != nil {
			maxKey := sample.maxKey()
			// prob := sample.Tfidf[maxKey]
			if _, ok := votes[maxKey]; ok {
				votes[maxKey].Count++
				votes[maxKey].Score += (sample.Prob)
			} else {
				votes[maxKey] = &Vote{
					Count: 1,
					Score: sample.Prob,
				}
			}
		}
	}
	best, maxProb := maxVote(votes)

	return &BestCategory{
		ID:    best,
		Name:  cats[best],
		Score: maxProb,
	}
}
