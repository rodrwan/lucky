# Lucky

To train a new classifier you should create two files and pass it to params. After train the classifier, two files are created to speed up the training fase, (model.bin and categories.bin)

## Example files
### Training file:
```text
120#Pago Pesos TEF
26#COMISION MENSUAL POR MANTENCION
109#COMPRAS APL* ITUNES.COM/BILL
89#PAGO EN LINEA AUTOPISTA VESPUCIO SUR
89#Compra Internet COSTANERA NORTE
...
```
### Labels file:
```text
94#Ajuste por cierre de cuenta
12#Gimnasio
2#Bonos
86#Intereses
35#Jardín
13#Ropa Deportiva
...
```

## Basic model
```go
type Samples struct {
    Ngram    string
    Freq     float64
    Classes  map[uint]float64
    Probs    map[uint]float64
    Maximum  float64
    Minimum  float64
    Weighted bool
}
```

## Setup structure
```go
type Lucky struct {
    Model            map[string]*model.Samples
    CatNum           map[uint]float64
    CatStr           map[uint]string
    LabelsPath       string
    TrainingDataPath string
    URL              string
}
```

## Public response
```go
type BestCategory struct {
  ID    uint    // category id
  Name  string  // category name
  Score float64 // category probability
}
```

### Public Methods
```go
Fit() void
Predict(test string) (*model.BestCategory)
```

To see the classifier in action see `lucky_test.go`
