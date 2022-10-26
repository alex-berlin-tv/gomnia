package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/pasztorpisti/qs"
)

// Parameters for the approve ManagementAPI call. The documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#approve
type Approve struct {
	// Only valid on UGC Elements. If an Approval also claims the Media Item, the Item will
	// lose its UGC Flag and become an "standard" Media Item afterwards.
	AndClaim bool `qs:"andClaim,omitempty"`
	// A free Text as Reason
	Reason string `qs:"reason,omitempty"`
	// Restrict the Item to a dedicated Age Class
	RestrictToAge enums.AgeRestriction `qs:"restictToAge,omitempty"`
	// Flag the Item with certain Warning
	ContentModerationAspects string `qs:"contentModerationAspects,omitempty"`
}

func (a Approve) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&a)
}
