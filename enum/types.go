package enum

// A boolean value is expressed as a 0 for `false` and 1 for `true`.
// String is used as type as it's not possible to nil integer values
// (which is needed in order to omit unset parameters as the query parameter).
type Bool string

const (
	NoBool  = Bool("0")
	YesBool = Bool("1")
)

// All instances of the Bool
func (b Bool) Instances() []Bool {
	return []Bool{
		NoBool,
		YesBool,
	}
}

func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[Bool](YesBool, data)
	*(*Bool)(b) = *value
	return err
}

// Used to state the desired image format of a requested image.
type ImageFormat string

const (
	WebpImageFormat = ImageFormat("webp")
	AvifImageFormat = ImageFormat("avif")
	// Will return a jpg, png or gif.
	ClassicImageFormat = ImageFormat("classic")
)

// All instances of the ImageFormat
func (i ImageFormat) Instances() []ImageFormat {
	return []ImageFormat{
		WebpImageFormat,
		AvifImageFormat,
		ClassicImageFormat,
	}
}

func (i *ImageFormat) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[ImageFormat](WebpImageFormat, data)
	*(*ImageFormat)(i) = *value
	return err
}

// Possible rich text formats.
type RichTextFormat string

const (
	PlainFormat      = RichTextFormat("plain")
	CoverLinksFormat = RichTextFormat("converlinks")
	HtmlFormat       = RichTextFormat("html")
	XmlStrictFormat  = RichTextFormat("xmlstrict")
)

// All instances of the RichTextFormat
func (i RichTextFormat) Instances() []RichTextFormat {
	return []RichTextFormat{
		PlainFormat,
		CoverLinksFormat,
		HtmlFormat,
		XmlStrictFormat,
	}
}

func (i *RichTextFormat) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[RichTextFormat](PlainFormat, data)
	*(*RichTextFormat)(i) = *value
	return err
}

// Metric or imperial distance units.
type DistanceUnit string

const (
	MetricUnit   = DistanceUnit("metric")
	ImperialUnit = DistanceUnit("imperial")
)

// All instances of the DistanceUnit
func (i DistanceUnit) Instances() []DistanceUnit {
	return []DistanceUnit{
		MetricUnit,
		ImperialUnit,
	}
}

func (i *DistanceUnit) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[DistanceUnit](MetricUnit, data)
	*(*DistanceUnit)(i) = *value
	return err
}

// Metric or imperial temperature units.
type TemperatureUnit string

const (
	CelsiusUnit    = TemperatureUnit("celsius")
	FahrenheitUnit = TemperatureUnit("fahrenheit")
)

// All instances of the TemperatureUnit
func (i TemperatureUnit) Instances() []TemperatureUnit {
	return []TemperatureUnit{
		CelsiusUnit,
		FahrenheitUnit,
	}
}

func (i *TemperatureUnit) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[TemperatureUnit](CelsiusUnit, data)
	*(*TemperatureUnit)(i) = *value
	return err
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

// All instances of the Gateway
func (i Gateway) Instances() []Gateway {
	return []Gateway{
		AllGateway,
		DesktopGateway,
		MobileGateway,
		SmartTvGateway,
		CarGateway,
	}
}

func (i *Gateway) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[Gateway](AllGateway, data)
	*(*Gateway)(i) = *value
	return err
}

// Direction of ordering elements.
type OrderDirection string

const (
	AscendingOrder  = OrderDirection("ASC")
	DescendingOrder = OrderDirection("DESC")
)

// All instances of the OrderDirection
func (i OrderDirection) Instances() []OrderDirection {
	return []OrderDirection{
		AscendingOrder,
		DescendingOrder,
	}
}

func (i *OrderDirection) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[OrderDirection](AscendingOrder, data)
	*(*OrderDirection)(i) = *value
	return err
}

// Different streamtypes used in the API call.
type StreamType string

const (
	VideoStreamType = StreamType("videos")
	AudioStreamType = StreamType("audio")
	ShowStreamType  = StreamType("shows")
)

// All instances of the StreamType
func (i StreamType) Instances() []StreamType {
	return []StreamType{
		VideoStreamType,
		AudioStreamType,
		ShowStreamType,
	}
}

func (i *StreamType) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[StreamType](VideoStreamType, data)
	*(*StreamType)(i) = *value
	return err
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

// All instances of the ContentType
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

func (i *ContentType) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[ContentType](VideoContentType, data)
	*(*ContentType)(i) = *value
	return err
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

// All instances of the AgeRestriction
func (i AgeRestriction) Instances() []AgeRestriction {
	return []AgeRestriction{
		AgeRestriction0,
		AgeRestriction6,
		AgeRestriction12,
		AgeRestriction16,
		AgeRestriction18,
	}
}

func (i *AgeRestriction) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[AgeRestriction](AgeRestriction0, data)
	*(*AgeRestriction)(i) = *value
	return err
}

// Geometric dimension of a media file.
type Dimension string

const (
	HdDimension     = Dimension("hd")
	FullHdDimension = Dimension("fullhd")
	I2kDimension    = Dimension("2K")
	I4kDimension    = Dimension("4K")
)

// All instances of the Dimension
func (i Dimension) Instances() []Dimension {
	return []Dimension{
		HdDimension,
		FullHdDimension,
		I2kDimension,
		I4kDimension,
	}
}

func (i *Dimension) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[Dimension](HdDimension, data)
	*(*Dimension)(i) = *value
	return err
}

// Media orientation.
type Orientation string

const (
	PortraitOrientation  = Orientation("portrait")
	LandscapeOrientation = Orientation("landscape")
)

// All instances of the Orientation
func (i Orientation) Instances() []Orientation {
	return []Orientation{
		PortraitOrientation,
		LandscapeOrientation,
	}
}

func (i *Orientation) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[Orientation](PortraitOrientation, data)
	*(*Orientation)(i) = *value
	return err
}

// Output modifier used to define the detail level.
type OutputModifier string

const (
	FullOutputModifier    = OutputModifier("full")
	DefaultOutputModifier = OutputModifier("default")
	IdOutputModifier      = OutputModifier("ID")
	GidOutputModifier     = OutputModifier("GID")
)

// All instances of the OutputModifier
func (i OutputModifier) Instances() []OutputModifier {
	return []OutputModifier{
		FullOutputModifier,
		DefaultOutputModifier,
		IdOutputModifier,
		GidOutputModifier,
	}
}

func (i *OutputModifier) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[OutputModifier](FullOutputModifier, data)
	*(*OutputModifier)(i) = *value
	return err
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

// All instances of the AutoFill
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

func (i *AutoFill) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[AutoFill](RandomAutoFill, data)
	*(*AutoFill)(i) = *value
	return err
}

// Query modes.
type QueryMode string

const (
	ClassicWithAndQueryMode = QueryMode("classicwithand")
	ClassicWithOrQueryMode  = QueryMode("classicwithor")
	FulltextQueryMode       = QueryMode("fulltext")
)

// All instances of the QueryMode
func (i QueryMode) Instances() []QueryMode {
	return []QueryMode{
		ClassicWithAndQueryMode,
		ClassicWithOrQueryMode,
		FulltextQueryMode,
	}
}

func (i *QueryMode) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[QueryMode](ClassicWithAndQueryMode, data)
	*(*QueryMode)(i) = *value
	return err
}

// Action after rejection of an item.
type ActionAfterRejection string

const (
	DeleteAfterRejection     = ActionAfterRejection("delete")
	ArchiveAfterRejection    = ActionAfterRejection("archive")
	BlockAfterRejection      = ActionAfterRejection("block")
	NewVersionAfterRejection = ActionAfterRejection("newversion")
)

// All instances of the ActionAfterRejection
func (i ActionAfterRejection) Instances() []ActionAfterRejection {
	return []ActionAfterRejection{
		DeleteAfterRejection,
		ArchiveAfterRejection,
		BlockAfterRejection,
		NewVersionAfterRejection,
	}
}

func (i *ActionAfterRejection) UnmarshalJSON(data []byte) (err error) {
	value, err := EnumByByteValue[ActionAfterRejection](DeleteAfterRejection, data)
	*(*ActionAfterRejection)(i) = *value
	return err
}
