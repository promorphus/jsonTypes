package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONBool struct {
	JsonValue bool
	JSONValidation
}

func (j JSONBool) Value() (value driver.Value, err error) {
	value = j.JsonValue
	return
}

func (j *JSONBool) Scan(src interface{}) (err error) {
	bv, err := driver.Bool.ConvertValue(src)
	if err != nil {
		return
	}
	j.JsonValue = bv.(bool)
	return
}

func (j JSONBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.JsonValue)
}

func (j *JSONBool) UnmarshalJSON(data []byte) (err error) {
	j.Set = true //unmarshalled, so the key was set

	var value bool
	if err = json.Unmarshal(data, &value); err != nil {
		return
	}
	j.JsonValue = value
	j.Valid = true
	return
}

func (j *JSONBool) isSet() bool {
	if j.Set == true {
		return true
	}
	return false
}
