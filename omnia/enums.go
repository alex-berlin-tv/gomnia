package omnia

// A boolean value is expressed as a 0 for `false` and 1 for `true`.
// String is used as type as it's not possible to nil integer values
// (which is needed in order to omit unset parameters as the query parameter).
type Bool string

const (
	NoBool  Bool = "0"
	YesBool      = "1"
)

// Used to state the desired image format of a requested image.
type ImageFormat string

const (
	WebpImageFormat ImageFormat = "webp"
	AvifImageFormat             = "avif"
	// Will return a jpg, png or gif.
	ClassicImageFormat = "classic"
)

// Possible rich text formats.
type RichTextFormat string

const (
	PlainFormat      RichTextFormat = "plain"
	CoverLinksFormat                = "converlinks"
	HtmlFormat                      = "html"
	XmlStrictFormat                 = "xmlstrict"
)

// Metric or imperial distance units.
type DistanceUnit string

const (
	MetricUnit   DistanceUnit = "metric"
	ImperialUnit              = "imperial"
)

// Metric or imperial temperature units.
type TemperatureUnit string

const (
	CelisusUnit    TemperatureUnit = "celsius"
	FahrenheitUnit                 = "fahrenheit"
)

// Used in conjection with [QueryParameters.ForceGateway].
type Gateway string

const (
	AllGateway     Gateway = "all"
	DesktopGateway         = "desktop"
	MobileGateway          = "mobile"
	SmartRvGateway         = "smarttv"
	CarGateway             = "car"
)

// Direction of ordering elements.
type OrderDirection string

const (
	AscendingOrder  OrderDirection = "ASC"
	DescendingOrder                = "DESC"
)

// Different streamtypes used in the API call.
type StreamType string

const (
	VideoType StreamType = "videos"
	AudioType            = "audio"
)
