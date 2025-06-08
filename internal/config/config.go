package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"golang.org/x/exp/constraints"
)

type Config struct {
	Version  string         `json:"version,omitempty" mapstructure:"version"`
	Log      LogConfig      `json:"log" mapstructure:"log"`
	Cors     CorsConfig     `json:"cors" mapstructure:"cors"`
	Database DatabaseConfig `json:"db" mapstructure:"db"`
	Server   ServerConfig   `json:"server" mapstructure:"server"`
}

type defaulter interface {
	setDefaults(v *viper.Viper)
}

var decodeHooks = mapstructure.ComposeDecodeHookFunc(
	mapstructure.StringToTimeDurationHookFunc(),
	stringToSliceHookFunc(),
	stringToEnumHookFunc(stringToLevelEncoding),
	stringToEnumHookFunc(stringToLogEncoding),
)

func (cfg *Config) Load(path string) error {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("loading configuration: %w", err)
	}

	// Calls set interface methods on all fields that implement the defaulter interface
	// This allows us to set default values for fields that implement the defaulter interface
	f := func(field any) {
		if d, ok := field.(defaulter); ok {
			d.setDefaults(v)
		}
	}

	// invoke field visitors on the root config firsts
	root := reflect.ValueOf(cfg).Interface()
	f(root)

	val := reflect.ValueOf(cfg).Elem()
	for i := range val.NumField() {
		field := val.Field(i).Addr().Interface()
		f(field)
	}

	// Unmarshal config into struct
	if err := v.Unmarshal(cfg, viper.DecodeHook(decodeHooks)); err != nil {
		return err
	}

	return nil
}

func stringToEnumHookFunc[T constraints.Integer](mappings map[string]T) mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String || t != reflect.TypeOf(T(0)) {
			return data, nil
		}

		enum := mappings[data.(string)]
		return enum, nil
	}
}

func stringToSliceHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Kind, t reflect.Kind, data interface{}) (interface{}, error) {
		if f != reflect.String || t != reflect.Slice {
			return data, nil
		}

		raw := data.(string)
		if raw == "" {
			return []string{}, nil
		}

		return strings.Fields(raw), nil
	}
}
