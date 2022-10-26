package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/pasztorpisti/qs"
)

// Parameters for the reject ManagementAPI call. The documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#reject
type Reject struct {
	// A free text as reason.
	Reason string `qs:"reason,omitempty"`
	// How to handle the rejected media item after rejection.
	Action enums.ActionAfterRejection `qs:"action,omitempty"`
}

func (r Reject) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&r)
}
