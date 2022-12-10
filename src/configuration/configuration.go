package configuration

type Parser func(source string) (*Configuration, error)

type Git struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Repositories []string `json:"repositories"`
}

type Folder struct {
	Path []string `json:"path"`
}

type Configuration struct {
	Destination  string `json:"destination"`
	Compression  string `json:"compression"`
	FolderSource Folder `json:"folder"`
	GitSource    Git    `json:"git"`
}

func SupportedConfigurationFormat() map[string]Parser {
	configurationMap := make(map[string]Parser)
	configurationMap[".json"] = ParseJSON
	return configurationMap
}
