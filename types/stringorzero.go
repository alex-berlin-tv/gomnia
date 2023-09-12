// Additional types which are used in requests to the API but are not built-ins
// and thus not implemented by the json package.
package types

import (
	"fmt"
	"strings"
)

// Workaround for Omnia's very handy feature of using an integer zero
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
