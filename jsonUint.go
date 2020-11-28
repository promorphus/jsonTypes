package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONUint struct {
	JsonValue uint
	JSONValidation
}

func (j JSONUint) Value() (value driver.Value, err error) {
	value = int64(j.JsonValue)
	return value, nil

}

func (j *JSONUint) Scan(src interface{}) (err error) {
	int64Val, ok := src.(int64)
	if !ok {
		err = errors.New("unable to scan")
		return
	}

	j.JsonValue = uint(int64Val)
	return
}

func (j JSONUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.JsonValue)
}

func (j *JSONUint) UnmarshalJSON(data []byte) (err error) {
	j.Set = true // unmarshalled, so the key was set

	var value uint
	if err = json.Unmarshal(data, &value); err != nil {
		return
	}
	j.JsonValue = value
	j.Valid = true
	return

}

func (j *JSONUint) isSet() bool {
	if j.Set == true {
		return true
	}
	return false
}
