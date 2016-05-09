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

// Lucky ...
type Lucky struct {
	Model            map[string]*model.Sample
	CatStr           map[uint]string
	LabelsPath       string
	TrainingDataPath string
	Verbose          bool
	AsPkg            bool
	InvalidWords     []string
}

// Fit ...
func (newModel *Lucky) Fit() {
	if !newModel.AsPkg {
		flag.Parse()
	}
	if !newModel.Verbose {
		newModel.Verbose = *verb
	}
	if newModel.Verbose {
		log.Println(">> Init model.")
	}
	if newModel.LabelsPath == "" {
		newModel.LabelsPath = *labels
	}
	if newModel.Verbose {
		log.Printf("> Loading %s\n", newModel.LabelsPath)
	}
	if newModel.TrainingDataPath == "" {
		newModel.TrainingDataPath = *path
	}
	if newModel.Verbose {
		log.Printf("> Loading %s\n", newModel.TrainingDataPath)
	}

	err := newModel.getCategories()
	if err != nil {
		log.Fatalln(err)
	}
	if newModel.Verbose {
		log.Println(">> Fit model.")
	}

	// enable all cpu to speedup model fit
	maxProcs := runtime.NumCPU()
	log.Printf("Available CPU: %d\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)

	start := time.Now()
	m := model.Fit(newModel.TrainingDataPath, maxProcs, newModel.InvalidWords)
	elapsed := time.Since(start)
	if newModel.Verbose {
		log.Printf("Fit took %s\n\n", elapsed)
	}

	newModel.Model = m
	log.Println(">> Ready to categorize.")
	// just leave 1 cpu to the rest of work
	runtime.GOMAXPROCS(1)
}

// Predict ...
func (newModel *Lucky) Predict(test string) (res *model.BestCategory) {
	m := newModel.Model
	cats := newModel.CatStr
	res = model.Predict(m, test, cats, newModel.InvalidWords)
	return
}

func (newModel *Lucky) getCategories() error {
	if model.Exists(newModel.LabelsPath) {
		newModel.CatStr = model.LoadLabels(newModel.LabelsPath)
		return nil
	}

	return errors.New("can't load labels file")
}
