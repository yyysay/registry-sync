package config

type DestinationConfig struct {
	Registry  string      `yaml:"registry"`
	Namespace string      `yaml:"namespace,omitempty"`
	Flatten   bool        `yaml:"flatten,omitempty"`
	Auth      *AuthConfig `yaml:"auth,omitempty"`
}
