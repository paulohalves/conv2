package conv2

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
	x2j "github.com/basgys/goxml2json"
	"github.com/cstockton/go-conv"
)

// XMLToJSON -
func XMLToJSON(xmlString string) (string, error) {

	xml := strings.NewReader(xmlString)

	json, err := x2j.Convert(xml)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return json.String(), nil
}

// JSONToMap -
func JSONToMap(jsonString string) (*gabs.Container, error) {

	jsonParsed, err := gabs.ParseJSON([]byte(jsonString))
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return jsonParsed, nil
}

// ArrayJSONToMap -
func ArrayJSONToMap(data string) ([]map[string]interface{}, error) {

	var ret []map[string]interface{}

	b := []byte(data)

	var m interface{}

	json.Unmarshal(b, &m)

	switch v := m.(type) {
	case []interface{}:
		for _, x := range v {
			ret = append(ret, x.(map[string]interface{}))
		}
	}

	return ret, nil
}

// GetInMap -
func GetInMap(data *gabs.Container, info string) interface{} {
	return data.Path(info)
}

// GetStringInMap -
func GetStringInMap(data *gabs.Container, info string) (string, error) {

	var ret string

	ret, err := conv.String(data.Path(info))
	if err != nil {
		return "", errors.New(err.Error())
	}

	ret = strings.ReplaceAll(ret, "\"", "")

	return ret, nil
}

// GetBoolInMap -
func GetBoolInMap(data *gabs.Container, info string) (bool, error) {

	var ret bool

	item, err := conv.String(data.Path(info))
	if err != nil {
		return false, errors.New(err.Error())
	}

	ret, err2 := conv.Bool(strings.Replace(item, "\"", "", -1))
	if err2 != nil {
		return false, errors.New(err.Error())
	}

	return ret, nil
}

// ToString -
func ToString(data interface{}) (string, error) {

	var ret string

	switch data.(type) {
	case int:
		ret = fmt.Sprintf("%v", data)
	case float64:
		ret = fmt.Sprintf("%v", data)
	default:
		return "", errors.New("Unknown")
	}

	return ret, nil
}

// ToBool -
func ToBool(data interface{}) (bool, error) {

	var ret bool

	switch data.(type) {
	case int:
		val := InterfaceToInt(data)

		if val != 0 {
			ret = true
		}
	case float64:
		val := InterfaceToFloat64(data)

		if val != 0 {
			ret = true
		}
	case string:
		val := InterfaceToString(data)

		if val != "" {
			ret = true
		}
	default:
		return false, errors.New("Unknown")
	}

	return ret, nil
}

// InterfaceToInt -
func InterfaceToInt(data interface{}) int {
	return data.(int)
}

// InterfaceToFloat64 -
func InterfaceToFloat64(data interface{}) float64 {
	return data.(float64)
}

// InterfaceToString -
func InterfaceToString(data interface{}) string {
	return data.(string)
}

// DebugStruct -
func DebugStruct(data interface{}) (string, error) {

	empJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string(empJSON), nil
}
