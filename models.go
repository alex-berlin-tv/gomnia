package gomnia

import (
	"encoding/json"

	"github.com/alex-berlin-tv/gomnia/enum"
	"github.com/alex-berlin-tv/gomnia/types"
	log "github.com/sirupsen/logrus"
)

// Metadata part of an API response.
type ResponseMetadata struct {
	// The HTTP Status for this Call.
	Status int `json:"status"` // OK
	// Version of the API.
	ApiVersion string `json:"apiversion,omitempty"` // OK
	// The used HTTP Verb.
	Verb string `json:"verb"` // OK
	// Internal Duration, needed to create the response.
	ProcessingTime float64 `json:"processingtime"` // OK
	// The called Endpoint and Parameter.
	CalledWith *string `json:"calledwith,omitempty"`
	// The `cfo` Parameter from the API Call.
	CalledFor *string `json:"calledfor,omitempty"`
	// The calling Domain ID.
	ForDomain *int `json:"fordomain,omitempty"`
	// The result was created by a Stage or Productive Server.
	FromStage *int `json:"fromstage,omitempty"`
	// If the Call uses deprecated Functionality, find here a Hint, what Attributes
	// should be changed.
	Notice *string `json:"notice"`
	// If the Call failed, a Hint for the Failure Reason.
	ErrorHint *string `json:"errorhint,omitempty"`
	// States whether result came from cache
	FromCache *int `json:"fromcache,omitempty"`
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

// Response represents the response structure obtained from a nexxOmnia
// API call. It encapsulates the metadata, result, and paging information
// as documented [here].
//
// [here]: https://api.docs.nexx.cloud/api-design/response-object
type Response[T any] struct {
	Metadata ResponseMetadata `json:"metadata"`
	Result   T
	Paging   *ResponsePaging `json:"paging"`
}

// MediaResult is a collection of MediaResultItem instances, each representing
// a media result. The decision for using a dedicated type for a slice of
// [MediaResultItem]s was made to ease the further work with results outside
// this package.
type MediaResult []MediaResultItem

// MediaResultItem holds detailed information about a single media result item.
type MediaResultItem struct {
	General        MediaResultGeneral        `json:"general"`
	ImageData      MediaResultImageData      `json:"imagedata"`
	ConnectedMedia MediaResultConnectedMedia `json:"connectedmedia"`
}

// MediaResultGeneral provides general information about a media item, including
// its ID, title, and more.
//
// Some fields (like Channel) are optional (or as omnia calls them »additional«)
// fields. You have to use the »additionalFields« parameter for the request. For
// example:
//
//	client := omnia.NewClient("23", "Secret", "42")
//	rsl, err := client.All(enum.AudioStreamType, id, params.Custom{
//		"additionalFields": "channel",
//	})
//
// In order to get all additional fields use the keyword »all«.
type MediaResultGeneral struct {
	Id                       int          `json:"ID"`
	Gid                      int          `json:"GID"`
	Hash                     string       `json:"hash"`
	Title                    string       `json:"title"`
	Subtitle                 string       `json:"subtitle"`
	GenreRaw                 string       `json:"genre_raw"`
	Genre                    string       `json:"genre"`
	ContentModerationAspects string       `json:"contentModerationAspects"`
	Uploaded                 types.UnixTS `json:"uploaded"`
	Created                  types.UnixTS `json:"created"`
	AudioType                string       `json:"audiotype"`
	Runtime                  string       `json:"runtime"`
	IsPicked                 enum.Bool    `json:"isPicked"`
	ForKids                  enum.Bool    `json:"forKids"`
	Channel                  int          `json:"channel,omitempty"` // Optional field.
	IsPay                    enum.Bool    `json:"isPay"`
	IsUgc                    enum.Bool    `json:"isUGC"`
}

// MediaResultImageData contains image-related data for a media item, including
// thumbnails and descriptions.
type MediaResultImageData struct {
	Language          string `json:"language"`
	Thumb             string `json:"thumb"`
	ThumbHasXS        int    `json:"thumb_hasXS"`
	ThumbHasXL        int    `json:"thumb_hasXL"`
	ThumbHasX2        int    `json:"thumb_hasX2"`
	ThumbHasX3        int    `json:"thumb_hasX3"`
	CoversShowTitle   int    `json:"coversShowTitle"`
	Description       string `json:"description"`
	ThumbAction       string `json:"thumb_action"`
	DescriptionAction string `json:"description_action"`
	ThumbBanner       string `json:"thumb_banner"`
	ThumbQuad         string `json:"thumb_quad"`
	ThumbAbt          string `json:"thumb_abt"`
	DescriptionAbt    string `json:"description_abt"`
	Waveform          string `json:"waveform"`
}

// MediaResultConnectedMedia represents connected media items associated with
// a media item.
type MediaResultConnectedMedia struct {
	Shows []MediaResultGeneral `json:"shows"`
}

// EditableAttributesResponse is a map that associates attribute names with their
// editable properties.
type EditableAttributesResponse map[string]EditableAttributesProperties

// EditableAttributesProperties contains information about the editable properties
// of an attribute.
type EditableAttributesProperties struct {
	Type         string `json:"type"`
	MaxLength    int    `json:"maxlength"`
	Format       string `json:"format,omitempty"`
	Hint         string `json:"hint,omitempty"`
	AllowedInUgc int    `json:"allowedInUGC"`
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
