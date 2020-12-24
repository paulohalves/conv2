# conv2
Biblioteca para conversão de dados em Go

## Utilização

### Import
```go
import (
	"github.com/paulohalves/conv2"
)
```

### XML to JSON
```go
sample = "<xml>test</xml>"

result, err := XMLToJSON(sample)
if err != nil {
    panic("XMLToJSON: Erro durante a conversão do XML para o JSON")
}

// result: {"xml": "test"}
```

### JSON to Map
```go
sample := "{\"test\": \"json\", \"ok\": true}"

result, err := JSONToMap(sample)
if err != nil {
    panic("JSONToMap: Erro durante a conversão do JSON para Map")
}
```

### Get String In Map
```go
sample := "{\"test\": \"json\", \"ok\": true}"
val, _ := JSONToMap(sample)

result, err := GetStringInMap(val, "test")
if err != nil {
    panic("GetStringInMap: Erro durante a localização da string")
}

// result: "json"
```

### Get Bool In Map
```go
sample := "{\"test\": \"json\", \"ok\": true}"
val, _ := JSONToMap(sample)

result, err := GetBoolInMap(val, "ok")
if err != nil {
    panic("GetBoolInMap: Erro durante a localização do bool")
}

// result: true
```

### Int To String
```go
sampleInt := 12

result, err := ToString(sampleInt)
if err != nil {
    panic("ToString - Int: Erro durante a conversão do Int para String")
}

// result: "12"
```

### Float To String
```go
sampleFloat := 12.34

result, err := ToString(sampleFloat)
if err != nil {
    panic("ToString - Float: Erro durante a conversão do Float para String")
}

// result: "12.34"
```

### Int To Bool
```go
sampleInt := 1

result, err := ToBool(sampleInt)
if err != nil {
    panic("ToBool - Int: Erro durante a conversão do Bool para Int")
}

// result: true
```

### Float To Bool
```go
sampleFloat := 12.34

result, err := ToBool(sampleFloat)
if err != nil {
    panic("ToBool - Float: Erro durante a conversão do Bool para Float")
}

// result: true
```

### String To Bool
```go
sampleString := "1"

result, err := ToBool(sampleString)
if err != nil {
    panic("ToBool - String: Erro durante a conversão do Bool para String")
}

// result: true
```