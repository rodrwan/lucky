package lucky

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/rodrwan/lucky/model"
)

func TestCatStr(t *testing.T) {
	newModel := new(Lucky)
	actual := newModel.CatStr
	if actual != nil {
		t.Errorf("CatStr should be nil, got %v", actual)
	}
}

func TestLabelsPath(t *testing.T) {
	newModel := new(Lucky)
	expected := ""
	actual := newModel.LabelsPath
	if actual != expected {
		t.Errorf("LabelsPath should be \"\", got %v", actual)
	}
}
func TestModel(t *testing.T) {
	newModel := new(Lucky)
	actual := newModel.Model
	if actual != nil {
		t.Errorf("Model should be nil, got %v", actual)
	}
}

func TestTrainingDataPath(t *testing.T) {
	newModel := new(Lucky)
	expected := ""
	actual := newModel.TrainingDataPath
	if actual != expected {
		t.Errorf("TrainingDataPath should be \"\", got %v", actual)
	}
}

func TestVerbose(t *testing.T) {
	newModel := new(Lucky)
	expected := false
	actual := newModel.Verbose
	if !actual == expected {
		t.Errorf("TrainingDataPath should be \"\", got %v", actual)
	}
}

func TestClassifier(t *testing.T) {
	newModel := new(Lucky)
	newModel.Verbose = true
	newModel.TrainingDataPath = "training_data2.txt"

	newModel.Fit()
	// cats := newModel.CatStr
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
	}
	desc := "ADIDAS PARQUE ARAUCO"
	best := newModel.Predict(desc)
	fmt.Println(best.Name)
	if best.ID != 13 {
		t.Errorf(" Should return ID equal to 13, got %d", best.ID)
	}
	if newModel.Model == nil {
		t.Errorf(" Should not be equal to nil, got %v", newModel.Model)
	}
	if !model.Exists("model.bin") {
		t.Errorf(" Should create model.bin, got %v", model.Exists("model.bin"))
	}
	fmt.Println()

	// var testCases = []struct {
	// 	phrases    string
	// 	categoryID uint
	// }{
	// 	{
	// 		phrases:    "compra normal ripley parque arauco",
	// 		categoryID: 48,
	// 	},
	// 	{
	// 		phrases:    "compra normal paris parque arauco",
	// 		categoryID: 48,
	// 	},
	// 	{
	// 		phrases:    "compra normal falabella alto las condes",
	// 		categoryID: 48,
	// 	},
	// 	{
	// 		phrases:    "COMPRA SUBWAY METRO ESC",
	// 		categoryID: 9,
	// 	},
	// 	{
	// 		phrases:    "Giro en Cajero Automático",
	// 		categoryID: 27,
	// 	},
	// 	{
	// 		phrases:    "COMPRA NIKE SHOP ALTO LA",
	// 		categoryID: 13,
	// 	},
	// 	{
	// 		phrases:    "COMPRA LIQUIDOS LA REINA",
	// 		categoryID: 99,
	// 	},
	// 	{
	// 		phrases:    "COMPRA CASTANO LAS CONDE",
	// 		categoryID: 162,
	// 	},
	// 	{
	// 		phrases:    "COMPRA DONER HOUSE",
	// 		categoryID: 9,
	// 	},
	// 	{
	// 		phrases:    "PAC Seguro Frau 000005500832946",
	// 		categoryID: 85,
	// 	},
	// 	{
	// 		phrases:    "COMPRA PIZZA PIZZA COLON",
	// 		categoryID: 9,
	// 	},
	// 	{
	// 		phrases:    "COMPRA OK MARKET SAN PAS",
	// 		categoryID: 8,
	// 	},
	// 	{
	// 		phrases:    "COMPRA OK MARKET SUB CEN",
	// 		categoryID: 8,
	// 	},
	// 	{
	// 		phrases:    "COMPRA VIVE SNACK MOVIST",
	// 		categoryID: 100,
	// 	},
	// 	{
	// 		phrases:    "COMPRA BARBAZUL",
	// 		categoryID: 7,
	// 	},
	// 	{
	// 		phrases:    "COMPRA CASTANO LAS CONDE",
	// 		categoryID: 162,
	// 	},
	// 	{
	// 		phrases:    "COMPRA FUENTE CHILENA AP",
	// 		categoryID: 9,
	// 	},
	// 	{
	// 		phrases:    "PAGO EN LINEA VTR",
	// 		categoryID: 70,
	// 	},
	// 	{
	// 		phrases:    "PAGO EN LINEA AUTOPISTA COSTANERA NORTE",
	// 		categoryID: 89,
	// 	},
	// 	{
	// 		phrases:    "compra normal todo juegos cl",
	// 		categoryID: 25,
	// 	},
	// 	{
	// 		phrases:    "Transf otro Bco FINCIERO SPA",
	// 		categoryID: 82,
	// 	},
	// 	{
	// 		phrases:    "PAGO BELSPORT S.A. PARQUE ARAU",
	// 		categoryID: 13,
	// 	},
	// 	{
	// 		phrases:    "COPEC MAIPU 29336",
	// 		categoryID: 56,
	// 	},
	// 	{
	// 		phrases:    "JUMBO MAIPU SUPERMERCADO",
	// 		categoryID: 8,
	// 	},
	// }

	// for _, test := range testCases {
	// 	best := newModel.Predict(test.phrases)
	// 	fmt.Printf("Description: %s\n", test.phrases)
	// 	fmt.Printf("Category: %d -> %s\n", best.ID, best.Name)
	// 	if best.ID != test.categoryID {
	// 		t.Errorf("Description %s has failed\n", test.phrases)
	// 		t.Errorf("Category: %s\n", best.Name)
	// 		t.Errorf(" Should best.ID be equal to %d, got %d", test.categoryID, best.ID)
	// 	}
	// 	fmt.Println()
	// }

	path := "training_data2.txt"
	fileBytes, _ := ioutil.ReadFile(path)
	fileAsString := string(fileBytes)
	lines := strings.Split(fileAsString, "\n")
	lines = lines[0 : len(lines)-1]

	accuracy := 0.0
	for _, line := range lines {
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
		// else {
		// 	fmt.Printf("Description: %s\nExpected: %s\nActual: %s\n\n", description, cats[best.ID], cats[uint(i)])
		// }
	}

	fmt.Printf("Accuracy: %03.2f%%\n", 100*(accuracy/float64(len(lines))))
}
