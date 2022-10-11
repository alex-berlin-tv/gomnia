package omnia

import (
	"net/url"

	"github.com/pasztorpisti/qs"
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

// General parameters for a query to the nexxOmnia API. The corresponding API
// documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/api-design/query-parameters
type BasicParameters struct {
	// If set to [YesBool], the API will disable Cached Results (which will
	// take longer, so only use this Parameter, if absolutely necessary).
	NoCache Bool `qs:"noc,omitempty"`
	// A custom Reference. It will be returned in the API Response Object for
	// further Processing by the calling Domain
	CustomReference string `qs:",omitempty"`
	// Will return Image Assets as WebP or AVIF, if possible or Classic (jpg/png/gif).
	ImageFormat ImageFormat `qs:"imageFormat,omitempty"`
	// Preprocesses Rich-Text Parts of all Result Elemtns – can be combined with.
	// The API will also accept a Combination of Values, combined by "," (currently
	// not supported by this library). Example: `d.m.Y`.
	RichTextFormat RichTextFormat `qs:"richTextFormat,omitempty"`
	// A valid Date Format to pre-format Date Values (which come as Unix
	// Timestamps by default)
	DateFormat string `qs:"dateFormat,omitempty"`
	// If a dateFormat is given, the default Timezone is used – if a different
	// Timezone is desired, use this Parameter. Example: `Europe/Berlin`.
	DateFormatTimezone string `qs:"dateFormatTimezone,omitempty"`
	// Distances (for example in Geo Searches) will be returned in this Unit.
	DistanceUnit DistanceUnit `qs:"distanceUnit,omitempty"`
	// Temperatures (for example in Weather Requests) will be returned in this Unit.
	TemperatureUnit TemperatureUnit `qs:"temperatureUnit,omitempty"`
	// If set to [YesBool], the API will include AspectRatio and Low-Res Cover
	// DataURIs for each returned Media Item Cover (which results in far more
	// transferred Data).
	ExtendCoverGeometry Bool `qs:"extendCoverGeometry,omitempty"`
	// For Item or Item List Calls, add those Item Attributes to each Item Result
	// Object. List of attributes or `all`.
	AdditionalFields []string `qs:"additionalFields,omitempty"`
	// In a Frontend Call, the Item Set will be automatically reduced to those
	// Items, that are  available for the current Frontend Gateway. If this is
	// not desired, this can be overwritten with this Parameter.
	ForceGateway Gateway `qs:"forceGateway,omitempty"`
	// In a Frontend Call for a Domain, that supports multiple Language, the
	// Text Attributes of each Item will automatically be returned in the current
	// Session Language. If not desired, this can overwritten with this Parameter
	// (if supported by the Domain and existing in the current Item). 2-Letter-Code
	// of a supported Frontend Language.
	ForceLanguage string `qs:"forceLanguage,omitempty"`
	// In a Frontend List Call, by default, all Elements are returned and
	// Geo-Restrictions are computed on Item-Level. If the List Calls should also
	// respect the Domain/Item Geo Restrictions, set this Parameter. 2-Letter-Code
	// of target Country or `auto`.
	RespectGeoRestrictions string `qs:"respectGeoRestrictions,omitempty"`
	// If the calling Domain belongs to a network, by default, all valid Elements
	// for all Network-Mode controlled Domains in this Network are returned. If only
	// the "real" Elements of the calling Domain are desired, use this Parameter
	// with 1.
	RestrictToCurrentDomain Bool `qs:"restrictToCurrentDomain,omitempty"`
	// If the calling Domain belongs to a network, by default, all valid Elements
	// for all Network-Mode controlled Domains in this Network are returned. If only
	// the Elements of a Child Domain of the calling Domain are desired (and the
	// calling Domain is the Network Mother Domain), use this Parameter with the ID
	// of that Child Domain.
	RestrictToChildDomain int `qs:"restrictToChildDomain,omitempty"`
	// Orders the Resultset by the given Attribute. If omitted, the Items will be
	// ordered by date DESC (notice, that date in this case is notuploaded or
	// created, but apiuploaded (i.e., a virtual Attribute, that can be rewritten
	// via API/nexxOMNIA).
	OrderBy string `qs:"orderBy,omitempty"`
	// The order direction.
	OrderDirection OrderDirection `qs:"oderDir,omitempty"`
	// The Result Set will start at this Item Number.
	Start int `qs:"start,omitempty"`
	// The maximal Size of the Result Set. Has to be between 1 and 100.
	Limit int `qs:"limit,omitempty"`
	// If the API Calls targets a Container Streamtype and forces the Inclusion of.
	// Child Elements, limit the Number of Child Elements to this Value
	ChildLimit int `qs:"childLimit,omitempty"`
	// Add an Object of Publishing States and Restrictions to each Item. When adding
	// this Output Modifier, it is possible (and that’s he only accepted way) to
	// query for inactive/unpublished Objects.
	AddPublishingDetails Bool `qs:"addPublishingDetails,omitempty"`
	// Add technical Details about Origin, Delivery and CDN Locations to each Item.
	AddStreamDetails Bool `qs:"addStreamDetails,omitempty"`
	// Add statistical Data to each Result Set Item.
	AddStatistics Bool `qs:"addStatistics,omitempty"`
}

func (b BasicParameters) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&b)
}

type ByQueryParameters struct {
	BasicParameters
	// Defines the Way, the Query is executed. Fore more results, "classicwithor"
	// is optimal. For a Lucene Search with Relevance, use "fulltext".
	QueryMode QueryMode `qs:"queryMode,omitempty"`
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

func (b ByQueryParameters) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&b)
}

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
	RespectChannelHierarchy Bool `qs:"respectChannelHierarchy,omitempty"`
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
	ContentType ContentType `qs:"contentType,omitempty"`
	// Restricts the result to media from the given country
	Country string `qs:"country,omitempty"`
	// Restrict result set to items with defined age level < 13 and without
	// Content Moderation Aspects
	NoExplicit Bool `qs:"noExplicit,omitempty"`
	// Restrict result set to items without content moderation aspects
	NoContentModerationHints Bool `qs:"noContentModerationHints,omitempty"`
	// Restrict result set to items with maximally the given age level
	MaxAge AgeRestriction `qs:"maxAge,omitempty"`
	// Restrict result set to items with at least the given age level
	MinAge AgeRestriction `qs:"minAge,omitempty"`
	// Restrict the result to media with the given height (alternatively, this
	// Parameter also accepts a numeric Value, which will be mapped to the Media
	// Height)
	Dimension Dimension `qs:"dimension,omitempty"`
	// Restrict the result to media with the given orientation
	Orientation Orientation `qs:"orientation,omitempty"`
	// If set to 1, only media in hdr quality will be returned
	OnlyHdr Bool `qs:"onlyHDR,omitempty"`
	// The given items wont be included into the result set
	ExcludeItems []int `qs:"excludeItems,omitempty"`
	// Add items, uploaded by the community, to the result set
	IncludeUGC Bool `qs:"includeUGC,omitempty"`
	// Restrict result set to items, uploaded by the community
	OnlyUGC Bool `qs:"onlyUGC,omitempty"`
	// Also include files, that originate not by nexxomnia, but a partner
	// Provider
	IncludeRemote Bool `qs:"includeRemote,omitempty"`
	// Only include files, that originate not by nexxomnia, but a partner
	// Provider
	OnlyRemote Bool `qs:"onlyRemote,omitempty"`
	// Also include media items, that are marked as not listable.  this
	// Parameter should be used only in very specific Usecases.
	IncludeNotListables Bool `qs:"includeNotListables,omitempty"`
	// Only valid for container calls with addchildmedia parameter.  this
	// Parameter will add currently invalid Elements to the Child Listing.  To
	// make this Parameter work, an active eternal Session must be used.
	IncludeInvalidChildMedia Bool `qs:"includeInvalidChildMedia,omitempty"`
	// Also include media, that are not valid yet, but will be in the near
	// Future and allow Premiere Functionality
	IncludePremieres Bool `qs:"includePremieres,omitempty"`
	// Restrict result set to items with payment attributes
	OnlyPay Bool `qs:"onlyPay,omitempty"`
	// Restrict result set to items with payment attributes and premium payment
	// Attributes
	OnlyPremiumPay Bool `qs:"onlyPremiumPay,omitempty"`
	// Restrict result set to items with payment attributes and standard payment
	// Attributes
	OnlyStandardPay Bool `qs:"onlyStandardPay,omitempty"`
	// Only possible if addpublishingdetails is active.  if set to 1, only
	// planned Elements will be returned.
	OnlyPlanned Bool `qs:"onlyPlanned,omitempty"`
	// Only possible, if addpublishingdetails is active.  if set to 1, only
	// unpublished Elements will be returned.
	OnlyInactive Bool `qs:"onlyInactive,omitempty"`
	// Restrict result set to items, that matches the given user (only valid for
	// User-targeting Calls, that shall not match the currently loggedin User)
	forUserID int `qs:"forUserID,omitempty"`
	// If the api call wont find enough items, fill the result set with the
	// given Method to the given Limit
	AutoFillResults AutoFill `qs:"autoFillResults,omitempty"`
	// If the output modifier addconnectedmedia is used, this parameter defines
	// the Detail Level for each connected Item.
	ConnectedMediaDetails OutputModifier `qs:"connectedMediaDetails,omitempty"`
	// If the output modifier addcparentmedia is used, this parameter defines
	// the Detail Level for each parent Item.
	ParentMediaDetails OutputModifier `qs:"parentMediaDetails,omitempty"`
	// If the output modifier addchildmedia is used, this parameter defines the
	// Detail Level for each Child Item
	ChildMediaDetails OutputModifier `qs:"childMediaDetails,omitempty"`
	// If the output modifier addreferencingmedia is used, this parameter
	// defines the Detail Level for each referencing Item.
	ReferencingMediaDetails OutputModifier `qs:"referencingMediaDetails,omitempty"`
}

func (g GeneralParameters) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&g)
}
