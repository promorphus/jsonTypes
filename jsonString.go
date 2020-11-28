package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONString struct {
	JsonValue string
	JSONValidation
}

func (j JSONString) Value() (value driver.Value, err error) {
	value = j.JsonValue
	return
}

func (j *JSONString) Scan(src interface{}) (err error) {
	val, ok := src.([]byte)
	if !ok {
		err = errors.New("unable to scan")
		return
	}

	j.JsonValue = string(val)
	return
}

func (j *JSONString) UnmarshalJSON(data []byte) (err error) {
	j.Set = true //unmarshalled, so the key was set

	var value string
	if err = json.Unmarshal(data, &value); err != nil {
		return
	}
	j.JsonValue = value
	j.Valid = true
	return
}

func (j JSONString) MarshalJSON() (value []byte, err error) {
	return json.Marshal(j.JsonValue)
}

func (j *JSONString) isSet() bool {
	if j.Set == true {
		return true
	}
	return false
}
