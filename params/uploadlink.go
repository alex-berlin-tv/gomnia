package params

import (
	"fmt"

	"github.com/alex-berlin-tv/gomnia/enum"
	"github.com/pasztorpisti/qs"
)

// Parameters for adding a UploadLink to the domain. The corresponding
// documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/domain-management#uploadlinks
type UploadLink struct {
	// A Title for the new UploadLink.
	Title string `qs:"title"`
	// A comma-separated List of Streamtypes, that can be uploaded
	// via this Link. Possible is [video, audio, image, file].
	SelectedStreamtypes string `qs:"selectedStreamtypes"`
	// A 2-Letter coded Language for the Frontend. Currently supported
	// is [de, en, es, fr].
	Language string `qs:"language"`
	// If the UploadLink shall be restricted in Usage, use this Parameter
	// for the maximal Number of Usages.
	MaxUsages int `qs:"maxUsages,omitempty"`
	// If desired, an optional code for further Protection of the Link.
	Password string `qs:"code,omitempty"`
	// If set to 1, the UploadLink will force the User to add some "Notes"
	// as additional Info for this Upload.
	AskForNotes enum.Bool `qs:"asForNotes,omitempty"`
	// If set to 1, the UploadLink UI will use the target Domain Colors
	// and Icons.
	UseDomainStyle enum.Bool `qs:"useDomainStyle"`
}

func (u UploadLink) UrlEncode() (string, error) {
	return qs.Marshal(&u)
}

// Checks if the instance is valid for the API. Returns an error with an
// explanation.
func (u UploadLink) Validate() error {
	if u.Title == "" {
		return fmt.Errorf("title has to be set")
	}
	if u.SelectedStreamtypes == "" {
		return fmt.Errorf("selectedStreamtypes has to be set")
	}
	if u.Language == "" {
		return fmt.Errorf("language has to be set")
	}
	found := false
	for _, lang := range []string{"de", "en", "es", "fr"} {
		if u.Language == lang {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("language has to be de, en, es, or fr")
	}
	return nil
}
