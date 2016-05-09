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

func maxVote(votes map[uint]uint) (uint, bool) {
	var maximum uint
	equals := false
	var maxKey uint

	for key, count := range votes {
		last := maximum
		if count > maximum {
			maximum = count
			maxKey = key
		}

		if last == count {
			equals = true
		}
	}

	return maxKey, equals
}

func maxFreq(freq map[float64]uint) uint {
	var maxFreqKey uint = 1
	maxFreq := 0.0

	tmp := make(map[float64]uint)
	for key, value := range freq {
		tmp[key] += value
	}

	for key, value := range tmp {
		if key > maxFreq {
			maxFreq = key
			maxFreqKey = value
		}
	}

	return maxFreqKey
}

func bestOption(votes map[uint]uint, freq map[float64]uint) uint {
	// if the keys have the same number of votes, then we use
	// key frequency ratio to tiebreaker this equality
	maxKey, equals := maxVote(votes)

	if equals {
		return maxFreq(freq)
	}

	return maxKey
}

// Predict function to get best category
func Predict(m map[string]*Sample, test string, cats map[uint]string, words []string) *BestCategory {
	total := ngrams.Make(test, 3, words)
	freq := make(map[float64]uint)
	votes := make(map[uint]uint)

	for _, value := range total {
		sample := m[value]
		if sample != nil {
			maxKey := sample.maxKey()
			prob := sample.Classes[maxKey]
			// fmt.Printf("%s\t%f\t%f\t%d\t%s\n", value, prob, sample.Freq, maxKey, cats[maxKey])
			freq[prob] = maxKey
			votes[maxKey]++
		}
	}

	best := bestOption(votes, freq)

	return &BestCategory{
		ID:    best,
		Name:  cats[best],
		Score: 0.0, // TODO: implement score
	}
}
