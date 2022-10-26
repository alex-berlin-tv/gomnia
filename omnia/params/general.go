package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/pasztorpisti/qs"
)

// General parameters for an MediaAPI call. Documentation is available [here].
//
// [here]: https://api.docs.nexx.cloud/media-api/usage
type GeneralParameters struct {
	// Restrict result to elements, created after the given time
	CreatedAfter string `qs:"createdAfter,omitempty"`
	// Restrict result to elements, modified after the given time
	ModifiedAfter string `qs:"modifiedAfter,omitempty"`
	// Restrict result to elements, published after the given time
	PublishedAfter string `qs:"publishedAfter,omitempty"`
	// Restrict result set to items in this channel
	Channel int `qs:"channel,omitempty"`
	// If the target channel is a main channel, and the contents of its
	// Subchannels shall also be included, set this Parameter to 1.
	RespectChannelHierarchy enums.Bool `qs:"respectChannelHierarchy,omitempty"`
	// Restrict result set to items in this format
	Format int `qs:"format,omitempty"`
	// Restrict result set to items in this category
	Category int `qs:"category,omitempty"`
	// Restrict result set to items in this genre (only for video, playlist,
	// Series, Audio and Audio Album)
	Genre int `qs:"genre,omitempty"`
	// Many media items have certain types to define their purpose.  if
	// necessary, you can filter by this Enum.
	ItemType string `qs:"type,omitempty"`
	// Many media items have a certain contenttype that defines some
	// Characteristics.  If necessary, you can filter by this Enum.
	ContentType enums.ContentType `qs:"contentType,omitempty"`
	// Restricts the result to media from the given country
	Country string `qs:"country,omitempty"`
	// Restrict result set to items with defined age level < 13 and without
	// Content Moderation Aspects
	NoExplicit enums.Bool `qs:"noExplicit,omitempty"`
	// Restrict result set to items without content moderation aspects
	NoContentModerationHints enums.Bool `qs:"noContentModerationHints,omitempty"`
	// Restrict result set to items with maximally the given age level
	MaxAge enums.AgeRestriction `qs:"maxAge,omitempty"`
	// Restrict result set to items with at least the given age level
	MinAge enums.AgeRestriction `qs:"minAge,omitempty"`
	// Restrict the result to media with the given height (alternatively, this
	// Parameter also accepts a numeric Value, which will be mapped to the Media
	// Height)
	Dimension enums.Dimension `qs:"dimension,omitempty"`
	// Restrict the result to media with the given orientation
	Orientation enums.Orientation `qs:"orientation,omitempty"`
	// If set to 1, only media in hdr quality will be returned
	OnlyHdr enums.Bool `qs:"onlyHDR,omitempty"`
	// The given items wont be included into the result set
	ExcludeItems []int `qs:"excludeItems,omitempty"`
	// Add items, uploaded by the community, to the result set
	IncludeUGC enums.Bool `qs:"includeUGC,omitempty"`
	// Restrict result set to items, uploaded by the community
	OnlyUGC enums.Bool `qs:"onlyUGC,omitempty"`
	// Also include files, that originate not by nexxomnia, but a partner
	// Provider
	IncludeRemote enums.Bool `qs:"includeRemote,omitempty"`
	// Only include files, that originate not by nexxomnia, but a partner
	// Provider
	OnlyRemote enums.Bool `qs:"onlyRemote,omitempty"`
	// Also include media items, that are marked as not listable.  this
	// Parameter should be used only in very specific Usecases.
	IncludeNotListables enums.Bool `qs:"includeNotListables,omitempty"`
	// Only valid for container calls with addchildmedia parameter.  this
	// Parameter will add currently invalid Elements to the Child Listing.  To
	// make this Parameter work, an active eternal Session must be used.
	IncludeInvalidChildMedia enums.Bool `qs:"includeInvalidChildMedia,omitempty"`
	// Also include media, that are not valid yet, but will be in the near
	// Future and allow Premiere Functionality
	IncludePremieres enums.Bool `qs:"includePremieres,omitempty"`
	// Restrict result set to items with payment attributes
	OnlyPay enums.Bool `qs:"onlyPay,omitempty"`
	// Restrict result set to items with payment attributes and premium payment
	// Attributes
	OnlyPremiumPay enums.Bool `qs:"onlyPremiumPay,omitempty"`
	// Restrict result set to items with payment attributes and standard payment
	// Attributes
	OnlyStandardPay enums.Bool `qs:"onlyStandardPay,omitempty"`
	// Only possible if addpublishingdetails is active.  if set to 1, only
	// planned Elements will be returned.
	OnlyPlanned enums.Bool `qs:"onlyPlanned,omitempty"`
	// Only possible, if addpublishingdetails is active.  if set to 1, only
	// unpublished Elements will be returned.
	OnlyInactive enums.Bool `qs:"onlyInactive,omitempty"`
	// Restrict result set to items, that matches the given user (only valid for
	// User-targeting Calls, that shall not match the currently loggedin User)
	forUserID int `qs:"forUserID,omitempty"`
	// If the api call wont find enough items, fill the result set with the
	// given Method to the given Limit
	AutoFillResults enums.AutoFill `qs:"autoFillResults,omitempty"`
	// If the output modifier addconnectedmedia is used, this parameter defines
	// the Detail Level for each connected Item.
	ConnectedMediaDetails enums.OutputModifier `qs:"connectedMediaDetails,omitempty"`
	// If the output modifier addcparentmedia is used, this parameter defines
	// the Detail Level for each parent Item.
	ParentMediaDetails enums.OutputModifier `qs:"parentMediaDetails,omitempty"`
	// If the output modifier addchildmedia is used, this parameter defines the
	// Detail Level for each Child Item
	ChildMediaDetails enums.OutputModifier `qs:"childMediaDetails,omitempty"`
	// If the output modifier addreferencingmedia is used, this parameter
	// defines the Detail Level for each referencing Item.
	ReferencingMediaDetails enums.OutputModifier `qs:"referencingMediaDetails,omitempty"`
}

func (g GeneralParameters) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&g)
}
