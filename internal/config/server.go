package config

import "github.com/spf13/viper"

var _ defaulter = (*ServerConfig)(nil)

type ServerConfig struct {
	Port int `json:"port,omitempty" mapstructure:"port"`
}

func (s *ServerConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("server", map[string]any{
		"port": 8080,
	})
}
