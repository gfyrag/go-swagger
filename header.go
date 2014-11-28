package swagger

import (
	"encoding/json"

	"github.com/casualjim/go-swagger/reflection"
)

// Describes a header for a response of the API
//
// For more information: http://goo.gl/8us55a#headerObject
type Header struct {
	Description      string        `swagger:"description,omitempty"`
	Maximum          float64       `swagger:"maximum,omitempty"`
	ExclusiveMaximum bool          `swagger:"exclusiveMaximum,omitempty"`
	Minimum          float64       `swagger:"minimum,omitempty"`
	ExclusiveMinimum bool          `swagger:"exclusiveMinimum,omitempty"`
	MaxLength        int64         `swagger:"maxLength,omitempty"`
	MinLength        int64         `swagger:"minLength,omitempty"`
	Pattern          string        `swagger:"pattern,omitempty"`
	MaxItems         int64         `swagger:"maxItems,omitempty"`
	MinItems         int64         `swagger:"minItems,omitempty"`
	UniqueItems      bool          `swagger:"uniqueItems,omitempty"`
	MultipleOf       float64       `swagger:"multipleOf,omitempty"`
	Enum             []interface{} `swagger:"enum,omitempty"`
	Type             string        `swagger:"type,omitempty"`
	Format           string        `swagger:"format,omitempty"`
	Default          interface{}   `swagger:"default,omitempty"`
	Items            *Items        `swagger:"-"`
}

func (h Header) MarshalJSON() ([]byte, error) {
	return json.Marshal(reflection.MarshalMapRecursed(h))
}

func (h Header) MarshalYAML() (interface{}, error) {
	return reflection.MarshalMap(h), nil
}