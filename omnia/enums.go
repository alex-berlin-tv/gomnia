package omnia

type Bool int

const (
	NoBool  Bool = 0
	YesBool      = 1
)

type ImageFormat string

const (
	Webp    ImageFormat = "webp"
	Avif                = "avif"
	Classic             = "classic"
)

type RichTextFormat string

const (
	Plain      RichTextFormat = "plain"
	CoverLinks                = "converlinks"
	Html                      = "html"
	XmlStrict                 = "xmlstrict"
)

type DistanceUnit string

const (
	Metric   DistanceUnit = "metric"
	IMPERIAL              = "imperial"
)

type TemperatureUnit string

const (
	Celisus    TemperatureUnit = "celsius"
	Fahrenheit                 = "fahrenheit"
)

type Gateway string

const (
	All     Gateway = "all"
	Desktop         = "desktop"
	Mobile          = "mobile"
	SmartRv         = "smarttv"
	Car             = "car"
)

type OrderDirection string

const (
	Ascending  OrderDirection = "ASC"
	Descending                = "DESC"
)

type StreamType string

const (
	Video StreamType = "videos"
	Audio            = "audio"
)
