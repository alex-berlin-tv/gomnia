package omnia

type QueryParameters struct {
	NoCache                 *Bool            `json:"noc,omitempty"`
	CustomReference         *string          `json:",omitempty"`
	ImageFormat             *ImageFormat     `json:"imageFormat,omitempty"`
	RichTextFormat          *RichTextFormat  `json:"richTextFormat,omitempty"`
	DateFormat              *string          `json:"dateFormat,omitempty"`
	DateFormatTimezone      *string          `json:"dateFormatTimezone,omitempty"`
	DistanceUnit            *DistanceUnit    `json:"distanceUnit,omitempty"`
	TemperatureUnit         *TemperatureUnit `json:"temperatureUnit,omitempty"`
	ExtendCoverGeometry     *Bool            `json:"extendCoverGeometry,omitempty"`
	AdditionalFields        *[]string        `json:"additionalFields,omitempty"`
	ForceGateway            *Gateway         `json:"forceGateway,omitempty"`
	ForceLanguage           *string          `json:"forceLanguage,omitempty"`
	RespectGeoRestrictions  *string          `json:"respectGeoRestrictions,omitempty"`
	RestrictToCurrentDomain *Bool            `json:"restrictToCurrentDomain,omitempty"`
	RestrictToChildDomain   *int             `json:"restrictToChildDomain,omitempty"`
	OrderBy                 *string          `json:"orderBy,omitempty"`
	OrderDirection          *OrderDirection  `json:"oderDir,omitempty"`
	Start                   *int             `json:"start,omitempty"`
	Limit                   *int             `json:"limit,omitempty"`
	ChildLimit              *int             `json:"childLimit,omitempty"`
	AddPublishingDetails    *Bool            `json:"addPublishingDetails,omitempty"`
	AddStreamDetails        *Bool            `json:"addStreamDetails,omitempty"`
	AddStatistics           *Bool            `json:"addStatistics,omitempty"`
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

type ResponsePaging struct {
	Start       int `json:"start"`
	Limit       int `json:"limit"`
	ResultCount int `json:"resultcount"`
}

type Response struct {
	Metadata ResponseMetadata `json:"metadata"`
	Result   interface{}      `json:"result"`
	Paging   *ResponsePaging  `json:"paging"`
}
