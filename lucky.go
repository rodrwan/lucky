package lucky

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/lucky/db"
	"github.com/rodrwan/lucky/model"
)

var (
	path   = flag.String("train", "training_data.txt", "Training data path")
	labels = flag.String("labels", "labels.txt", "Labels sample path")
	dbURL  = flag.String("dburl", "database-url", "DB url connection")
)

// Lucky ...
type Lucky struct {
	Model            map[string]*model.Samples
	CatNum           map[uint]float64
	CatStr           map[uint]string
	LabelsPath       string
	TrainingDataPath string
	URL              string
}

// Fit ...
func (newModel *Lucky) Fit() {
	flag.Parse()
	fmt.Println(">> Init model.")
	if newModel.LabelsPath == "" {
		newModel.LabelsPath = *labels
	}
	fmt.Printf("> Loading %s\n", newModel.LabelsPath)
	if newModel.TrainingDataPath == "" {
		newModel.TrainingDataPath = *path
	}
	fmt.Printf("> Loading %s\n", newModel.TrainingDataPath)
	if newModel.URL == "" {
		newModel.URL = *dbURL
	}
	fmt.Printf("> Getting data from %s...\n", newModel.URL[0:10])
	newModel.getCategories()

	fmt.Println(">> Fit model.")
	start := time.Now()
	m, c := model.Fit(newModel.TrainingDataPath)
	elapsed := time.Since(start)
	log.Printf("Fit took %s\n\n", elapsed)

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

func (newModel *Lucky) getCategories() {
	if model.Exists(newModel.LabelsPath) {
		newModel.CatStr = db.Load(newModel.LabelsPath)
	}

	start := time.Now()
	conn, _ := sqlx.Open("postgres", newModel.URL)
	cats, _ := db.Categories(conn)
	elapsed := time.Since(start)
	log.Printf("Load data took %s", elapsed)
	newModel.CatStr = cats
	return
}

func main() {
	flag.Parse()
	newModel := new(Lucky)
	newModel.Fit()
	desc := "ADIDAS PARQUE ARAUCO"
	best := newModel.Predict(desc)
	fmt.Println(best)
}
