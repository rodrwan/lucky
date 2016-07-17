package lucky

import (
	"errors"
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/rodrwan/lucky/model"
)

var (
	path   = flag.String("train", "training_data.txt", "Training data path")
	labels = flag.String("labels", "labels.txt", "Labels sample path")
	verb   = flag.Bool("verbose", false, "Verbose mode true to display")
)

// Config ...
type Config struct {
	Model            map[string]*model.Sample
	CatStr           map[uint]string
	LabelsPath       string
	TrainingDataPath string
	Verbose          bool
	AsPkg            bool
	InvalidWords     []string
	Threshold        float64
}

// Fit ...
func (config *Config) Fit() {
	if !config.AsPkg {
		flag.Parse()
	}
	if !config.Verbose {
		config.Verbose = *verb
	}
	if config.Verbose {
		log.Println(">> Init model.")
	}
	if config.LabelsPath == "" {
		config.LabelsPath = *labels
	}
	if config.Verbose {
		log.Printf("> Loading %s\n", config.LabelsPath)
	}
	if config.TrainingDataPath == "" {
		config.TrainingDataPath = *path
	}
	if config.Verbose {
		log.Printf("> Loading %s\n", config.TrainingDataPath)
	}

	err := config.getCategories()
	if err != nil {
		log.Fatalln(err)
	}
	if config.Verbose {
		log.Println(">> Fit model.")
	}

	// enable all cpu to speedup model fit
	maxProcs := runtime.NumCPU()
	// log.Printf("Available CPU: %d\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)

	start := time.Now()
	m := model.Fit(config.TrainingDataPath, maxProcs, config.InvalidWords)
	elapsed := time.Since(start)
	if config.Verbose {
		log.Printf("Fit took %s\n\n", elapsed)
	}

	config.Model = m
	// log.Println(">> Ready to categorize.")
	// just leave 1 cpu to the rest of work
	runtime.GOMAXPROCS(1)
}

// Predict ...
func (config *Config) Predict(test string) (res *model.BestCategory) {
	m := config.Model
	cats := config.CatStr
	res = model.Predict(m, test, cats, config.InvalidWords, config.Threshold)
	return
}

func (config *Config) getCategories() error {
	if model.Exists(config.LabelsPath) {
		config.CatStr = model.LoadLabels(config.LabelsPath)
		return nil
	}

	return errors.New("can't load labels file")
}
