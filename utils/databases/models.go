// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package databases

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Model 数据库base model
type Model struct {
	ID        int64      `json:"id,string" gorm:"primary_key"`
	CreatedAt TimeInt64  `json:"createdAt" gorm:"type:datetime"`
	UpdatedAt TimeInt64  `json:"updatedAt" gorm:"type:datetime"`
	DeletedAt *TimeInt64 `json:"deletedAt" gorm:"type:datetime"`
}

// Int64Str Int64 String
type Int64Str int64

// MarshalJSON Int64Str MarshalJSON
func (i Int64Str) MarshalJSON() ([]byte, error) {
	var s string
	s = strconv.FormatInt(int64(i), 10)
	if i == 0 {
		s = ""
	}
	return json.Marshal(s)
}

// UnmarshalJSON Int64Str UnmarshalJSON
func (i *Int64Str) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Int64Str(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*int64)(i))
}

// TimeInt64 自定义时间结论.json为 unix time string
type TimeInt64 time.Time

// MarshalJSON Int64Str MarshalJSON
func (t TimeInt64) MarshalJSON() ([]byte, error) {
	var s string
	i := time.Time(t).Unix() * 1000
	s = strconv.FormatInt(int64(i), 10)
	if i == 0 {
		s = ""
	}
	return json.Marshal(s)
}

// UnmarshalJSON Int64Str UnmarshalJSON
func (t *TimeInt64) UnmarshalJSON(b []byte) error {
	var s string
	s = string(b)
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	value = value / 1000

	*t = TimeInt64(time.Unix(value, 0))
	return nil
}

// Value insert timestamp into mysql need this function.
func (t TimeInt64) Value() (driver.Value, error) {
	var zeroTime time.Time
	ti := time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (t *TimeInt64) Scan(v interface{}) error {
	ti, ok := v.(time.Time) // NOT directly assertion v.(TimeNormal)
	if ok {
		*t = TimeInt64(ti)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
