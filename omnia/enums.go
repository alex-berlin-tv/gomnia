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
	VideoStreamType StreamType = "videos"
	AudioStreamType            = "audio"
)

// Content type of media items.
type ContentType string

const (
	VideoContentType   ContentType = "video"
	ComicContentType               = "comic"
	CgiContentType                 = "cgi"
	FotoContentType                = "foto"
	DrawingContentType             = "drawing"
	ClipartContentType             = "clipart"
)

// Age categories for age restrictions.
type AgeRestriction string

const (
	AgeRestriction0  AgeRestriction = "0"
	AgeRestriction6                 = "6"
	AgeRestriction12                = "12"
	AgeRestriction16                = "16"
	AgeRestriction18                = "18"
)

// Geometric dimension of a media file.
type Dimension string

const (
	HdDimension     Dimension = "hd"
	FullHdDimension           = "fullhd"
	I2kDimension              = "2K"
	I4kDimension              = "4K"
)

// Media orientation.
type Orientation string

const (
	PortraitOrientation  Orientation = "portrait"
	LandscapeOrientation             = "landscape"
)

// Ouptut modifier used to define the detail level.
type OutputModifier string

const (
	FullOutputModifier    OutputModifier = "full"
	DefaultOutputModifier                = "default"
	IdOutputModifier                     = "ID"
	GidOutputModifier                    = "GID"
)

// Method for the auto fill method of the API.
type AutoFill string

const (
	RandomAutoFill    AutoFill = "random"
	LatestAutoFill             = "latest"
	TopItemsAutoFill           = "topitems"
	TopItemsExternal           = "topitemsexternal"
	ForkIdsAutoFill            = "forkids"
	EvergreenAutoFill          = "evergreen"
)

// Query modes.
type QueryMode string

const (
	ClassicWithAndQueryMode QueryMode = "classicwithand"
	ClassicWithOrQueryMode            = "classicwithor"
	FulltextQueryMode                 = "fulltext"
)
