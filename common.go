package jsonTypes

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type ParamList struct {
	Params []Param
}

type Param interface {
	isSet() bool
}

type JSONValidation struct {
	Valid bool
	Set   bool
}

func RequiredJsonParameters(params ...Param) (err error) {
	for _, param := range params {
		if param.isSet() != true {
			err = errors.New(fmt.Sprintf("%s", RequiredJsonParameterMissing))
			return
		}
	}
	return
}

func RequiredParameters(params ...interface{}) (err error) {
	for _, param := range params {
		typeOf := reflect.TypeOf(param)

		switch param.(type) {
		case string:
			if param == "" {
				err = errors.New(fmt.Sprintf("%s :: type: %s", MissingParameters, typeOf.String()))
				break
			}

		case uint:
			if param == 0 {
				err = errors.New(fmt.Sprintf("%s :: type: %s", MissingParameters, typeOf.String()))
				break
			}
		case int:
			if param == 0 {
				err = errors.New(fmt.Sprintf("%s :: type: %s", MissingParameters, typeOf.String()))
				break
			}
		case uuid.UUID:
			var uuidParam = param.(uuid.UUID)

			if uuidParam.String() == EMPTYUUID {
				err = errors.New(fmt.Sprintf("%s :: type: %s", MissingParameters, typeOf.String()))
				break
			}
		default:
			errorString := fmt.Sprintf("Unhandled type in requiredParameters -- %s", reflect.TypeOf(param).String())
			log.Errorf(errorString)
			err = errors.New(errorString)
			return
		}
	}

	return
}
