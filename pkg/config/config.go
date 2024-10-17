package config

import "time"

type RestConfig struct {
	HTTPClient    ClientConfig            `yaml:"http_client"`
	ApiDomain     string                  `yaml:"api_domain"`
	ExternalCalls map[string]ExternalCall `yaml:"external_calls"`
}

type ClientConfig struct {
	Timeout time.Duration `yaml:"timeout"`
}

type ExternalCall struct {
	RequestUri string `yaml:"request_uri"`
}

type SQLiteConfig struct {
	Name string `yaml:"name"`
}
