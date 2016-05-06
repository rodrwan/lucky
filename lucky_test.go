package lucky

import (
	"fmt"
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
	newModel.Fit()
	desc := "ADIDAS PARQUE ARAUCO"
	best := newModel.Predict(desc)
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

	var stringCase = []string{
		"compra normal ripley parque arauco",
		"compra normal paris parque arauco",
		"compra normal falabella alto las condes",
		"COMPRA SUBWAY METRO ESC",
		"Giro en Cajero Automático",
		"COMPRA NIKE SHOP ALTO LA",
		"COMPRA LIQUIDOS LA REINA",
		"COMPRA CASTANO LAS CONDE",
		"COMPRA DONER HOUSE",
		"PAC Seguro Frau 000005500832946",
		"COMPRA PIZZA PIZZA COLON",
		"COMPRA OK MARKET SAN PAS",
		"COMPRA OK MARKET SUB CEN",
		"COMPRA VIVE SNACK MOVIST",
		"COMPRA BARBAZUL",
		"COMPRA CASTANO LAS CONDE",
		"COMPRA FUENTE CHILENA AP",
		"PAGO EN LINEA VTR",
		"PAGO EN LINEA AUTOPISTA COSTANERA NORTE",
		"compra normal todo juegos cl",
	}

	for _, desc := range stringCase {
		best := newModel.Predict(desc)
		fmt.Printf("Description: %s\n", desc)
		fmt.Printf("Category: %d -> %s\n", best.ID, best.Name)
		fmt.Println()
	}
}
