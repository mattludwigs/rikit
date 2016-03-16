package config

type ConfigObject struct {
	URL  string `json:"url,omitempty"`
	Auth string `json:"auth,omitempty"`
}

type SitesConfig map[string]ConfigObject

type Config struct {
	Sites SitesConfig `json:"sites, omitempty"`
}
