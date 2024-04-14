package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Banner struct {
	Content `json:",inline"`

	ID        int   `json:"id"`
	TagIDs    []int `json:"tag_ids"`
	FeatureID int   `json:"feature_id"`
	IsActive  bool  `json:"is_active"`
}

type BannerFilter struct {
	FeatureID int
	TagID     int
	Offset    int
	Limit     int
}

type BannerCreated struct {
	ID int `json:"banner_id"`
}

type BannerDeleted struct {
	ID int `json:"banner_id"`
}

type Content struct {
	Content any `json:"content"`
}

func (r Content) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (c *Content) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}
