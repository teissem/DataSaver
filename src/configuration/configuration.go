package configuration

type Parser func(source string) (*Configuration, error)

type SrcDest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Git struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Repositories []SrcDest `json:"repositories"`
}

type Folder struct {
	Path []SrcDest `json:"path"`
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
