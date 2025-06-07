package config

import "github.com/spf13/viper"

var _ defaulter = (*DatabaseConfig)(nil)

type DatabaseConfig struct {
	URL         string `json:"url" mapstructure:"url"`
	MaxIdleConn int    `json:"maxIdleConn,omitempty" mapstructure:"max_idle_conn"`
}

func (c *DatabaseConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("db", map[string]any{
		"max_idle_time": 2,
	})
}
