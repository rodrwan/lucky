package lucky

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/rodrwan/lucky/model"
)

func TestCatStr(t *testing.T) {
	newModel := new(Config)
	actual := newModel.CatStr
	if actual != nil {
		t.Errorf("CatStr should be nil, got %v", actual)
	}
}

func TestLabelsPath(t *testing.T) {
	newModel := new(Config)
	expected := ""
	actual := newModel.LabelsPath
	if actual != expected {
		t.Errorf("LabelsPath should be \"\", got %v", actual)
	}
}
func TestModel(t *testing.T) {
	newModel := new(Config)
	actual := newModel.Model
	if actual != nil {
		t.Errorf("Model should be nil, got %v", actual)
	}
}

func TestTrainingDataPath(t *testing.T) {
	newModel := new(Config)
	expected := ""
	actual := newModel.TrainingDataPath
	if actual != expected {
		t.Errorf("TrainingDataPath should be \"\", got %v", actual)
	}
}

func TestVerbose(t *testing.T) {
	newModel := new(Config)
	expected := false
	actual := newModel.Verbose
	if !actual == expected {
		t.Errorf("TrainingDataPath should be \"\", got %v", actual)
	}
}

func TestClassifier(t *testing.T) {
	newModel := new(Config)
	newModel.Verbose = false
	newModel.TrainingDataPath = "train.txt"
	newModel.Threshold = 0.0
	newModel.Fit()
	newModel.InvalidWords = []string{
	// "compra",
	// "pago",
	// "normal",
	// "cl",
	// "linea",
	// "parque",
	// "arauco",
	// "arauc",
	// "arau",
	// "mall",
	// "reina",
	// "condes",
	// "conde",
	// "alto",
	// "maipu",
	}

	desc := "ADIDAS PARQUE ARAUCO"
	best := newModel.Predict(desc)

	var expected uint = 13
	if best.ID != expected {
		t.Errorf(" Should return ID equal to %d, got %d", expected, best.ID)
	}
	if newModel.Model == nil {
		t.Errorf(" Should not be equal to nil, got %v", newModel.Model)
	}
	if !model.Exists("model.bin") {
		t.Errorf(" Should create model.bin, got %v", model.Exists("model.bin"))
	}

	var testCases = []struct {
		phrases    string
		categoryID uint
	}{
		{
			phrases:    "ADIDAS PARQUE ARAUCO",
			categoryID: 48,
		},
		{
			phrases:    "compra normal ripley parque arauco",
			categoryID: 48,
		},
		{
			phrases:    "compra normal paris parque arauco",
			categoryID: 48,
		},
		{
			phrases:    "compra normal falabella alto las condes",
			categoryID: 48,
		},
		{
			phrases:    "COMPRA SUBWAY METRO ESC",
			categoryID: 9,
		},
		{
			phrases:    "Giro en Cajero Automático",
			categoryID: 27,
		},
		{
			phrases:    "COMPRA NIKE SHOP ALTO LA",
			categoryID: 13,
		},
		{
			phrases:    "COMPRA LIQUIDOS LA REINA",
			categoryID: 99,
		},
		{
			phrases:    "COMPRA CASTANO LAS CONDE",
			categoryID: 100,
		},
		{
			phrases:    "COMPRA DONER HOUSE",
			categoryID: 9,
		},
		{
			phrases:    "PAC Seguro Frau 000005500832946",
			categoryID: 85,
		},
		{
			phrases:    "COMPRA PIZZA PIZZA COLON",
			categoryID: 9,
		},
		{
			phrases:    "COMPRA OK MARKET SAN PAS",
			categoryID: 8,
		},
		{
			phrases:    "COMPRA OK MARKET SUB CEN",
			categoryID: 8,
		},
		{
			phrases:    "COMPRA VIVE SNACK MOVIST",
			categoryID: 9,
		},
		{
			phrases:    "COMPRA BARBAZUL",
			categoryID: 7,
		},
		{
			phrases:    "COMPRA CASTANO LAS CONDE",
			categoryID: 100,
		},
		{
			phrases:    "COMPRA FUENTE CHILENA AP",
			categoryID: 9,
		},
		{
			phrases:    "PAGO EN LINEA VTR",
			categoryID: 70,
		},
		{
			phrases:    "PAGO EN LINEA AUTOPISTA COSTANERA NORTE",
			categoryID: 89,
		},
		{
			phrases:    "compra normal todo juegos cl",
			categoryID: 25,
		},
		{
			phrases:    "Transf otro Bco FINCIERO SPA",
			categoryID: 6,
		},
		{
			phrases:    "PAGO BELSPORT S.A. PARQUE ARAU",
			categoryID: 13,
		},
		{
			phrases:    "COPEC MAIPU 29336",
			categoryID: 56,
		},
		{
			phrases:    "JUMBO MAIPU SUPERMERCADO",
			categoryID: 8,
		},
	}

	good := 0
	for _, test := range testCases {
		best := newModel.Predict(test.phrases)
		if best.ID != test.categoryID {
			t.Errorf("Description: %s, has failed\n", test.phrases)
			t.Errorf("Category: %s\n", best.Name)
			t.Errorf("Probability: %f\n", best.Score)
			t.Errorf(" Should best.ID be equal to %d, got %d", test.categoryID, best.ID)
		} else {
			good++
		}
		acc := 100 * (float64(good) / float64(len(testCases)))
		fmt.Printf("\rAcc: %d/%d -> %03.2f%%", good, len(testCases), acc)
		time.Sleep(time.Second / 4)
	}
	fmt.Printf("\n\n")
	maxProcs := runtime.NumCPU()
	log.Printf("Available CPU: %d\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)

	testPath := "test.txt"
	fileBytes, _ := ioutil.ReadFile(testPath)
	fileAsString := string(fileBytes)
	lines := strings.Split(fileAsString, "\n")
	lines = lines[0:len(lines)]
	totalLines := len(lines)
	chunkSize := totalLines / maxProcs
	rest := totalLines % maxProcs
	wg := &sync.WaitGroup{}

	accuracy := 0.0
	for idx := 0; idx < maxProcs; idx++ {
		chunk := lines[chunkSize*(idx) : chunkSize*(idx+1)+rest]
		wg.Add(1)

		go func(chunk []string) {
			for _, line := range chunk {
				if len(line) == 0 {
					continue
				}
				splitedStr := strings.Split(line, "#")
				category, description := splitedStr[0], splitedStr[1]
				best := newModel.Predict(description)
				i, err := strconv.Atoi(category)
				if err != nil {
					continue
				}
				if best.ID == uint(i) {
					accuracy += 1.0
				}
			}
			wg.Done()
		}(chunk)
	}
	wg.Wait()
	fmt.Printf("Accuracy: %03.2f%%\n", 100*(accuracy/float64(len(lines))))
}
