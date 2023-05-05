package types

import "github.com/zgwit/iot-master/v3/model"

type Alarm struct {
	Id       int64      `json:"id"`
	Device   string     `json:"device" xorm:"-"`
	DeviceId string     `json:"device_id" xorm:"index"`
	Type     string     `json:"type"`
	Title    string     `json:"title"`
	Message  string     `json:"message,omitempty"`
	Level    uint       `json:"level"`
	Read     bool       `json:"read,omitempty"`
	Created  model.Time `json:"created,omitempty" xorm:"created"`
}
