package omnia

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type QueryParameters struct {
	NoCache                 *Bool            `url:"noc,omitempty"`
	CustomReference         *string          `url:",omitempty"`
	ImageFormat             *ImageFormat     `url:"imageFormat,omitempty"`
	RichTextFormat          *RichTextFormat  `url:"richTextFormat,omitempty"`
	DateFormat              *string          `url:"dateFormat,omitempty"`
	DateFormatTimezone      *string          `url:"dateFormatTimezone,omitempty"`
	DistanceUnit            *DistanceUnit    `url:"distanceUnit,omitempty"`
	TemperatureUnit         *TemperatureUnit `url:"temperatureUnit,omitempty"`
	ExtendCoverGeometry     *Bool            `url:"extendCoverGeometry,omitempty"`
	AdditionalFields        *[]string        `url:"additionalFields,omitempty"`
	ForceGateway            *Gateway         `url:"forceGateway,omitempty"`
	ForceLanguage           *string          `url:"forceLanguage,omitempty"`
	RespectGeoRestrictions  *string          `url:"respectGeoRestrictions,omitempty"`
	RestrictToCurrentDomain *Bool            `url:"restrictToCurrentDomain,omitempty"`
	RestrictToChildDomain   *int             `url:"restrictToChildDomain,omitempty"`
	OrderBy                 *string          `url:"orderBy,omitempty"`
	OrderDirection          *OrderDirection  `url:"oderDir,omitempty"`
	Start                   *int             `url:"start,omitempty"`
	Limit                   *int             `url:"limit,omitempty"`
	ChildLimit              *int             `url:"childLimit,omitempty"`
	AddPublishingDetails    Bool             `url:"addPublishingDetails,omitempty"`
	AddStreamDetails        *Bool            `url:"addStreamDetails,omitempty"`
	AddStatistics           *Bool            `url:"addStatistics,omitempty"`
}

type ResponseMetadata struct {
	Status         int     `json:"status"`
	ApiVersion     string  `json:"apiversion,omitempty"`
	Verb           string  `json:"verb"`
	ProcessingTime float64 `json:"processingtime"`
	CalledWith     *string `json:"calledwith,omitempty"`
	CalledFor      *string `json:"calledfor,omitempty"`
	ForDomain      *string `json:"fordomain,omitempty"`
	FromStage      *int    `json:"fromstage,omitempty"`
	Notice         *string `json:"notice"`
	ErrorHint      *string `json:"errorhint,omitempty"`
}

func (m ResponseMetadata) toMap() map[string]interface{} {
	return structToMap(m)
}

type ResponsePaging struct {
	Start       int `json:"start"`
	Limit       int `json:"limit"`
	ResultCount int `json:"resultcount"`
}

func (p ResponsePaging) toMap() map[string]interface{} {
	return structToMap(p)
}

type Response struct {
	Metadata ResponseMetadata `json:"metadata"`
	Result   interface{}      `json:"result"`
	Paging   *ResponsePaging  `json:"paging"`
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
