package config

import "github.com/spf13/viper"

var _ defaulter = (*CorsConfig)(nil)

type CorsConfig struct {
	Enabled        bool     `json:"enabled" mapstructure:"enabled"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty" mapstructure:"allowed_origins"`
}

func (c *CorsConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("cors", map[string]any{
		"enabled":         false,
		"allowed_origins": "*",
	})
}
