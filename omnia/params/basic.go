package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/enum"
	"github.com/pasztorpisti/qs"
)

// General parameters for a query to the nexxOmnia API. The corresponding API
// documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/api-design/query-parameters
type Basic struct {
	// If set to [YesBool], the API will disable Cached Results (which will
	// take longer, so only use this Parameter, if absolutely necessary).
	NoCache enum.Bool `qs:"noc,omitempty"`
	// A custom Reference. It will be returned in the API Response Object for
	// further Processing by the calling Domain
	CustomReference string `qs:",omitempty"`
	// Will return Image Assets as WebP or AVIF, if possible or Classic (jpg/png/gif).
	ImageFormat enum.ImageFormat `qs:"imageFormat,omitempty"`
	// Preprocesses Rich-Text Parts of all Result Elemtns – can be combined with.
	// The API will also accept a Combination of Values, combined by "," (currently
	// not supported by this library). Example: `d.m.Y`.
	RichTextFormat enum.RichTextFormat `qs:"richTextFormat,omitempty"`
	// A valid Date Format to pre-format Date Values (which come as Unix
	// Timestamps by default)
	DateFormat string `qs:"dateFormat,omitempty"`
	// If a dateFormat is given, the default Timezone is used – if a different
	// Timezone is desired, use this Parameter. Example: `Europe/Berlin`.
	DateFormatTimezone string `qs:"dateFormatTimezone,omitempty"`
	// Distances (for example in Geo Searches) will be returned in this Unit.
	DistanceUnit enum.DistanceUnit `qs:"distanceUnit,omitempty"`
	// Temperatures (for example in Weather Requests) will be returned in this Unit.
	TemperatureUnit enum.TemperatureUnit `qs:"temperatureUnit,omitempty"`
	// If set to [YesBool], the API will include AspectRatio and Low-Res Cover
	// DataURIs for each returned Media Item Cover (which results in far more
	// transferred Data).
	ExtendCoverGeometry enum.Bool `qs:"extendCoverGeometry,omitempty"`
	// For Item or Item List Calls, add those Item Attributes to each Item Result
	// Object. List of attributes or `all`.
	AdditionalFields []string `qs:"additionalFields,omitempty"`
	// In a Frontend Call, the Item Set will be automatically reduced to those
	// Items, that are  available for the current Frontend Gateway. If this is
	// not desired, this can be overwritten with this Parameter.
	ForceGateway enum.Gateway `qs:"forceGateway,omitempty"`
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
	RestrictToCurrentDomain enum.Bool `qs:"restrictToCurrentDomain,omitempty"`
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
	OrderDirection enum.OrderDirection `qs:"oderDir,omitempty"`
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
	AddPublishingDetails enum.Bool `qs:"addPublishingDetails,omitempty"`
	// Add technical Details about Origin, Delivery and CDN Locations to each Item.
	AddStreamDetails enum.Bool `qs:"addStreamDetails,omitempty"`
	// Add statistical Data to each Result Set Item.
	AddStatistics enum.Bool `qs:"addStatistics,omitempty"`
}

func (b Basic) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&b)
}
