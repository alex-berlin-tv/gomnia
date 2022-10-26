package params

import (
	"net/url"
)

// Provides parameters for an API call.
type QueryParameters interface {
	UrlEncode(extra map[string]interface{}) (string, error)
}

type CustomParameters map[string]string

func (c CustomParameters) UrlEncode(extra map[string]interface{}) (string, error) {
	values := url.Values{}
	for key, value := range c {
		values.Set(key, value)
	}
	return values.Encode(), nil
}
