package conv2

import (
	"strings"
	"testing"
)

func TestXMLToJSON(t *testing.T) {

	var sample string

	//XMLToJSON
	sample = "<xml>test</xml>"

	val, err := XMLToJSON(sample)
	if err != nil {
		t.Errorf("XMLToJSON: Erro durante a conversão do XML para o JSON")
	}

	val = strings.TrimSpace(val)
	expected := `{"xml": "test"}`
	if val != expected {
		t.Errorf("XMLToJSON divergente -> esperado: %s | resultado: %s", expected, val)
	}

}

func TestJSONToMap(t *testing.T) {

	//JSONToMap
	sample := "{\"test\": \"json\", \"ok\": true}"

	val, err := JSONToMap(sample)
	if err != nil {
		t.Errorf("JSONToMap: Erro durante a conversão do JSON para Map")
	}

	// GetStringInMap
	s, err2 := GetStringInMap(val, "test")
	if err2 != nil {
		t.Errorf("GetStringInMap: Erro durante a localização da string")
	}

	s = strings.TrimSpace(s)
	expected := "json"
	if s != expected {
		t.Errorf("GetStringInMap divergente -> esperado: %s | resultado: %s", expected, s)
	}

	// GetBoolInMap
	b, err2 := GetBoolInMap(val, "ok")
	if err2 != nil {
		t.Errorf("GetBoolInMap: Erro durante a localização do bool")
	}

	expected2 := true
	if b != expected2 {
		t.Errorf("GetBoolInMap divergente -> esperado: %s | resultado: %s", expected, s)
	}
}

func TestToString(t *testing.T) {

	// Int
	sampleInt := 12

	val, err := ToString(sampleInt)
	if err != nil {
		t.Errorf("ToString - Int: Erro durante a conversão do Int para String")
	}

	expected := "12"
	if val != expected {
		t.Errorf("ToString - Int divergente -> esperado: %s | resultado: %s", expected, val)
	}

	// Float
	sampleFloat := 12.34

	val2, err := ToString(sampleFloat)
	if err != nil {
		t.Errorf("ToString - Float: Erro durante a conversão do Float para String")
	}

	expected2 := "12.34"
	if val2 != expected2 {
		t.Errorf("ToString - Float divergente -> esperado: %s | resultado: %q", expected2, val2)
	}

}

func TestToBool(t *testing.T) {

	// Int
	sampleInt := 1

	val, err := ToBool(sampleInt)
	if err != nil {
		t.Errorf("ToBool - Int: Erro durante a conversão do Bool para Int")
	}

	expected := true
	if val != expected {
		t.Errorf("ToBool - Int divergente -> esperado: %t | resultado: %t", expected, val)
	}

	// Float
	sampleFloat := 12.34

	val, err2 := ToBool(sampleFloat)
	if err2 != nil {
		t.Errorf("ToBool - Float: Erro durante a conversão do Bool para Float")
	}

	expected2 := true
	if val != expected2 {
		t.Errorf("ToBool - Float divergente -> esperado: %t | resultado: %t", expected, val)
	}

	// String
	sampleString := "1"

	val, err3 := ToBool(sampleString)
	if err3 != nil {
		t.Errorf("ToBool - String: Erro durante a conversão do Bool para String")
	}

	expected3 := true
	if val != expected3 {
		t.Errorf("ToBool - String divergente -> esperado: %t | resultado: %t", expected, val)
	}

}
