package util

import (
	"bytes"
	"encoding/json"
)

// ToJson 将对象转换为JSON
func ToJson(obj interface{}) string {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func JsonToMap(json string) (map[string]interface{}, error) {
	var resultMap map[string]interface{}
	err := ReadJsonByByte([]byte(json), &resultMap)
	return resultMap, err
}

func ReadJsonByByte(body []byte, obj interface{}) error {
	mdz := json.NewDecoder(bytes.NewBuffer(body))

	mdz.UseNumber()
	err := mdz.Decode(obj)

	if err != nil {
		return err
	}
	return nil
}
