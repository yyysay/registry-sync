package config

type SourceConfig struct {
	Sync     []string                `yaml:"sync,omitempty"`
	Platform []string                `yaml:"platform,omitempty"`
	Auth     *AuthConfig             `yaml:"auth,omitempty"`
	Images   map[string]*ImageConfig `yaml:"images"`
}
