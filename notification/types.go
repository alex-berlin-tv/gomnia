package notification

import (
	"fmt"
	"strconv"
	"strings"
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

// Workaround for Omnia's *very handy* feature of using an integer zero
// when no string data is present.
type StringOrZero string

func (s StringOrZero) MarshallJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", s)), nil
}

func (s *StringOrZero) UnmarshalJSON(data []byte) (err error) {
	rsl := string(data)
	rsl = strings.TrimPrefix(rsl, "\"")
	rsl = strings.TrimSuffix(rsl, "\"")
	*(*string)(s) = rsl
	return nil
}
