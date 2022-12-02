package notification

import (
	"strconv"
	"time"
)

// Date and time represented as a UNIX-timestamp.
type UnixTS time.Time

func (t UnixTS) MarshallJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *UnixTS) UnmarshalJSON(data []byte) (err error) {
	value, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(value, 0)
	return nil
}
