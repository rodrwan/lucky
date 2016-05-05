package lucky

import (
	"fmt"
	"testing"

	"github.com/rodrwan/lucky/model"
)

func TestLucky(t *testing.T) {
	fmt.Println("Lucky")
	newModel := new(Lucky)
	fmt.Printf("\t> CatNum")
	if newModel.CatNum != nil {
		fmt.Println(" Fail")
		t.Errorf("CatNum should be nil, got %v", newModel.CatNum)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> CatStr")
	if newModel.CatStr != nil {
		fmt.Println(" Fail")
		t.Errorf("CatStr should be nil, got %v", newModel.CatStr)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> LabelsPath")
	if newModel.LabelsPath != "" {
		fmt.Println(" Fail")
		t.Errorf("LabelsPath should be \"\", got %v", newModel.LabelsPath)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> Model")
	if newModel.Model != nil {
		fmt.Println(" Fail")
		t.Errorf("Model should be nil, got %v", newModel.Model)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> TrainingDataPath")
	if newModel.TrainingDataPath != "" {
		fmt.Println(" Fail")
		t.Errorf("TrainingDataPath should be \"\", got %v", newModel.TrainingDataPath)
	}
	fmt.Println(" OK")
	fmt.Println()
}

func TestClassifier(t *testing.T) {
	fmt.Println("Lucky classifier")
	newModel := new(Lucky)
	newModel.Fit()
	desc := "ADIDAS PARQUE ARAUCO"
	best := newModel.Predict(desc)
	// fmt.Println(best)
	fmt.Printf("\t> Predict:")
	if best.ID != 13 {
		fmt.Println(" Fail")
		t.Errorf(" Should return ID equal to 13, got %d", best.ID)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> Model:")
	if newModel.Model == nil {
		fmt.Println(" Fail")
		t.Errorf(" Should not be equal to nil, got %v", newModel.Model)
	}
	fmt.Println(" OK")
	fmt.Printf("\t> Model.bin:")
	if !model.Exists("model.bin") {
		fmt.Println(" Fail")
		t.Errorf(" Should create model.bin, got %v", model.Exists("model.bin"))
	}
	fmt.Println(" OK")
	fmt.Printf("\t> categories.bin:")
	if !model.Exists("categories.bin") {
		fmt.Println(" Fail")
		t.Errorf("Should create categires.bin, got %v", model.Exists("categories.bin"))
	}
	fmt.Println(" OK")
	fmt.Println()
}
