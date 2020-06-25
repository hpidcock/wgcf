package util

import (
	"time"

	"github.com/pkg/errors"

	"github.com/hpidcock/wgcf/cloudflare/structs/resp"
)

func GetTimestamp() string {
	return getTimestamp(time.Now())
}

func getTimestamp(t time.Time) string {
	timestamp := t.Format(time.RFC3339Nano)
	return timestamp
}

func FindDevice(result *resp.BoundDevicesData, deviceId string) (*resp.BoundDevice, error) {
	for _, device := range *result {
		if device.ID == deviceId {
			return &device, nil
		}
	}
	return nil, errors.New("device not found in list")
}
