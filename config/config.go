package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Action uint

const (
	Undefined Action = iota
	UpdateDateFromMetadata
	UpdateMetadata
	UploadAlbums
)

var ActionFromString = map[string]Action{
	"UNDEFINED":                 Undefined,
	"UPDATE_DATE_FROM_METADATA": UpdateDateFromMetadata,
	"UPDATE_METADATA":           UpdateMetadata,
	"UPLOAD_ALBUMS":             UploadAlbums,
}

func (a *Action) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}
	var v Action
	var ok bool
	if v, ok = ActionFromString[s]; !ok {
		return fmt.Errorf("unknown user type %s", s)
	}
	*a = v
	return nil
}

type Config struct {
	Action                   Action               `yaml:"action" `
	Path                     string               `yaml:"path"`
	UpdateMetadataDateConfig UpdateMetadataConfig `yaml:"update_metadata_config"`
}

type UpdateMetadataConfig struct {
	Date     string                            `yaml:"date"`
	Override bool                              `yaml:"override"`
	Replace  []UpdateMetadataDateConfigReplace `yaml:"replace"`
}
type UpdateMetadataDateConfigReplace struct {
	Day    string `yaml:"day"`
	NewDay string `yaml:"new_day"`
}
