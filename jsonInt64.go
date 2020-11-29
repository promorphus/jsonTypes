package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

type JSONInt struct {
	JsonValue int
	JSONValidation
}

func (j JSONInt) Value() (value driver.Value, err error) {
	value = int64(j.JsonValue)
	return
}

func (j *JSONInt) Scan(src interface{}) (err error) {
	var sv reflect.Value
	sv = reflect.ValueOf(src)

	switch {
	case sv.Kind() == reflect.Int:
		val, ok := src.(int)
		if !ok {
			err = errors.New("unable to scan")
			return
		}

		j.JsonValue = val
		return

	case sv.Kind() == reflect.Uint8:
		val, ok := src.([]uint8)
		if !ok {
			err = errors.New("unable to scan")
			return
		}

		var i int
		// i64, err = strconv.ParseInt(string(val), 10, 32)
		i, err = strconv.Atoi(string(val))
		if err != nil {
			log.WithFields(
				log.Fields{
					"Error": err.Error(),
				}).Error("Error scanning uint8")
			return
		}

		j.JsonValue = i
		return

	case sv.Kind() == reflect.Slice:

		t := reflect.TypeOf(src)
		switch t.Elem().Kind() {
		case reflect.Uint8:
			val, ok := src.([]uint8)
			if !ok {
				err = errors.New("unable to scan")
				return
			}

			var intint int
			intint, err = strconv.Atoi(string(val))
			// i, err = strconv.ParseInt(string(val), 10, 64)
			if err != nil {
				log.WithFields(
					log.Fields{
						"Error": err.Error(),
					}).Error("Error scanning []uint8")
				return
			}

			j.JsonValue = intint

		}
	}

	return
}

func (j JSONInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.JsonValue)
}

func (j *JSONInt) UnmarshalJSON(data []byte) (err error) {

	j.Set = true // unmarshalled, so the key was set

	if string(data) == "null" {
		j.Valid = false // The key was set to null
		return
	}

	var value int
	if err = json.Unmarshal(data, &value); err != nil {
		return
	}
	j.JsonValue = value
	j.Valid = true
	return
}

func (j *JSONInt) isSet() bool {
	if j.Set == true {
		return true
	}
	return false
}
