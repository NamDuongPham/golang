package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDelete
)

var allItemStatuses = [3]string{"Doing", "Done", "Delete"}

func (item *ItemStatus) String() string {
	return allItemStatuses[*item]
}
func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid item status")
}

// đọc sql dưới cơ sở dữ liệu ra ItemmStatus
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan item status: %v", value)
	}

	v, err := parseStr2ItemStatus(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to scan item status: %v", value)
	}
	*item = v
	return nil

}

// lấy từ ItemStatus ra sql
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

// hỗ trợ đổi ItemStatus sang json value
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// hỗ trợ đổi json value  sang ItemStatus
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemValue
	return nil
}
