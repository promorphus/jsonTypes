package jsonTypes

import (
	"strconv"
	"time"
)

type JSONTimeFormatted time.Time

func (j JSONTimeFormatted) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(j).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (j *JSONTimeFormatted) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(j) = time.Unix(q, 0)
	return nil
}


// Unix returns j as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
func (j JSONTimeFormatted) Unix() int64 {
	return time.Time(j).Unix()
}

// Time returns the JSONTimeFormatted as a time.Time instance in UTC
func (j JSONTimeFormatted) Time() time.Time {
	return time.Time(j).UTC()
}

// String returns j as a formatted string
func (j JSONTimeFormatted) String() string {
	return j.Time().String()
}
