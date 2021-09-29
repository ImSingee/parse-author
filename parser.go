package pauthor

import (
	"fmt"
	"strings"
)

type Author struct {
	Name  string `json:"name" toml:"name" yaml:"name" mapstructure:"name"`
	Email string `json:"email" toml:"email" yaml:"email" mapstructure:"email"`
	Url   string `json:"url" toml:"url" yaml:"url" mapstructure:"url"`
}

var ErrIsEmpty = fmt.Errorf("author is empty")
var ErrInvalid = fmt.Errorf("author is invalid")

func Parse(s string) (*Author, error) {
	s = strings.TrimSpace(s)

	if s == "" {
		return nil, ErrIsEmpty
	}

	as := Regex().FindStringSubmatch(s)
	if len(as) != 4 {
		return nil, ErrInvalid
	}

	return &Author{Name: as[1], Email: as[2], Url: as[3]}, nil
}
