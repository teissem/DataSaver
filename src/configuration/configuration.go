package configuration

type Parser func(path string) (*Configuration, error)

type GitConfiguration struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Repositories []string `json:"repositories"`
}

type Configuration struct {
	Destination string           `json:"destination"`
	Compression string           `json:"compression"`
	Folder      []string         `json:"folder"`
	Git         GitConfiguration `json:"git"`
}

func SupportedConfigurationFormat() map[string]Parser {
	configurationMap := make(map[string]Parser)
	configurationMap[".json"] = ParseJSON
	return configurationMap
}
