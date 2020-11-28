package jsonTypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	fc "github.com/promorphus/jsonTypes/functionalConstants"
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

func (j *JSONUUID) Generate() { // for backwards compatibility
	j.GenerateV4()
}

func (j *JSONUUID) GenerateV4() {
	j.Set = true
	newUUID, _ := uuid.NewV4()
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5DNS(name string) {
	j.Set = true
	newUUID := uuid.NewV5(fc.GetnamespaceDNS(), name)
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5URL(name string) {
	j.Set = true
	newUUID := uuid.NewV5(fc.GetnamespaceURL(), name)
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5OID(name string) {
	j.Set = true
	newUUID := uuid.NewV5(fc.GetnamespaceOID(), name)
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5X500(name string) {
	j.Set = true
	newUUID := uuid.NewV5(fc.GetnamespaceX500(), name)
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5(namespace string, name string) {
	namespaceUUID, err := uuid.FromString(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}

	j.Set = true
	newUUID := uuid.NewV5(namespaceUUID, name)
	j.JsonValue = newUUID
}

func (j *JSONUUID) Generatev5FailOnError(namespace string, name string) {
	namespaceUUID, err := uuid.FromString(namespace)
	if err != nil {
		log.Fatal(err.Error())
	}

	j.Set = true
	newUUID := uuid.NewV5(namespaceUUID, name)
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
