package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

var _ defaulter = (*LogConfig)(nil)

type LogConfig struct {
	Level    LevelEncoding `json:"level,omitempty" mapstructure:"level"`
	Encoding LogEncoding   `json:"encoding,omitempty" mapstructure:"encoding"`
}

func (l *LogConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("log", map[string]any{
		"level":    "debug",
		"encoding": "json",
	})
}

type LogEncoding uint8

const (
	_ LogEncoding = iota
	LogEncodingConsole
	LogEncodingJSON
)

var (
	logEncodingToString = [...]string{
		LogEncodingConsole: "console",
		LogEncodingJSON:    "json",
	}

	stringToLogEncoding = map[string]LogEncoding{
		"console": LogEncodingConsole,
		"json":    LogEncodingJSON,
	}
)

func (e LogEncoding) String() string {
	return logEncodingToString[e]
}

func (e LogEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

type LevelEncoding uint8

const (
	_ LevelEncoding = iota
	LevelEncodingDebug
	LevelEncodingInfo
	LevelEncodingWarn
	LevelEncodingError
)

var (
	levelEncodingToString = [...]string{
		LevelEncodingDebug: "debug",
		LevelEncodingInfo:  "info",
		LevelEncodingWarn:  "warn",
		LevelEncodingError: "error",
	}

	stringToLevelEncoding = map[string]LevelEncoding{
		"debug": LevelEncodingDebug,
		"info":  LevelEncodingInfo,
		"warn":  LevelEncodingWarn,
		"error": LevelEncodingError,
	}
)

func (e LevelEncoding) String() string {
	return levelEncodingToString[e]
}

func (e LevelEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}
