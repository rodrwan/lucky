package lucky

import (
	"fmt"
	"testing"
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
	fmt.Printf("\t> URL")
	if newModel.URL != "" {
		fmt.Println(" Fail")
		t.Errorf("URL should be \"\", got %v", newModel.URL)
	}
	fmt.Println(" OK")
}
