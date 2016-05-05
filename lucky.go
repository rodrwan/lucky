package lucky

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rodrwan/lucky/db"
	"github.com/rodrwan/lucky/model"
)

var (
	path   = flag.String("train", "training_data.txt", "Training data path")
	labels = flag.String("labels", "labels.txt", "Labels sample path")
	verb   = flag.Bool("verbose", false, "Verbose mode true to display")
)

// Lucky ...
type Lucky struct {
	Model            map[string]*model.Samples
	CatNum           map[uint]float64
	CatStr           map[uint]string
	LabelsPath       string
	TrainingDataPath string
	verbose          bool
}

// Fit ...
func (newModel *Lucky) Fit() {
	flag.Parse()
	if !newModel.verbose {
		newModel.verbose = *verb
	}
	if newModel.verbose {
		fmt.Println(">> Init model.")
	}
	if newModel.LabelsPath == "" {
		newModel.LabelsPath = *labels
	}
	if newModel.verbose {
		fmt.Printf("> Loading %s\n", newModel.LabelsPath)
	}
	if newModel.TrainingDataPath == "" {
		newModel.TrainingDataPath = *path
	}
	if newModel.verbose {
		fmt.Printf("> Loading %s\n", newModel.TrainingDataPath)
	}

	err := newModel.getCategories()
	if err != nil {
		log.Fatalln(err)
	}
	if newModel.verbose {
		fmt.Println(">> Fit model.")
	}

	start := time.Now()
	m, c := model.Fit(newModel.TrainingDataPath)
	elapsed := time.Since(start)
	if newModel.verbose {
		log.Printf("Fit took %s\n\n", elapsed)
	}

	newModel.Model = m
	newModel.CatNum = c
	fmt.Println(">> Ready to categorize.")
}

// Predict ...
func (newModel *Lucky) Predict(test string) (res *model.BestCategory) {
	m := newModel.Model
	cats := newModel.CatStr
	c := newModel.CatNum
	res = model.Predict(m, test, cats, c)
	return
}

func (newModel *Lucky) getCategories() error {
	start := time.Now()
	if model.Exists(newModel.LabelsPath) {
		newModel.CatStr = db.Load(newModel.LabelsPath)
		elapsed := time.Since(start)
		if *verb {
			log.Printf("Load labels took %s", elapsed)
		}
		return nil
	}

	return errors.New("can't load labels file")
}
