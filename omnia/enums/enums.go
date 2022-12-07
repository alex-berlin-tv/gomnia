package enums

// A boolean value is expressed as a 0 for `false` and 1 for `true`.
// String is used as type as it's not possible to nil integer values
// (which is needed in order to omit unset parameters as the query parameter).
type Bool string

const (
	NoBool  = Bool("0")
	YesBool = Bool("1")
)

// All instances of the Bool enum.
func (b Bool) Instances() []Bool {
	return []Bool{
		NoBool,
		YesBool,
	}
}

// Used to state the desired image format of a requested image.
type ImageFormat string

const (
	WebpImageFormat = ImageFormat("webp")
	AvifImageFormat = ImageFormat("avif")
	// Will return a jpg, png or gif.
	ClassicImageFormat = ImageFormat("classic")
)

// All instances of the ImageFormat enum.
func (i ImageFormat) Instances() []ImageFormat {
	return []ImageFormat{
		WebpImageFormat,
		AvifImageFormat,
		ClassicImageFormat,
	}
}

// Possible rich text formats.
type RichTextFormat string

const (
	PlainFormat      = RichTextFormat("plain")
	CoverLinksFormat = RichTextFormat("converlinks")
	HtmlFormat       = RichTextFormat("html")
	XmlStrictFormat  = RichTextFormat("xmlstrict")
)

// All instances of the RichTextFormat enum.
func (i RichTextFormat) Instances() []RichTextFormat {
	return []RichTextFormat{
		PlainFormat,
		CoverLinksFormat,
		HtmlFormat,
		XmlStrictFormat,
	}
}

// Metric or imperial distance units.
type DistanceUnit string

const (
	MetricUnit   = DistanceUnit("metric")
	ImperialUnit = DistanceUnit("imperial")
)

// All instances of the DistanceUnit enum.
func (i DistanceUnit) Instances() []DistanceUnit {
	return []DistanceUnit{
		MetricUnit,
		ImperialUnit,
	}
}

// Metric or imperial temperature units.
type TemperatureUnit string

const (
	CelsiusUnit    = TemperatureUnit("celsius")
	FahrenheitUnit = TemperatureUnit("fahrenheit")
)

// All instances of the TemperatureUnit enum.
func (i TemperatureUnit) Instances() []TemperatureUnit {
	return []TemperatureUnit{
		CelsiusUnit,
		FahrenheitUnit,
	}
}

// Used in conjunction with [QueryParameters.ForceGateway].
type Gateway string

const (
	AllGateway     = Gateway("all")
	DesktopGateway = Gateway("desktop")
	MobileGateway  = Gateway("mobile")
	SmartTvGateway = Gateway("smarttv")
	CarGateway     = Gateway("car")
)

// All instances of the Gateway enum.
func (i Gateway) Instances() []Gateway {
	return []Gateway{
		AllGateway,
		DesktopGateway,
		MobileGateway,
		SmartTvGateway,
		CarGateway,
	}
}

// Direction of ordering elements.
type OrderDirection string

const (
	AscendingOrder  = OrderDirection("ASC")
	DescendingOrder = OrderDirection("DESC")
)

// All instances of the OrderDirection enum.
func (i OrderDirection) Instances() []OrderDirection {
	return []OrderDirection{
		AscendingOrder,
		DescendingOrder,
	}
}

// Different streamtypes used in the API call.
type StreamType string

const (
	VideoStreamType = StreamType("videos")
	AudioStreamType = StreamType("audio")
	ShowStreamType  = StreamType("shows")
)

// All instances of the StreamType enum.
func (i StreamType) Instances() []StreamType {
	return []StreamType{
		VideoStreamType,
		AudioStreamType,
		ShowStreamType,
	}
}

// Content type of media items.
type ContentType string

const (
	VideoContentType   = ContentType("video")
	ComicContentType   = ContentType("comic")
	CgiContentType     = ContentType("cgi")
	FotoContentType    = ContentType("foto")
	DrawingContentType = ContentType("drawing")
	ClipartContentType = ContentType("clipart")
)

// All instances of the ContentType enum.
func (i ContentType) Instances() []ContentType {
	return []ContentType{
		VideoContentType,
		ComicContentType,
		CgiContentType,
		FotoContentType,
		DrawingContentType,
		ClipartContentType,
	}
}

// Age categories for age restrictions.
type AgeRestriction string

const (
	AgeRestriction0  = AgeRestriction("0")
	AgeRestriction6  = AgeRestriction("6")
	AgeRestriction12 = AgeRestriction("12")
	AgeRestriction16 = AgeRestriction("16")
	AgeRestriction18 = AgeRestriction("18")
)

// All instances of the AgeRestriction enum.
func (i AgeRestriction) Instances() []AgeRestriction {
	return []AgeRestriction{
		AgeRestriction0,
		AgeRestriction6,
		AgeRestriction12,
		AgeRestriction16,
		AgeRestriction18,
	}
}

// Geometric dimension of a media file.
type Dimension string

const (
	HdDimension     = Dimension("hd")
	FullHdDimension = Dimension("fullhd")
	I2kDimension    = Dimension("2K")
	I4kDimension    = Dimension("4K")
)

// All instances of the Dimension enum.
func (i Dimension) Instances() []Dimension {
	return []Dimension{
		HdDimension,
		FullHdDimension,
		I2kDimension,
		I4kDimension,
	}
}

// Media orientation.
type Orientation string

const (
	PortraitOrientation  = Orientation("portrait")
	LandscapeOrientation = Orientation("landscape")
)

// All instances of the Orientation enum.
func (i Orientation) Instances() []Orientation {
	return []Orientation{
		PortraitOrientation,
		LandscapeOrientation,
	}
}

// Output modifier used to define the detail level.
type OutputModifier string

const (
	FullOutputModifier    = OutputModifier("full")
	DefaultOutputModifier = OutputModifier("default")
	IdOutputModifier      = OutputModifier("ID")
	GidOutputModifier     = OutputModifier("GID")
)

// All instances of the OutputModifier enum.
func (i OutputModifier) Instances() []OutputModifier {
	return []OutputModifier{
		FullOutputModifier,
		DefaultOutputModifier,
		IdOutputModifier,
		GidOutputModifier,
	}
}

// Method for the auto fill method of the API.
type AutoFill string

const (
	RandomAutoFill    = AutoFill("random")
	LatestAutoFill    = AutoFill("latest")
	TopItemsAutoFill  = AutoFill("topitems")
	TopItemsExternal  = AutoFill("topitemsexternal")
	ForkIdsAutoFill   = AutoFill("forkids")
	EvergreenAutoFill = AutoFill("evergreen")
)

// All instances of the AutoFill enum.
func (i AutoFill) Instances() []AutoFill {
	return []AutoFill{
		RandomAutoFill,
		LatestAutoFill,
		TopItemsAutoFill,
		TopItemsExternal,
		ForkIdsAutoFill,
		EvergreenAutoFill,
	}
}

// Query modes.
type QueryMode string

const (
	ClassicWithAndQueryMode = QueryMode("classicwithand")
	ClassicWithOrQueryMode  = QueryMode("classicwithor")
	FulltextQueryMode       = QueryMode("fulltext")
)

// All instances of the QueryMode enum.
func (i QueryMode) Instances() []QueryMode {
	return []QueryMode{
		ClassicWithAndQueryMode,
		ClassicWithOrQueryMode,
		FulltextQueryMode,
	}
}

// Action after rejection of an item.
type ActionAfterRejection string

const (
	DeleteAfterRejection     = ActionAfterRejection("delete")
	ArchiveAfterRejection    = ActionAfterRejection("archive")
	BlockAfterRejection      = ActionAfterRejection("block")
	NewVersionAfterRejection = ActionAfterRejection("newversion")
)

// All instances of the ActionAfterRejection enum.
func (i ActionAfterRejection) Instances() []ActionAfterRejection {
	return []ActionAfterRejection{
		DeleteAfterRejection,
		ArchiveAfterRejection,
		BlockAfterRejection,
		NewVersionAfterRejection,
	}
}
