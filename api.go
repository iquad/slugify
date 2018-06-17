package slugify

import "strings"

// API structs creates a wrapper
// to be used for custom configurations
type API struct {
	replaceSet []string
}

// SetReplaceSet changes default replace table
// with given table.
func (api *API) SetReplaceSet(userReplaceSet []string) {
	api.replaceSet = userReplaceSet
}

// Slugify converts given string into unicode compatible
// slug with given configuration
func (api *API) Slugify(str string) (result string, err error) {
	replacer := strings.NewReplacer(api.replaceSet...)
	return transformizer(replacer, str)
}

// GetWithCustomReplacer returns API instance with
// given configuration
func GetWithCustomReplacer(replacer []string) (*API) {
	return &API{
		replaceSet:replacer,
	}
}