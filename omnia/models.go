package omnia

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Metadata part of an API response.
type ResponseMetadata struct {
	// The HTTP Status for this Call.
	Status int `json:"status"`
	// Version of the API.
	ApiVersion string `json:"apiversion,omitempty"`
	// The used HTTP Verb.
	Verb string `json:"verb"`
	// Internal Duration, needed to create the response.
	ProcessingTime float64 `json:"processingtime"`
	// The called Endoint and Parameter.
	CalledWith *string `json:"calledwith,omitempty"`
	// The `cfo` Parameter from the API Call.
	CalledFor *string `json:"calledfor,omitempty"`
	// The calling Domain ID.
	ForDomain *string `json:"fordomain,omitempty"`
	// The result was created by a Stage or Productive Server.
	FromStage *int `json:"fromstage,omitempty"`
	// If the Call uses deprecated Functionality, find here a Hint, what Attributes
	// should be changed.
	Notice *string `json:"notice"`
	// If the Call failed, a Hint for the Failure Reason.
	ErrorHint *string `json:"errorhint,omitempty"`
}

func (m ResponseMetadata) toMap() map[string]interface{} {
	return structToMap(m)
}

// Information on the paging of an result.
type ResponsePaging struct {
	// The Start of the Query Range.
	Start int `json:"start"`
	// The given maximal Item List Length.
	Limit int `json:"limit"`
	// The maximally available Number of Items.
	ResultCount int `json:"resultcount"`
}

func (p ResponsePaging) toMap() map[string]interface{} {
	return structToMap(p)
}

// The response of an nexxOmnia API call. As documented [here].
//
// [here]: https://api.docs.nexx.cloud/api-design/response-object
type Response struct {
	// Metadata.
	Metadata ResponseMetadata `json:"metadata"`
	// Acutal result of the call. Structure can vary widely.
	Result interface{} `json:"result"`
	// Optional information on the paging.
	Paging *ResponsePaging `json:"paging"`
}

func structToMap(data interface{}) map[string]interface{} {
	var rsl map[string]interface{}
	tmp, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(tmp, &rsl)
	return rsl
}
