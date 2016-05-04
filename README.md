### Training data:
```text
120#Pago Pesos TEF
26#COMISION MENSUAL POR MANTENCION
109#COMPRAS APL* ITUNES.COM/BILL
89#PAGO EN LINEA AUTOPISTA VESPUCIO SUR
89#Compra Internet COSTANERA NORTE
...
```

### Basic model
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

### Setup structure
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

### Public response
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
