package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/pasztorpisti/qs"
)

// Parameters for the byquery MediaAPI call. The documentation is available [here].
//
// [here]: https://api.docs.nexx.cloud/media-api/endpoints/media-endpoint#byquery
type ByQuery struct {
	Basic
	// Defines the Way, the Query is executed. Fore more results, "classicwithor"
	// is optimal. For a Lucene Search with Relevance, use "fulltext".
	QueryMode enums.QueryMode `qs:"queryMode,omitempty"`
	// A comma seperated List of Attributes, to search within. If omitted, the
	// Search will use all available Text Attributes.
	QueryFields []string `qs:"queryFields,omitempty"`
	// Skip Results with a Query Score lower than the given Value. Only usefull
	// for querymode "fulltext".
	MinimalQueryScore int `qs:"minimalQueryScore,omitempty"`
	// By default, the Query will only return Results on  full Words. If also
	// Substring Matches shall be returned, set this Parameter to 1. Only usefull,
	// if querymode is not "fulltext".
	IncludeSubstringMatches bool `qs:"includeSubstringMatches,omitempty"`
	// By default, the Query will only return Results on full Words. If also
	// Substring Matches shall be returned, set this Parameter to 1. Only usefull,
	// if querymode is not "fulltext".
	SkipReporting bool `qs:"skipReporting,omitempty"`
}

func (b ByQuery) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&b)
}
