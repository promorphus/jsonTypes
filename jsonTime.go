package jsonTypes

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"time"
)

type JSONTimeFormatted struct {
	JsonValue time.Time
	JSONValidation
}


func (j JSONTimeFormatted) Value() (value driver.Value, err error) {
	return j.JsonValue, nil

}

func (j *JSONTimeFormatted) Scan(src interface{}) (err error) {
	t, ok := src.(time.Time)
	if !ok {
		err = errors.New("unable to scan")
		return
	}

	j.JsonValue = t
	return
}


func (j JSONTimeFormatted) MarshalJSON() ([]byte, error) {
	s := []byte(strconv.FormatInt(time.Time(j.JsonValue).Unix(), 10))
	return s, nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (j *JSONTimeFormatted) UnmarshalJSON(s []byte) (err error) {
	// r := string(s)
	// q, err := strconv.ParseInt(r, 10, 64)
	// if err != nil {
	// 	return err
	// }
	// *(*time.Time)(j.Time()) = time.Unix(q, 0)

	r := strings.Trim(string(s), "\"")
	j.JsonValue, err =  time.Parse("2006-01-02", r)
	if err != nil {
		return
	}

	return nil
}


// Unix returns j as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// func (j JSONTimeFormatted) Unix() int64 {
// 	return time.Time(j).Unix()
// }
//
// // Time returns the JSONTimeFormatted as a time.Time instance in UTC
// func (j JSONTimeFormatted) Time() time.Time {
// 	return time.Time(j).UTC()
// }

// String returns j as a formatted string
// func (j JSONTimeFormatted) String() string {
// 	return j.Time().String()
// }
