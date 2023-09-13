package params

import (
	"net/url"
)

// Provides parameters for an API call.
type QueryParameters interface {
	UrlEncode() (string, error)
}

// Set custom parameters using a string map.
type Custom map[string]string

func (c Custom) UrlEncode() (string, error) {
	values := url.Values{}
	for key, value := range c {
		values.Set(key, value)
	}
	return values.Encode(), nil
}
