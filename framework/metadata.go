package framework

import (
	"github.com/Meduzz/helper/meta/schema"
)

type (
	Registry struct {
		summer *Summer
		meta   map[string]*Metadata
	}

	Metadata struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Input       *schema.Schema `json:"input"`
		Output      *schema.Schema `json:"output,omitempty"`
	}
)
