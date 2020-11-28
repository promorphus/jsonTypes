package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/gofrs/uuid"
)

type JSONUUID struct {
	JsonValue uuid.UUID
	JSONValidation
}

func (j JSONUUID) IsUUIDNull() bool {
	if j.JsonValue.String() == EMPTYUUID {
		return true
	}
	return false
}

func (j *JSONUUID) Generate() {
	j.Set = true
	newUUID, _ := uuid.NewV4()
	j.JsonValue = newUUID
}

func (j JSONUUID) Value() (value driver.Value, err error) {
	if j.JsonValue.String() == EMPTYUUID {
		return "", nil
	} else {
		return j.JsonValue.String(), nil
	}
}

func (j *JSONUUID) Scan(src interface{}) (err error) {
	return j.JsonValue.Scan(src)

}

func (j JSONUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.JsonValue)
}

func (j *JSONUUID) UnmarshalJSON(data []byte) (err error) {
	j.Set = true // unmarshalled, so the key was set

	var value uuid.UUID
	if err = json.Unmarshal(data, &value); err != nil {
		return
	}
	j.JsonValue = value
	j.Valid = true
	return
}

func (j JSONUUID) isSet() bool {
	if j.Set == true {
		return true
	}
	return false
}
