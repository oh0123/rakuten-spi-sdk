package types

import (
	"database/sql/driver"
	"errors"
	"rakuten-spi-sdk/config"
	"time"

	"github.com/bytedance/sonic"
)

type JpTime time.Time

func (j JpTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(config.CusTimeFormat)+len(`""`))
	b = append(b, '"')
	b = append(b, []byte(time.Time(j).Format(time.RFC1123)+"+0900")...)
	b = append(b, '"')
	return b, nil
}

func (j *JpTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	var err error
	t, err := time.Parse("2006-01-02T15:04:05+0900", string(data))
	*j = JpTime(t)
	return err
}

func (c *JpTime) Scan(value interface{}) error {
	t := value.([]byte)
	b := make([]byte, 0, len("2006-01-02T15:04:05+0900")+len(`""`))
	b = append(b, '"')
	b = append(b, t...)
	b = append(b, '"')
	return sonic.Unmarshal(b, c)
}

func (c JpTime) Value() (driver.Value, error) {
	b, err := sonic.Marshal(c)
	if err != nil {
		return nil, err
	}
	b = b[len(`"`) : len(b)-len(`"`)]
	return b, err
}

type JpDate time.Time

func (j JpDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateOnly)+len(`""`))
	b = append(b, '"')
	b = append(b, []byte(time.Time(j).Format(time.DateOnly))...)
	b = append(b, '"')
	return b, nil
}

func (j *JpDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	var err error
	t, err := time.Parse(time.DateOnly, string(data))
	*j = JpDate(t)
	return err
}

func (c *JpDate) Scan(value interface{}) error {
	t := value.([]byte)
	b := make([]byte, 0, len(time.DateOnly)+len(`""`))
	b = append(b, '"')
	b = append(b, t...)
	b = append(b, '"')
	return sonic.Unmarshal(b, c)
}

func (c JpDate) Value() (driver.Value, error) {
	b, err := sonic.Marshal(c)
	if err != nil {
		return nil, err
	}
	b = b[len(`"`) : len(b)-len(`"`)]
	return b, err
}

type Slice[T int | string | float32 | int32 | int64 | float64 | JpTime] []T

func (c *Slice[T]) Scan(value interface{}) error {
	return sonic.Unmarshal(value.([]byte), c)
}

func (c Slice[T]) Value() (driver.Value, error) {
	b, err := sonic.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, err
}

type Map[K string, V any] map[K]V

func (c *Map[K, V]) Scan(value interface{}) error {
	return sonic.Unmarshal(value.([]byte), c)
}

func (c Map[K, V]) Value() (driver.Value, error) {
	b, err := sonic.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, err
}

type Maps[K string, V any] []Map[K, V]

func (c *Maps[K, V]) Scan(value interface{}) error {
	return sonic.Unmarshal(value.([]byte), c)
}

func (c Maps[K, V]) Value() (driver.Value, error) {
	b, err := sonic.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, err
}
