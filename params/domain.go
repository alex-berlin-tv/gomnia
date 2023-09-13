package params

import "github.com/pasztorpisti/qs"

// Parameters for adding or updating a channel. Documentation is available [here].
//
// [here]: https://api.nexx.cloud/v3.1/manage/channels/:channelid/update
type Channel struct {
	// Required.
	Title       string `qs:"title"`
	Subtitle    string `qs:"subtitle,omitempty"`
	Refnr       string `qs:"refnr,omitempty"`
	Teaser      string `qs:"teaser,omitempty"`
	Description string `qs:"description,omitempty"`
	// If the channel shall be a sub channel of a parent channel, add the parent
	// Channel ID here
	Parent int `qs:"parent,omitempty"`
	// An optional sorting parameter for publically visible channels
	Pos   int    `qs:"pos,omitempty"`
	Color string `qs:"color,omitempty"`
}

func (c Channel) UrlEncode() (string, error) {
	return qs.Marshal(&c)
}
