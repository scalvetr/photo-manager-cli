package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"time"
)

type Action uint

const (
	Undefined Action = iota
	UpdateDateFromMetadata
	UpdateMetadata
	UploadAlbums
	CheckAlbumDateMismatch
	IncreaseDate
)

var ActionFromString = map[string]Action{
	"UNDEFINED":                 Undefined,
	"UPDATE_DATE_FROM_METADATA": UpdateDateFromMetadata,
	"UPDATE_METADATA":           UpdateMetadata,
	"UPLOAD_ALBUMS":             UploadAlbums,
	"CHECK_ALBUM_DATE_MISMATCH": CheckAlbumDateMismatch,
	"INCREASE_DATE":             IncreaseDate,
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
	Regexp                   string               `yaml:"regexp"`
	ReportFile               string               `yaml:"report_file"`
	UpdateMetadataDateConfig UpdateMetadataConfig `yaml:"update_metadata_config"`
	AlbumInfoConfig          AlbumInfoConfig      `yaml:"album_info_config"`
	IncreaseDateConfig       IncreaseDateConfig   `yaml:"increase_date_config"`
}

type AlbumInfoConfig struct {
	FolderRegexp     string `yaml:"folder_regexp"`
	AlbumNamePattern string `yaml:"album_name_pattern"`
}
type IncreaseDateConfig struct {
	DateRangeFrom   time.Time `yaml:"date_range_from"`
	DateRangeTo     time.Time `yaml:"date_range_to"`
	IncreaseSeconds int       `yaml:"increase_seconds"`
}

func (c AlbumInfoConfig) GetFolderRegexp() string {
	if c.FolderRegexp == "" {
		return `(?P<year>\d{4}) - (?P<month>\d{2})(.*) - (?P<name>.*)`
	}
	return c.FolderRegexp

}
func (c AlbumInfoConfig) GetAlbumNamePattern() string {
	if c.AlbumNamePattern == "" {
		return `{{printf "%04d" .Year}}-{{printf "%02d" .Month}} - {{.Name}}`
	}
	return c.AlbumNamePattern
}

type UpdateMetadataConfig struct {
	Override         bool                              `yaml:"override"` // keep if there is one date
	FileDate         bool                              `yaml:"file_date"`
	Date             string                            `yaml:"date"`
	DateReplaces     []UpdateMetadataDateConfigReplace `yaml:"date_replaces"`
	DateFilePatterns []string                          `yaml:"date_file_patterns"`
}

type UpdateMetadataDateConfigReplace struct {
	Day    string `yaml:"day"`
	NewDay string `yaml:"new_day"`
}
