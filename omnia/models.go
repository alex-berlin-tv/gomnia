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
	ForDomain *int `json:"fordomain,omitempty"`
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

type Response interface {
	GetMetadata() ResponseMetadata
	GetResult() interface{}
	GetPaging() *ResponsePaging
}

// The response of an nexxOmnia API call. As documented [here].
//
// [here]: https://api.docs.nexx.cloud/api-design/response-object
type UniversalResponse struct {
	// Metadata.
	Metadata ResponseMetadata `json:"metadata"`
	// Acutal result of the call. Structure can vary widely.
	Result interface{} `json:"result"`
	// Optional information on the paging.
	Paging *ResponsePaging `json:"paging"`
}

func (u UniversalResponse) GetMetadata() ResponseMetadata {
	return u.Metadata
}

func (u UniversalResponse) GetResult() interface{} {
	return u.Result
}

func (u UniversalResponse) GetPaging() *ResponsePaging {
	return u.Paging
}

type MediaResult struct {
	General   MediaResultGeneral   `json:"general"`
	ImageData MediaResultImageData `json:"imagedata"`
}

type MediaResultGeneral struct {
	Id                       int    `json:"ID"`
	Gid                      int    `json:"GID"`
	Hash                     string `json:"hash"`
	Title                    string `json:"title"`
	Subtitle                 string `json:"subtitle"`
	GenreRaw                 string `json:"genre_raw"`
	Genre                    string `json:"genre"`
	ContentModerationAspects string `json:"contentModerationAspects"`
	Uploaded                 int    `json:"uploaded"`
	Created                  int    `json:"created"`
	AudioType                string `json:"audiotype"`
	Runtime                  string `json:"runtime"`
	IsPicked                 int    `json:"isPicked"`
	ForKids                  int    `json:"forKids"`
	IsPay                    int    `json:"isPay"`
	IsUgc                    int    `json:"isUGC"`
}

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

// Response for Media API queries.
type MediaResponse struct {
	// Metadata.
	Metadata ResponseMetadata `json:"metadata"`
	// Acutal result of the call. Structure can vary widely.
	Result []MediaResult `json:"result"`
	// Optional information on the paging.
	Paging *ResponsePaging `json:"paging"`
}

func (m MediaResponse) GetMetadata() ResponseMetadata {
	return m.Metadata
}

func (m MediaResponse) GetResult() interface{} {
	return m.Result
}

func (m MediaResponse) GetPaging() *ResponsePaging {
	return m.Paging
}

// Response for the system editable attributes call.
type EditableAttributesResponse struct {
	// Metadata.
	Metadata ResponseMetadata `json:"metadata"`
	// Acutal result of the call.
	Result []map[string]EditableAttributesProperties `json:"result"`
}

func (e EditableAttributesResponse) GetMetadata() ResponseMetadata {
	return e.Metadata
}

func (e EditableAttributesResponse) GetResult() interface{} {
	return e.Result
}

func (e EditableAttributesResponse) GetPaging() *ResponsePaging {
	return nil
}

type EditableAttributesProperties struct {
	Type         string `json:"type"`
	MaxLength    int    `json:"maxlength"`
	Format       string `json:"format"`
	Hint         string `json:"hint"`
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
