package notification

import (
	"encoding/json"
	"fmt"
)

// A notification based on an event within omnia.
type Notification struct {
	Trigger Trigger `json:"trigger"`
	Item    Item    `json:"item"`
	Data    Data    `json:"data"`
}

// Returns a Notification instance based on JSON data
func NotificationFromJson(raw []byte) (*Notification, error) {
	var rsl *Notification
	fmt.Println("HERE")
	if err := json.Unmarshal(raw, rsl); err != nil {
		return nil, err
	}
	return rsl, nil
}

// Information about the trigger of the notification.
type Trigger struct {
	// The »reason« for this trigger.
	Event string `json:"event"`
	// The user ID, that changed the media item (or 0, if created by nexxOMNIA).
	User int `json:"user"`
	// The Session ID, that changed the Media Item (or 0, if created by nexxOMNIA).
	Session int `json:"session"`
	// The Timestamp of the Change.
	Created UnixTS `json:"created"`
	// The Timestamp of the Trigger Processing.
	Sent UnixTS `json:"sent"`
	// The computed Secret for Comparison (if enabled).
	Secret string `json:"secret,omitempty"`
}

// Information about the item which has triggered by the event.
type Item struct {
	// The ID of the Media Item.
	ID string `json:"ID"`
	// The GlobalID of the Media Item.
	GID int `json:"GID"`
	// The external Reference of the Media Item.
	RefNr string `json:"refnr"`
	// The Domain of the Media Item.
	Domain int `json:"domain"`
	// The stream-type of the Media Item.
	StreamType string `json:"streamtype"`
}

// Data about the item of the notification. The data should be more or less analog
// to the editable fields type and should be implemented by the `omnia` package
// instead. This only implements a subset of all available fields.
type Data struct {
	General         GeneralData     `json:"general"`
	ChannelData     ChannelData     `json:"channeldata"`
	ImageData       ImageData       `json:"imagedata"`
	InteractionData InteractionData `json:"interactiondata"`
	PublishingData  PublishingData  `json:"publishingdata"`
}

// Part of the `Data` struct. Based on real world data and not on any documentation.
type GeneralData struct {
	ID          int    `json:"ID"`
	GID         int    `json:"GID"`
	Hash        string `json:"hash"`
	Title       string `json:"title"`
	SubTitle    string `json:"subtitle"`
	GenreRaw    string `json:"genre_raw"`
	Uploaded    UnixTS `json:"uploaded"`
	Created     UnixTS `json:"created"`
	Description string `json:"description"`
	RefNr       string `json:"refnr"`
}

// Part of the `Data` struct. Based on real world data and not on any documentation.
type ChannelData struct{}

// Part of the `Data` struct. Based on real world data and not on any documentation.
type ImageData struct{}

// Part of the `Data` struct. Based on real world data and not on any documentation.
type InteractionData struct{}

// Part of the `Data` struct. Based on real world data and not on any documentation.
type PublishingData struct {
	Origin string `json:"uploadLink"`
}
