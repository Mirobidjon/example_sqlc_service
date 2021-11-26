// Code generated by sqlc. DO NOT EDIT.

package sqlc

import (
	"fmt"
)

type AttributeTypes string

const (
	AttributeTypesDatetime AttributeTypes = "datetime"
	AttributeTypesText     AttributeTypes = "text"
	AttributeTypesNumber   AttributeTypes = "number"
)

func (e *AttributeTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AttributeTypes(s)
	case string:
		*e = AttributeTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for AttributeTypes: %T", src)
	}
	return nil
}

type Attribute struct {
	ID   string         `json:"id"`
	Name string         `json:"name"`
	Type AttributeTypes `json:"type"`
}

type Position struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ProfessionID string `json:"profession_id"`
}

type PositionAttribute struct {
	ID          string `json:"id"`
	AttributeID string `json:"attribute_id"`
	PositionID  string `json:"position_id"`
	Value       string `json:"value"`
}

type Profession struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
